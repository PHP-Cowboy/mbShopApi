package timeutil

import (
	"fmt"
	"time"
)

type JsonTime time.Time

const (
	MonthFormat     = "2006-01"
	DateFormat      = "2006-01-02"
	TimeFormat      = "2006-01-02 15:04:05"
	TimeFormNoSplit = "20060102150405"
)

func (t JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format(TimeFormat))
	return []byte(stamp), nil
}

func (t *JsonTime) UnmarshalJSON(data []byte) error {
	now, err := time.ParseInLocation(TimeFormat, string(data), time.Local)
	*t = JsonTime(now)
	return err
}
