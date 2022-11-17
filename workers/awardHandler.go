package workers

import (
	"fmt"
	"github.com/nicknikandish/vesting_program/models"
	"math"
	"sort"
	"sync"
	"time"
)

var awards = make(map[models.Award]float64)
var mu sync.Mutex

func Process(eventsList []models.VestingEvent, targetDate time.Time, precision int) {
	for _, event := range eventsList {
		if !event.GetDate().After(targetDate) {
			if _, ok := awards[models.Award(event.GetAward())]; ok {
				award := event.GetAward()
				quantity := event.GetQuantity()
				if event.GetVestType() == "VEST" {
					awards[models.Award(award)] += RoundFloat(quantity, precision)
				} else if event.GetVestType() == "CANCEL" {
					awards[models.Award(award)] -= RoundFloat(quantity, precision)
				}
			} else {
				awards[models.Award(event.GetAward())] = RoundFloat(event.GetQuantity(), precision)
			}
		} else {
			if _, ok := awards[models.Award(event.GetAward())]; !ok {
				awards[models.Award(event.GetAward())] = 0
			}
		}
	}

}

func PrintList(precision int) {
	keys := make([]models.Award, 0, len(awards))

	for key := range awards {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i].EmployeeId < keys[j].EmployeeId
	})
	for i, k := range keys {
		s := fmt.Sprintf("%s, %s,%s, %.[5]*[4]f", k.EmployeeId, k.EmployeeName, keys[i].GetAwardId(), awards[k], precision)
		fmt.Println(s)
	}
}
func RoundFloat(val float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Floor(val*ratio) / ratio
}
