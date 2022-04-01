package repository

import (
	"github.com/Aranyak-Ghosh/go-sample-clinic/infrastructure/datamodel"
	"github.com/jmoiron/sqlx"
)

type PatientRepository interface {
	RegisterUserAsPatient(datamodel.PatientModel) error
	UpdatePatientProfile(datamodel.PatientModel) error
	GetPatientProfileByUserId(string) (*datamodel.PatientModel, error)
}

type patientRepository struct {
	db *sqlx.DB
}

var _ PatientRepository = (*patientRepository)(nil)

func (pr *patientRepository) RegisterUserAsPatient(datamodel.PatientModel) error {
	return nil
}

func (pr *patientRepository) UpdatePatientProfile(datamodel.PatientModel) error {
	return nil
}

func (pr *patientRepository) GetPatientProfileByUserId(string) (*datamodel.PatientModel, error) {
	return nil, nil
}
