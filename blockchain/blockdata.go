package main

// BlockData contains all types of data that can be part of a block
type BlockData struct {
	ContractExecutionHistory []ContractExecution `json:"contract_execution_history"`
	Contracts                []SmartContract     `json:"contracts"`
	Transactions             []Transaction       `json:"transactions"`
}

// Transaction represents a blockchain transaction
type Transaction struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
	Miner  string  `json:"miner"`
}

// Validate checks if the transaction is valid
func (t Transaction) Validate(blockchain *Blockchain) (bool, string) {
	if t.From == t.To {
		return false, "Sender and recipient cannot be the same"
	}
	if t.Amount <= 0 {
		return false, "Transaction amount must be greater than zero"
	}
	if t.From == BLOCK_REWARD_WALLET || t.To == BLOCK_REWARD_WALLET {
		return false, "Direct transactions to or from the block reward wallet are not allowed"
	}
	balance := blockchain.getBalance(t.From)
	if balance < t.Amount+TRANSACTION_FEE {
		return false, "Insufficient balance to cover the transaction amount and fee"
	}
	return true, ""
}
