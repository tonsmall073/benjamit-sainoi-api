package utils

import (
	"os"
	"time"
)

func DatetimeZone(datetime time.Time) time.Time {
	defaultStr := "Asia/Bangkok"
	if getTimeZone := os.Getenv("SERVER_TIME_ZONE"); getTimeZone != "" {
		defaultStr = getTimeZone
	}
	loc, err := time.LoadLocation(defaultStr)
	if err != nil {
		return datetime
	}
	return datetime.In(loc)

}
