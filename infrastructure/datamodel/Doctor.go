package datamodel

import "time"

type DoctorModel struct {
	UserId         string    `db:"user_id"`
	Specialization string    `db:"specialization"`
	StartTime      time.Time `db:"start_time"`
	EndTime        time.Time `db:"end_time"`
}
