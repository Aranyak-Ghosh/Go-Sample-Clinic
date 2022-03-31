package datamodel

type CredentialModel struct {
	Id           string `db:"id"`
	Email        string `db:"email"`
	PasswordHash string `db:"hash_password"`
}
