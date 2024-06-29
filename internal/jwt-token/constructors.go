package jwttoken

import (
	"crypto/rsa"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type TokenGenerator struct {
	privateKey *rsa.PrivateKey
}

type TokenVerifier struct {
	publicKey *rsa.PublicKey
}

func NewTokenGenerator(pvtKeyFilepath string) (*TokenGenerator, error) {
	privateKeyData, err := os.ReadFile(pvtKeyFilepath)
	if err != nil {
		return nil, fmt.Errorf("error reading private key: %v", err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyData)
	if err != nil {
		return nil, fmt.Errorf("error parsing private key: %v", err)
	}
	return &TokenGenerator{
		privateKey: privateKey,
		// expiryTime: expiryTime,
	}, nil
}

func NewTokenVerifier(filepath string) (*TokenVerifier, error) {
	publicKeyData, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("error reading public key: %v", err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyData)
	if err != nil {
		return nil, fmt.Errorf("error parsing public key: %v", err)
	}
	return &TokenVerifier{publicKey: publicKey}, nil
}
