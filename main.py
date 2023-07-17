import hashlib

class neuralCoinBlock:

    def __init__(self,previousBlockHash,transactionList):
        self.previousBlockHash = previousBlockHash
        self.transactionList = transactionList

        self.blockData = "-".join(transactionList) +"-"+previousBlockHash
        self.blockHash = hashlib.sha256(self.blockData.encode()).hexdigest()