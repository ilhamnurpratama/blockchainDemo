package chaincode

import (
        "encoding/json"
        "fmt"

        "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
        contractapi.Contract
}

// Ledger definition parameter
type Asset struct {
    call_id string `json:"call_id`
	a_number string `json:"a_number"`
	b_number string `json:"b_number"`
	a_min_record string `json:"a_min_record`
	b_min_record string `json:"b_min_record`
	a_fee string `json:"a_fee"`
	b_fee string `json:"b_fee"`
	a_mdr string `json:"a_mdr"`
	b_mdr string `json:"b_mdr"`
}

// Create Call function
func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, call_id string,a_number string,b_number string,a_min_record string,b_min_record string,a_fee string,b_fee string,a_mdr string,b_mdr string) error {
	exists, err := s.AssetExists(ctx, call_id)
	if err != nil {
			return err
	}
	if exists {
			return fmt.Errorf("the asset %s already exists", call_id)
	}

	asset := Asset{
		call_id:	call_id,
		a_number:	a_number,
		b_number:	b_number,
		a_min_record:	a_min_record,
		b_min_record:	b_min_record,
		a_fee:		a_fee,
		b_fee:		b_fee,
		a_mdr:		a_mdr,
		b_mdr:		b_mdr,
	}

	assetJSON, err := json.Marshal(asset)
	if err != nil {
			return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// Update Call log
func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, id string, color string, size int, owner string, appraisedValue int) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
			return err
	}
	if !exists {
			return fmt.Errorf("the asset %s does not exist", id)
	}

	// overwriting original asset with new asset
	asset := Asset{
		a_min_record:	a_min_record,
		b_min_record:	b_min_record,
		a_fee:		a_fee,
		b_fee:		b_fee,
		a_mdr:		a_mdr,
		b_mdr:		b_mdr,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
			return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// Asset exist function
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, call_id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(call_id)
	if err != nil {
			return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

// get all assets
func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
			return nil, err
	}
	defer resultsIterator.Close()

	var assets []*Asset
	for resultsIterator.HasNext() {
			queryResponse, err := resultsIterator.Next()
			if err != nil {
					return nil, err
			}

			var asset Asset
			err = json.Unmarshal(queryResponse.Value, &asset)
			if err != nil {
					return nil, err
			}
			assets = append(assets, &asset)
	}

	return assets, nil
}