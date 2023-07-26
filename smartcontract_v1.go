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

// Command for creating Ledger parameter
type Ledger struct{
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

// Function to create a call
func (s *SmartContract) create_call(ctx contractapi.TransactionContextInterface, call_id string,a_number string,b_number string,a_min_record string,b_min_record string,a_fee string,b_fee string,a_mdr string,b_mdr string){
	exist, err := s.AssetExists(ctx,call_id)

	if err != nil {
			return err
	}
	if exists {
			return fmt.Errorf("the asset %s already exists", id)
	}

	ledger := Ledger{
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

	ledger_json, err := json.Marshal(asset)
	if err != nil {
			return err
	}

	return ctx.GetStub().PutState(id, ledger_json)
}

// AssetExists returns true when asset with given ID exists in world state
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, call_id string) (bool, error) {
	ledger_json, err := ctx.GetStub().GetState(call_id)
	if err != nil {
			return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return ledger_json != nil, nil
}

// Read Transactions of previous call
func (s *SmartContract) read_transaction(ctx contractapi.TransactionContextInterface, call_id string) (*Ledger, error) {
	ledger_json, err := ctx.GetStub().GetState(call_id)
	if err != nil {
			return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if ledger_json == nil {
			return nil, fmt.Errorf("the asset %s does not exist", call_id)
	}

	var ledger Ledger
	err = json.Unmarshal(ledger_json, &ledger)
	if err != nil {
			return nil, err
	}

	return &ledger, nil
}

// Update transactions on Telkomsel Ledger
func (s *SmartContract) update_transaction(ctx contractapi.TransactionContextInterface, call_id string, a_min_record string, b_min_record, a_fee string, b_fee string,a_mdr string,b_mdr string) error {
	exists, err := s.AssetExists(ctx, call_id)
	if err != nil {
			return err
	}
	if !exists {
			return fmt.Errorf("the asset %s does not exist", call_id)
	}

	// overwriting original asset with new asset
	ledger := Ledger{
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

	ledger_json, err := json.Marshal(ledger)
	if err != nil {
			return err
	}

	return ctx.GetStub().PutState(call_id, ledger_json)
}

// DeleteAsset deletes an given asset from the world state.
func (s *SmartContract) delete_record(ctx contractapi.TransactionContextInterface, call_id string) error {
	exists, err := s.AssetExists(ctx, call_id)
	if err != nil {
			return err
	}
	if !exists {
			return fmt.Errorf("the asset %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}