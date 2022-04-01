package datamodel

import "time"

type AppointmentModel struct {
	Id             string        `db:"id"`
	DoctorId       string        `db:"doctor_id"`
	PatientId      string        `db:"patient_id"`
	StartDateTime  time.Time     `db:"start_date_time"`
	Duration       time.Duration `db:"duration"`
	Summary        string        `db:"summary"`
	RequestMessage string        `db:"request_message"`
}
