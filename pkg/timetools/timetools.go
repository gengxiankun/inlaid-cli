package timetools

import (
	"time"
)

func TraverseTime(start string, end string) []string {
	layout := "2006-01-02"

	ended, _ := time.ParseInLocation(layout, end, time.Local)
	started, _ := time.ParseInLocation(layout, start, time.Local)

	dateArray := make([]string, 0)

	for ended.After(started) {
		dateArray = append(dateArray, started.Format("2006-01-02"))
		started = started.AddDate(0, 0, 1)
	}
	dateArray = append(dateArray, end)

	return dateArray
}
