package api

import (
	"time"
)

type Award struct {
	EmployeeId   string `db:"employee_id" json:"employee_id"`
	EmployeeName string `db:"employee_name" json:"employee_name"`
	AwardId      string `db:"award_id" json:"award_id"`
}

type VestingEvent struct {
	VestType string    `json:"vest_type" db:"vest_type"`
	Award    Award     `db:"award" json:"award"`
	Date     time.Time `db:"date" json:"date"`
	Quantity float64   `db:"quantity" json:"quantity"`
}
