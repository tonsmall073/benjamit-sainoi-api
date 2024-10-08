package utils

import "time"

func ConvDateStringToTimeType(datetime string) (time.Time, error) {
	layout := time.RFC3339 // รูปแบบวันที่
	date, dateErr := time.Parse(layout, datetime)
	if dateErr != nil {
		return time.Time{}, dateErr
	}
	return date, nil
}
