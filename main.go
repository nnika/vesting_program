package main

import (
	"fmt"
	"github.com/nicknikandish/vesting_program/workers"
	"os"
	"strconv"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) < 3 {
		fmt.Println("Usage: filename date [precision]")
		return
	}
	filename := os.Args[1]
	pwd, _ := os.Getwd()
	path := pwd + "/" + filename

	date, err := time.Parse("2006-01-02", os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	precision := 0
	if len(arguments) > 3 {
		precision, err = strconv.Atoi(arguments[3])
		if err != nil {
			fmt.Println(err)
		}
	}

	validate(precision)
	initialize(path, date, precision)
}

func initialize(filePath string, targetDate time.Time, precision int) {

	err := workers.ReadCSVFile(filePath, precision, targetDate)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	workers.PrintList(precision)

}
func validate(precision int) {
	if precision < 0 || precision > 6 {
		panic("Precision is supposed to be between 0 and 6")
	}

}
