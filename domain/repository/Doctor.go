package repository

import (
	"github.com/Aranyak-Ghosh/go-sample-clinic/infrastructure/datamodel"
	"github.com/jmoiron/sqlx"
)

type DoctorRepository interface {
	RegisterUserAsDoctor(datamodel.DoctorModel) error
	RemoveDoctor(UserId string) error
	UpdateDoctorProfile(datamodel.DoctorModel) error
	GetDoctorProfileByUserId(UserId string) (*datamodel.DoctorModel, error)
}

type doctorRepository struct {
	db *sqlx.DB
}

var _ DoctorRepository = (*doctorRepository)(nil)

func (dr *doctorRepository) RegisterUserAsDoctor(datamodel.DoctorModel) error {
	return nil
}

func (dr *doctorRepository) RemoveDoctor(UserId string) error {
	return nil
}

func (dr *doctorRepository) UpdateDoctorProfile(datamodel.DoctorModel) error {
	return nil
}

func (dr *doctorRepository) GetDoctorProfileByUserId(UserId string) (*datamodel.DoctorModel, error) {
	return nil, nil
}
