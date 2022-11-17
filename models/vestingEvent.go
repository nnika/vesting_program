package models

import (
	"github.com/nicknikandish/vesting_program/api"
	"time"
)

type VestingEvent api.VestingEvent

func (v *VestingEvent) TableName() string {
	return "vesting_events"
}

func (v *VestingEvent) GetVestType() string {
	return v.VestType
}

func (v *VestingEvent) SetVestType(vestType string) {
	v.VestType = vestType
}

func (v *VestingEvent) GetAward() api.Award {
	return v.Award
}

func (v *VestingEvent) SetAward(award api.Award) {
	v.Award = award
}

func (v *VestingEvent) GetDate() time.Time {
	return v.Date
}

func (v *VestingEvent) SetDate(date time.Time) {
	v.Date = date
}

func (v *VestingEvent) GetQuantity() float64 {
	return v.Quantity
}

func (v *VestingEvent) SetQuantity(quantity float64) {
	v.Quantity = quantity
}

func VestingEventReader(vestType string, awardId string, employeeName string, date time.Time, quantity float64, employeeId string) *VestingEvent {
	return &VestingEvent{
		VestType: vestType,
		Award: api.Award{
			EmployeeId:   employeeId,
			EmployeeName: employeeName,
			AwardId:      awardId,
		},
		Date:     date,
		Quantity: quantity,
	}
}
