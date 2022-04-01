package repository

import (
	"github.com/Aranyak-Ghosh/go-sample-clinic/infrastructure/datamodel"
	"github.com/jmoiron/sqlx"
)

type credentialRepository struct {
	db *sqlx.DB
}

var _ CredentialRespository = (*credentialRepository)(nil)

type CredentialRespository interface {
	AddUserCredentials(datamodel.CredentialModel) error
	FetchUserCredentialsById(string) (*datamodel.CredentialModel, error)
	FetchUserCredentialsByEmail(string) (*datamodel.CredentialModel, error)
}

func (cr *credentialRepository) AddUserCredentials(model datamodel.CredentialModel) error {
	return nil
}

func (cr *credentialRepository) FetchUserCredentialsById(id string) (*datamodel.CredentialModel, error) {
	return nil, nil
}

func (cr *credentialRepository) FetchUserCredentialsByEmail(email string) (*datamodel.CredentialModel, error) {
	return nil, nil
}
