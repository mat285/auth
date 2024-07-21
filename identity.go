package auth

type Identity struct {
	Identifier string
	Creds      StoredCreds
	Data       interface{}
}

func NewIdentity(id string, creds StoredCreds, data interface{}) Identity {
	return Identity{
		Identifier: id,
		Creds:      creds,
		Data:       data,
	}
}
