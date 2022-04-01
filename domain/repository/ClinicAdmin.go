package repository

import (
	"github.com/Aranyak-Ghosh/go-sample-clinic/infrastructure/datamodel"
	"github.com/jmoiron/sqlx"
)

type ClinicAdminRepository interface {
	IsUserClinicAdmin(UserId string) (bool, error)
	RegisterUserAsClinicAdmin(datamodel.ClinicAdminModel) error
	RemoveClinicAdmin(UserId string) error
}

type clinicAdminRepository struct {
	db *sqlx.DB
}

var _ ClinicAdminRepository = (*clinicAdminRepository)(nil)

func (cr *clinicAdminRepository) RegisterUserAsClinicAdmin(datamodel.ClinicAdminModel) error {
	return nil
}

func (cr *clinicAdminRepository) IsUserClinicAdmin(UserId string) (bool, error) {
	return false, nil
}

func (cr *clinicAdminRepository) RemoveClinicAdmin(UserId string) error {
	return nil
}
