package models

import "github.com/nicknikandish/vesting_program/api"

type Award api.Award

func (a *Award) TableName() string {
	return "awards"
}

func (a *Award) GetEmployeeId() string {
	return a.EmployeeId
}

func (a *Award) SetEmployeeId(employeeId string) {
	a.EmployeeId = employeeId
}

func (a *Award) GetEmployeeName() string {
	return a.EmployeeName
}

func (a *Award) SetEmployeeName(employeeName string) {
	a.EmployeeName = employeeName
}

func (a *Award) GetAwardId() string {
	return a.AwardId
}

func (a *Award) SetAwardId(awardId string) {
	a.AwardId = awardId
}
