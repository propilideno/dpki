package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"time"
	"net/http"
	"encoding/json"
	"os"
)

// Certificate implements the Code interface for smart contract of type Certificates
type Certificate struct {
	ContractID  string    `json:"contract_id"`
	Domain      string    `json:"domain"`
	Certificate string    `json:"string"`
	CreatedAt   time.Time `json:"created_at"`
	Status      bool      `json:"status"`
}

func (sc *Certificate) Execute(blockchain *Blockchain) (string, error) {
	BACKEND_URL := os.Getenv("BACKEND_URL")
	if BACKEND_URL == "" { BACKEND_URL = "http://localhost:9000/api" }
	// Construct the URL for the DNS check
	dnsURL := fmt.Sprintf("%s/dns/%s", BACKEND_URL,sc.Domain)
	resp, err := http.Get(dnsURL)
	if err != nil {
		return "invalid", nil // Mark as invalid if the request fails
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		// If the status code is not 200, return invalid result
		return "invalid", nil
	}

	// Parse the response
	var dnsResponse struct {
		ID        int       `json:"id"`
		Domain    string    `json:"domain"`
		TXT       string    `json:"txt"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&dnsResponse); err != nil {
		return "invalid", nil // Return invalid if decoding fails
	}

	if dnsResponse.TXT == sc.ContractID {
		sc.Status = true
		return "valid", nil
	}

	return "invalid", nil
}


func (sc *Certificate) Validate(blockchain *Blockchain) bool {
	// Add validation logic for the smart contract of type Certificate
	fmt.Println("Validating smart contract of type Certificate...")
	return true
}

func GetPublicKeyFromCertificate(base64Cert string) (string, error) {
	// Step 1: Decode the base64-encoded certificate
	certPEM, err := base64.StdEncoding.DecodeString(base64Cert)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64: %w", err)
	}
	// Step 2: Decode the PEM block
	block, _ := pem.Decode(certPEM)
	if block == nil || block.Type != "CERTIFICATE" {
		return "", fmt.Errorf("failed to parse PEM block containing the certificate")
	}
	// Step 3: Parse the certificate
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse certificate: %w", err)
	}
	// Step 4: Extract the RSA public key
	rsaPublicKey, ok := cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("certificate does not contain an RSA public key")
	}
	// Step 5: Encode the RSA public key in PEM format (PKIX format)
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(rsaPublicKey)
	if err != nil {
		return "", fmt.Errorf("failed to marshal public key: %w", err)
	}
	// Step 6: Encode the public key bytes into PEM format
	rsaPublicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	// Step 7: Encode the entire PEM block to base64
	base64EncodedPEM := base64.StdEncoding.EncodeToString(rsaPublicKeyPEM)
	// Return the base64-encoded public key including PEM headers and footers
	return base64EncodedPEM, nil
}

// getCertificateStatus returns the status of the last contract execution for a given certificate
func (bc *Blockchain) getCertificateStatus(base64Cert string) (bool, error) {
	for i := len(bc.Chain) - 1; i >= 0; i-- { // Iterate over the blockchain in reverse order
		block := bc.Chain[i]
		for _, exec := range block.Data.ContractExecutionHistory {
			contract := bc.findContractByID(exec.ContractID)
			if contract != nil {
				cert, ok := contract.Code.(*Certificate)
				if ok && cert.Certificate == base64Cert {
					return cert.Status, nil // Return the status of the certificate
				}
			}
		}
	}
	return false, fmt.Errorf("Certificate not found or has not been executed yet")
}
