/** logic.js 
 * Version 0.1.1
 * Last Update 14-08-2023
 * @param {org.tsel.network.Recon} recon - the recon of the call to be processed
 * @transaction
*/

async function reconProcess(recon){
    /** To set ledger owner to new trader */
    recon.ledger.owner = recon.newOwner

    /** To get the asset from the model */
    let assetRegistry = await getAssetRegistry('org.tsel.network.Ledger');
    await assetRegistry.update(recon.ledger);
}