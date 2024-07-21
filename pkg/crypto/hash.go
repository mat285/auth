package crypto

import (
	"crypto/rand"
	"crypto/sha256"
)

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 256)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func Hash(pass, salt []byte) []byte {
	r := append([]byte{}, pass...)
	r = append(r, salt...)
	return sha256.New().Sum(r)
}
