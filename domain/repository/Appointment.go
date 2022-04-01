package repository

import (
	"time"

	"github.com/Aranyak-Ghosh/go-sample-clinic/infrastructure/datamodel"
	"github.com/jmoiron/sqlx"
)

type AppointmentRepository interface {
	CreateAppointment(appointment datamodel.AppointmentModel) error
	UpdateAppointment(appointment datamodel.AppointmentModel) error
	DeleteAppointment(appointmentId string) error
	ListAppointmentsForDoctor(doctorId string, startDateTime, endDateTime time.Time) ([]datamodel.AppointmentModel, error)
	ListAppointmentsForPatient(patientId string, startDateTime, endDateTime time.Time) ([]datamodel.AppointmentModel, error)
	ListAppointments(startDateTime, endDateTime time.Time) ([]datamodel.AppointmentModel, error)
}

type appointmentRepository struct {
	db *sqlx.DB
}

var _ AppointmentRepository = (*appointmentRepository)(nil)

func (ar *appointmentRepository) CreateAppointment(appointment datamodel.AppointmentModel) error {
	return nil
}

func (ar *appointmentRepository) UpdateAppointment(appointment datamodel.AppointmentModel) error {
	return nil
}

func (ar *appointmentRepository) DeleteAppointment(appointmentId string) error {
	return nil
}

func (ar *appointmentRepository) ListAppointmentsForDoctor(doctorId string, startDateTime, endDateTime time.Time) ([]datamodel.AppointmentModel, error) {
	return nil, nil
}
func (ar *appointmentRepository) ListAppointmentsForPatient(patientId string, startDateTime, endDateTime time.Time) ([]datamodel.AppointmentModel, error) {
	return nil, nil
}
func (ar *appointmentRepository) ListAppointments(startDateTime, endDateTime time.Time) ([]datamodel.AppointmentModel, error) {
	return nil, nil
}
