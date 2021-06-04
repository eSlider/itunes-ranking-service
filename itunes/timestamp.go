package itunes

import (
	"database/sql/driver"
	"log"
	"time"
)

type TimeStamp LabelContainer

func (t TimeStamp) GetDate() time.Time {
	var d, err = time.Parse("2006-01-02T15:04:05-07:00", t.Label)
	if err != nil {
		log.Fatal(err.Error())
	}
	return d
}

func (t TimeStamp) Value() (driver.Value, error) {
	return t.Label, nil
}
