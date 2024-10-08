package utils

import "time"

func ConvTime(timestamp string) time.Time {
	datetime, _ := ConvDateStringToTimeType(timestamp)
	return datetime
}
