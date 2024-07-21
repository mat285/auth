package crypto

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func PemEncodeKeyPair(pub, priv interface{}) (pubPem, privPem *pem.Block, err error) {
	privPem, err = PemEncodePrivateKey(priv)
	if err != nil {
		return
	}
	pubPem, err = PemEncodePublicKeyKey(pub)
	return
}

func PemEncodePrivateKey(priv interface{}) (*pem.Block, error) {
	b, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return nil, err
	}

	return &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: b,
	}, nil
}

func PemEncodePublicKeyKey(pub interface{}) (*pem.Block, error) {
	b, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return nil, err
	}

	return &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: b,
	}, nil
}

func WritePemFilePair(privFile string, pub, priv *pem.Block) error {
	if len(privFile) == 0 {
		return fmt.Errorf("no file specified")
	}
	pubFile := fmt.Sprintf("%s.pub", privFile)

	err := os.WriteFile(privFile, pem.EncodeToMemory(priv), 0600)
	if err != nil {
		return err
	}
	return os.WriteFile(pubFile, pem.EncodeToMemory(pub), 0600)
}
