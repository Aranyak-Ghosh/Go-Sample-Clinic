package datamodel

import "time"

type PatientModel struct {
	UserId        string    `db:"user_id"`
	LastVisitedOn time.Time `db:"last_visted_on"`
}
