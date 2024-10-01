package utils

import "time"

type Date time.Time

// Implement UnmarshalJSON สำหรับ custom type
func (d *Date) UnmarshalJSON(b []byte) error {
	// ลบเครื่องหมายอัญประกาศ
	str := string(b[1 : len(b)-1])
	parsedTime, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return err
	}
	*d = Date(parsedTime)
	return nil
}

func (d Date) Time() time.Time {
	return time.Time(d)
}
