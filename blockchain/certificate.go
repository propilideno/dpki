package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"time"
)

// Certificate implements the Code interface for smart contract of type Certificates
type Certificate struct {
	Domain      string    `json:"domain"`
	Certificate string    `json:"string"`
	CreatedAt   time.Time `json:"created_at"`
	Status      bool      `json:"status"`
}

func (sc *Certificate) Execute(blockchain *Blockchain) error {
	// Add logic to process the smart contract of type Certificate
	fmt.Println("Executing smart contract of type Certificate...")
	return nil
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

