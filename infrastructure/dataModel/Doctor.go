package datamodel

import "time"

type DoctorModel struct {
	UserId         string
	Specialization string
	StartTime      time.Time
	EndTime        time.Time
}
