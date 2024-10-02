package utils

import "time"

func ConvUnixtime(datetime string) (time.Time, error) {
	layout := time.RFC3339 // รูปแบบวันที่
	unixTime, unixTimeErr := time.Parse(layout, datetime)
	if unixTimeErr != nil {
		return time.Time{}, unixTimeErr
	}
	return unixTime, nil
}
