package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const BLOCK_REWARD_WALLET string = "Block Reward"
const TRANSACTION_FEE float64 = 0.05
const GAS_FEE float64 = 0.1

func verifySignature(publicKey, message, signature string) (bool, error) {
	BACKEND_URL := os.Getenv("BACKEND_URL")
	if BACKEND_URL == "" { BACKEND_URL = "http://localhost:9000/api" }
	// Prepare the verification request
	verificationRequest := fiber.Map{
		"public_key": publicKey,
		"message":    message,
		"sign":       signature,
	}

	// Make the request to the verification API
	client := &http.Client{}
	reqBody, err := json.Marshal(verificationRequest)
	if err != nil {
		return false, fmt.Errorf("failed to prepare verification request: %v", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/verify",BACKEND_URL), strings.NewReader(string(reqBody)))
	if err != nil {
		return false, fmt.Errorf("failed to create verification request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to verify signature: %v", err)
	}
	defer resp.Body.Close()

	// Parse the verification response
	var verifyResponse struct {
		Verify bool `json:"verify"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&verifyResponse); err != nil {
		return false, fmt.Errorf("failed to parse verification response: %v", err)
	}

	// Return true if verification succeeded, otherwise false
	return verifyResponse.Verify, nil
}

// Block represents each 'item' in the blockchain
type Block struct {
	Data         BlockData `json:"data"`
	PreviousHash string    `json:"previous_hash"`
	Hash         string    `json:"hash"`
	Timestamp    time.Time `json:"timestamp"`
	Nonce        int       `json:"nonce"`
}

// Blockchain represents the entire chain
type Blockchain struct {
	GenesisBlock          Block
	Chain                 []Block
	TransactionPool       []Transaction
	ContractExecutionPool []ContractExecution
	Difficulty            int
	RewardPerBlock        float64
	MaxCoins              float64
}

func (bc *Blockchain) appendNewEmptyBlock() {
	block := Block{
		PreviousHash: bc.getLastBlock().Hash,
		Timestamp:    time.Now(),
	}
	block.Hash = block.calculateHash()
	bc.Chain = append(bc.Chain, block)
}

func (bc Blockchain) getLastBlock() *Block {
	return &bc.Chain[len(bc.Chain)-1]
}

// calculateHash calculates the hash of a block
func (b Block) calculateHash() string {
	data, _ := json.Marshal(b.Data)
	blockData := b.PreviousHash + string(data) + b.Timestamp.String() + strconv.Itoa(b.Nonce)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

// mine mines a block
func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.Hash, strings.Repeat("0", difficulty)) {
		b.Nonce++
		b.Hash = b.calculateHash()
	}
}

// CreateBlockchain creates a new blockchain with a genesis block
func CreateBlockchain(difficulty int, rewardPerBlock float64, maxCoins float64) Blockchain {
	genesisBlock := Block{
		Timestamp: time.Now(),
	}
	genesisBlock.Hash = genesisBlock.calculateHash() // Set initial hash without mining
	return Blockchain{
		GenesisBlock:   genesisBlock,
		Chain:          []Block{genesisBlock},
		Difficulty:     difficulty,
		RewardPerBlock: rewardPerBlock,
		MaxCoins:       maxCoins,
	}
}

func (bc *Blockchain) findContractByID(contractID string) *SmartContract {
	for i := len(bc.Chain) - 1; i >= 0; i-- {
		block := bc.Chain[i]
		for j := range block.Data.Contracts {
			if block.Data.Contracts[j].ContractID == contractID {
				return &block.Data.Contracts[j]
			}
		}
	}
	return nil
}

// addContract adds a smart contract directly to the current block
func (bc *Blockchain) addContract(contract SmartContract) {
	lastBlock := bc.getLastBlock()
	lastBlock.Data.Contracts = append(lastBlock.Data.Contracts, contract)
}

// addTransaction adds a transaction to the transaction pool after validating it
func (bc *Blockchain) addTransaction(tx Transaction) error {
	// Validate the transaction
	isValid, errMsg := tx.Validate(bc)
	if !isValid {
		return fmt.Errorf("%s", errMsg)
	}

	// If valid, add the transaction to the pool
	bc.TransactionPool = append(bc.TransactionPool, tx)
	return nil
}

// mineTransaction mines transactions from the transaction pool into the current block
func (bc *Blockchain) mineTransaction(miner string) error {
	if len(bc.TransactionPool) == 0 {
		return fmt.Errorf("no transactions to mine")
	}

	lastBlock := bc.getLastBlock()

	// Process the first transaction in the pool (FIFO)
	transaction := bc.TransactionPool[0]
	transaction.Miner = miner
	lastBlock.Data.Transactions = append(lastBlock.Data.Transactions, transaction)

	// Remove the processed transaction from the pool
	bc.TransactionPool = bc.TransactionPool[1:]

	return nil
}

// mineContractExecution mines contract executions from the execution pool into the current block
func (bc *Blockchain) mineContractExecution(miner string) float64 {
	lastBlock := bc.getLastBlock()

	if len(bc.ContractExecutionPool) > 0 {
		// Process the first contract execution in the pool (FIFO)
		execpool := bc.ContractExecutionPool[0]

		// Execute the contract and get the result
		contract := bc.findContractByID(execpool.ContractID)
		if contract != nil {
			result, err := contract.Execute(bc)
			if err != nil {
				execpool.Result = fmt.Sprintf("Execution failed: %s", err.Error())
			} else {
				execpool.Result = result
			}
			execpool.Miner = miner
			lastBlock.Data.ContractExecutionHistory = append(lastBlock.Data.ContractExecutionHistory, execpool)
		}

		// Remove the processed contract execution from the pool
		bc.ContractExecutionPool = bc.ContractExecutionPool[1:]
		return GAS_FEE
	}
	return 0
}

func (bc *Blockchain) mineBlock(miner string) (Block, error) {
	currentBlock := bc.getLastBlock()
	// Determine the block reward based on the maximum coins limit
	if bc.getMinedCoins()+bc.RewardPerBlock <= bc.MaxCoins {
		currentBlock.Data.Transactions = append(currentBlock.Data.Transactions, Transaction{
			From:   BLOCK_REWARD_WALLET,
			To:     miner,
			Amount: bc.RewardPerBlock,
		})
	}

	currentBlock.mine(bc.Difficulty)
	bc.appendNewEmptyBlock()

	// Return the mined block
	return *currentBlock, nil
}

// isValid checks if the blockchain is valid
func (bc Blockchain) isValid() bool {
	for i := 0; i < len(bc.Chain)-2; i++ {
		previousBlock := bc.Chain[i]
		currentBlock := bc.Chain[i+1]
		if currentBlock.Hash != currentBlock.calculateHash() || currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}

// getBalance calculates the balance of a specific address
func (bc *Blockchain) getBalance(address string) float64 {
	balance := 0.0
	for _, block := range bc.Chain {
		for _, tx := range block.Data.Transactions {
			if tx.From == address {
				balance -= (tx.Amount + TRANSACTION_FEE)
			} else if tx.To == address {
				balance += tx.Amount
			}

			if tx.Miner == address {
				balance += TRANSACTION_FEE
			}
		}
		for _, exec := range block.Data.ContractExecutionHistory {
			if exec.Miner == address {
				balance += GAS_FEE
			}
		}
	}
	return balance
}

// getMinedCoins calculates the total mined coins
func (bc Blockchain) getMinedCoins() float64 {
	totalMined := 0.0
	for _, block := range bc.Chain {
		for _, tx := range block.Data.Transactions {
			if tx.From == BLOCK_REWARD_WALLET {
				totalMined += tx.Amount
			}
		}
	}
	return totalMined
}

// generateRandomID generates a random 16-byte hex string
func generateRandomID() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// main sets up the server and routes
func main() {
	app := fiber.New()

	// Initialize the blockchain with a difficulty of 2, reward of 10 coins per block, and a maximum of 1000 coins
	blockchain := CreateBlockchain(2, 10, 1000)

	// Middleware to set blockchain in context
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("blockchain", &blockchain)
		return c.Next()
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// Mine a new block
	app.Post("/mine/block", func(c *fiber.Ctx) error {
		blockchain := c.Locals("blockchain").(*Blockchain)
		var request struct {
			Wallet string `json:"wallet"`
		}
		// Parse the request body
		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
		}

		miner := request.Wallet
		if miner == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Missing miner wallet")
		}
		block, err := blockchain.mineBlock(miner)
		if err != nil {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}

		response := fiber.Map{
			"message": "New Block Forged",
			"index":   len(blockchain.Chain) - 1,
			"block":   block,
		}
		return c.Status(fiber.StatusOK).JSON(response)
	})

	app.Post("/mine/transaction", func(c *fiber.Ctx) error {
		blockchain := c.Locals("blockchain").(*Blockchain)
		var request struct {
			Wallet string `json:"wallet"`
		}
		// Parse the request body
		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
		}

		miner := request.Wallet
		if miner == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Missing miner wallet")
		}

		// Mine the transaction
		err := blockchain.mineTransaction(miner)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "No transactions to mine"})
		}

		response := fiber.Map{
			"message": "Transaction mined successfully",
		}
		return c.Status(fiber.StatusOK).JSON(response)
	})

	// Mine contract executions
	app.Post("/mine/contract", func(c *fiber.Ctx) error {
		blockchain := c.Locals("blockchain").(*Blockchain)
		var request struct {
			Wallet string `json:"wallet"`
		}
		// Parse the request body
		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
		}

		miner := request.Wallet
		if miner == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Missing miner wallet")
		}

		// Mine and process the contract executions
		gas := blockchain.mineContractExecution(miner)

		if gas != 0 {
			response := fiber.Map{
				"message": "Contract Executed Successfully",
				"gas":     gas,
			}
			return c.Status(fiber.StatusOK).JSON(response)
		} else {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "No contracts to mine"})
		}
	})

	// Add new block data (transaction)
	app.Post("/transaction/new", func(c *fiber.Ctx) error {
		var tx Transaction
		if err := c.BodyParser(&tx); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
		}

		blockchain := c.Locals("blockchain").(*Blockchain)
		if err := blockchain.addTransaction(tx); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		response := fiber.Map{"message": "Transaction added to the pool"}
		return c.Status(fiber.StatusCreated).JSON(response)
	})

	// Add new smart contract
	app.Post("/certificate/request", func(c *fiber.Ctx) error {
		var request struct {
			Certificate string `json:"certificate"`
			Domain      string `json:"domain"`
		}

		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
		}

		contractID, err := generateRandomID()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Could not generate contract ID")
		}

		wallet, err := GetPublicKeyFromCertificate(request.Certificate)
		if err != nil { return c.Status(fiber.StatusInternalServerError).SendString("Could not extract public key from certificate") }

		smartContract := SmartContract{
			ContractID: contractID,
			Wallet: wallet,
			Code:       &Certificate {
				ContractID: contractID,
				Domain: request.Domain,
				Certificate: request.Certificate,
				CreatedAt: time.Now(),
			},
		}

		blockchain := c.Locals("blockchain").(*Blockchain)
		blockchain.addContract(smartContract)

		response := fiber.Map{
			"message":    "Certificate successfully added to the current block, complete acme challenge to turn it valid.",
			"contract_id": smartContract.ContractID,
			"domain": request.Domain,
			"wallet": smartContract.Wallet,
			"certificate": request.Certificate,
		}
		return c.Status(fiber.StatusCreated).JSON(response)
	})

	// Execute a smart contract (add to execution pool)
	app.Post("/contract/execute", func(c *fiber.Ctx) error {
		// Define a struct to parse the request body
		var request struct {
			ContractID string `json:"contract_id"`
		}

		// Parse the request body
		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
		}

		// Validate the contract ID
		if request.ContractID == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Missing contract ID")
		}

		// Extract the signature from the Authorization header
		signature := c.Get("Authorization")
		if signature == "" {
			return c.Status(fiber.StatusUnauthorized).SendString("Missing signature")
		}

		// Find the contract in the blockchain
		blockchain := c.Locals("blockchain").(*Blockchain)
		contract := blockchain.findContractByID(request.ContractID)
		if contract == nil {
			return c.Status(fiber.StatusNotFound).SendString("Contract not found")
		}

		// Verify the signature using the helper function
		isVerified, err := verifySignature(contract.Wallet, request.ContractID, signature)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		if !isVerified {
			return c.Status(fiber.StatusUnauthorized).SendString("Signature verification failed")
		}

		// Add the contract execution request to the ContractExecutionPool
		execution := ContractExecution{
			ContractID: request.ContractID,
			Timestamp:  time.Now(),
		}

		blockchain.ContractExecutionPool = append(blockchain.ContractExecutionPool, execution)

		response := fiber.Map{
			"message": "Contract execution added to the pool",
		}
		return c.Status(fiber.StatusCreated).JSON(response)
	})

	// Get the full blockchain
	app.Get("/chain", func(c *fiber.Ctx) error {
		blockchain := c.Locals("blockchain").(*Blockchain)
		response := fiber.Map{
			"chain":      blockchain.Chain,
			"length":     len(blockchain.Chain),
			"isValid":    blockchain.isValid(),
			"minedCoins": blockchain.getMinedCoins(),
		}
		return c.Status(fiber.StatusOK).JSON(response)
	})

	// Get data from the transaction pool
	app.Get("/memorypool", func(c *fiber.Ctx) error {
		blockchain := c.Locals("blockchain").(*Blockchain)
		response := fiber.Map{
			"transactionpool":       blockchain.TransactionPool,
			"contractexecutionpool": blockchain.ContractExecutionPool,
		}
		return c.Status(fiber.StatusOK).JSON(response)
	})

	// Get information of a wallet
	app.Get("/info/:wallet", func(c *fiber.Ctx) error {
		blockchain := c.Locals("blockchain").(*Blockchain)
		wallet := c.Params("wallet") // Get the wallet from the URL path

		response := fiber.Map{
			"balance": blockchain.getBalance(wallet),
		}
		return c.Status(fiber.StatusOK).JSON(response)
	})

	app.Get("/certificate/status/:base64cert", func(c *fiber.Ctx) error {
		blockchain := c.Locals("blockchain").(*Blockchain)
		base64Cert := c.Params("base64cert")

		status, err := blockchain.getCertificateStatus(base64Cert)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		response := fiber.Map{
			"certificate": base64Cert,
			"status":      status,
		}
		return c.Status(fiber.StatusOK).JSON(response)
	})

	app.Listen(":7000")
}
