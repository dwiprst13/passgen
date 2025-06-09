package model

type Credential struct {
	Service           string `json:"service"`            
	Username          string `json:"username"`           
	EncryptedPassword string `json:"encrypted_password"` 
}

type CredentialStore struct{
	Credentials []Credential `json:"credentials"` 
}