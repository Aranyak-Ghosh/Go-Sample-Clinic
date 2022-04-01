package datamodel

import "time"

type AppointmentModel struct {
	DoctorId       string
	PatientId      string
	StartDateTime  time.Time
	Duration       time.Duration
	Summary        string
	RequestMessage string
}
