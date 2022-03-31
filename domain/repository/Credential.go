package repository

import "github.com/Aranyak-Ghosh/go-sample-clinic/infrastructure"

type credentialRepository struct {
}

type CredentialRespository interface {
	AddUserCredentials(infrastructure.CredentialModel) error
	FetchUserCredentialsById(string) (infrastructure.CredentialModel, error)
	FetchUserCredentialsByEmail(string) (infrastructure.CredentialModel, error)
}
