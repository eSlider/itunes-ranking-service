package itunes

import (
	"database/sql/driver"
	"time"
)

type TimeStamp LabelContainer

func (t TimeStamp) GetDate() (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05-07:00", t.Label)
}

func (t TimeStamp) Value() (driver.Value, error) {
	return t.Label, nil
}
