package model

import (
	"strings"
	"time"
)

type DateTime struct {
	StringDate string
	StringTime string
}

func (dt *DateTime) DateTimeFormat() string {
	const (
		layoutDateTime = "2006-01-0215:04PM"
	)
	ist, _ := time.LoadLocation("Asia/Kolkata")

	str := dt.StringDate + dt.StringTime

	str = strings.ReplaceAll(str, " ", "")

	dateTime, _ := time.ParseInLocation(layoutDateTime, str, ist)

	return dateTime.Format(time.RFC3339Nano)

}

func ConvertDateTime(input_date, input_time string) string {
	dateTime := new(DateTime)
	dateTime.StringDate = input_date
	dateTime.StringTime = input_time
	return dateTime.DateTimeFormat()
}
