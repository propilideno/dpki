package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"
)

// SmartContract represents a smart contract in the blockchain
type SmartContract struct {
	ContractID    string `json:"contract_id"`
	Wallet        string `json:"wallet"`
	Code          Code   `json:"-"`
}

type ContractExecution struct {
	ContractID  string    `json:"contract_id"`
	Result      string    `json:"result"`
	Timestamp   time.Time `json:"timestamp"`
	Miner       string    `json:"miner"`
}

// Code interface defines the methods for a smart contract
type Code interface {
	Execute(blockchain *Blockchain) (string, error)
	Validate(blockchain *Blockchain) bool
}

// Execute calls the Execute method of the Code interface
func (sc *SmartContract) Execute(blockchain *Blockchain) (string, error) {
	return sc.Code.Execute(blockchain)
}

// Validate calls the Validate method of the Code interface
func (sc *SmartContract) Validate(blockchain *Blockchain) bool {
	return sc.Code.Validate(blockchain)
}

// calculateDigest generates a SHA256 digest of the contract data
func (sc *SmartContract) calculateDigest() string {
	data, _ := json.Marshal(sc)
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}
