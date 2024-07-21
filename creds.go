package auth

type Creds struct {
	Identifier string
	Password   []byte
}

type StoredCreds struct {
	PasswordHash []byte
	PasswordSalt []byte
}
