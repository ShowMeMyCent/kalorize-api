package utils

import "time"

func StringToDate(date string) (time.Time, error) {
	layout := "2006-01-02"
	return time.Parse(layout, date)
}
