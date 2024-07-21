package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func GenerateEd25519KeyPair() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	return ed25519.GenerateKey(rand.Reader)
}

func ParsePemEd25519PrivateKey(data []byte) (ed25519.PrivateKey, error) {
	block, _ := pem.Decode(data)
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	typed, ok := key.(ed25519.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("wrong key type")
	}
	return typed, nil
}

func ParsePemEd25519PublicKey(data []byte) (ed25519.PublicKey, error) {
	block, _ := pem.Decode(data)
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	typed, ok := key.(ed25519.PublicKey)
	if !ok {
		return nil, fmt.Errorf("wrong key type")
	}
	return typed, nil
}
