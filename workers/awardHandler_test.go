package workers

import (
	"github.com/nicknikandish/vesting_program/api"
	"github.com/nicknikandish/vesting_program/models"
	"testing"
	"time"
)

func TestPrintList(t *testing.T) {
	type args struct {
		precision int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Print List",
			args: args{precision: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintList(tt.args.precision)
		})
	}
}

func TestProcess(t *testing.T) {
	type args struct {
		eventsList []models.VestingEvent
		targetDate time.Time
		precision  int
	}
	tests := []struct {
		name string
		args args
		want []models.Award
	}{
		{
			name: "Test Process",
			args: args{
				eventsList: []models.VestingEvent{
					{
						VestType: "RSU",
						Award: api.Award(models.Award{
							EmployeeId:   "1",
							EmployeeName: "Nick",
							AwardId:      "111",
						}),
						Date:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						Quantity: 100,
					},
				},
			},
			want: []models.Award{
				{
					EmployeeId:   "1",
					EmployeeName: "Nick",
					AwardId:      "111",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Process(tt.args.eventsList, tt.args.targetDate, tt.args.precision)
		})
	}
}

func TestRoundFloat(t *testing.T) {
	type args struct {
		val       float64
		precision int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test Round Float",
			args: args{
				val:       1.23456789,
				precision: 2,
			},
			want: 1.23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoundFloat(tt.args.val, tt.args.precision); got != tt.want {
				t.Errorf("RoundFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
