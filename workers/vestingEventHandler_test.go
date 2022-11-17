package workers

import (
	"context"
	"github.com/nicknikandish/vesting_program/models"
	"io"
	"os"
	"sync"
	"testing"
	"time"
)

func TestReadCSVFile(t *testing.T) {
	type args struct {
		filepath   string
		precision  int
		targetDate time.Time
	}
	filename := "example1.csv"
	pwd, _ := os.Getwd()
	path := pwd + "/../" + filename
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test ReadCSVFile",
			args: args{
				filepath:   path,
				precision:  2,
				targetDate: time.Now(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReadCSVFile(tt.args.filepath, tt.args.precision, tt.args.targetDate); (err != nil) != tt.wantErr {
				t.Errorf("ReadCSVFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_declareWorker(t *testing.T) {
	type args struct {
		ctx          context.Context
		wg           *sync.WaitGroup
		src          chan []models.VestingEvent
		targetDate   time.Time
		precision    int
		numberOfCPUs int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test declareWorker",
			args: args{
				ctx:          context.Background(),
				wg:           &sync.WaitGroup{},
				src:          make(chan []models.VestingEvent),
				targetDate:   time.Now(),
				precision:    2,
				numberOfCPUs: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			declareWorker(tt.args.ctx, tt.args.wg, tt.args.src, tt.args.targetDate, tt.args.precision, tt.args.numberOfCPUs)
		})
	}
}

func Test_lineCounter(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Test lineCounter",
			args: args{
				r: os.Stdin,
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lineCounter(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("lineCounter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("lineCounter() got = %v, want %v", got, tt.want)
			}
		})
	}
}
