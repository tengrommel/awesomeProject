package main

import (
	"fmt"
	"time"
)

func getDayStringFromTime(timeItem time.Time) string {
	return timeItem.Format("2006-01-02")
}

func main() {
	dateList := make([]string, 0)
	startTime := time.Now().Add(-48 * time.Hour)
	endTime := time.Now().Add(10 * time.Hour)
	fmt.Println(startTime.Before(endTime))
	for startTime := startTime; startTime.Before(endTime); startTime = startTime.Add(time.Hour * 24) {
		dateList = append(dateList, getDayStringFromTime(startTime))
	}
	dateList = append(dateList, endTime.Format("2006-01-02"))
	fmt.Println(dateList)
}
