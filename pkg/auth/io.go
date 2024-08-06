package auth

import (
	"crypto/rsa"
	"encoding/pem"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

// LoadRSAPublicKeyFromDisk loads an RSA public key from a PEM file.
func LoadRSAPublicKeyFromDisk(path string) (*rsa.PublicKey, error) {
	// Read the public key file
	keyData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read public key file: %w", err)
	}

	// Decode the PEM block
	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}

	// Parse the public key
	pub, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		return nil, fmt.Errorf("could not parse RSA public key: %w", err)
	}

	return pub, nil
}

// LoadRSAPrivateKeyFromDisk loads an RSA private key from a PEM file.
func LoadRSAPrivateKeyFromDisk(path string) (*rsa.PrivateKey, error) {
	// Read the public key file
	keyData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read public key file: %w", err)
	}

	// Decode the PEM block
	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}

	// Parse the public key
	priv, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		return nil, fmt.Errorf("could not parse RSA public key: %w", err)
	}

	return priv, nil
}
