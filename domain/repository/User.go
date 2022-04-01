package repository

import (
	"github.com/Aranyak-Ghosh/go-sample-clinic/infrastructure/datamodel"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	RegisterUser(datamodel.UserModel) error
	FetchUserById(string) (*datamodel.UserModel, error)
}

type userRepository struct {
	db *sqlx.DB
}

var _ UserRepository = (*userRepository)(nil)

func (ur *userRepository) RegisterUser(datamodel.UserModel) error {
	return nil
}

func (ur *userRepository) FetchUserById(string) (*datamodel.UserModel, error) {
	return nil, nil
}
