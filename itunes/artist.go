package itunes

import (
	"database/sql/driver"
	"github.com/eSlider/itunes-ranking-service/itunes/link"
)

type Artist struct {
	link.Link
	LabelContainer
}

func (l Artist) Value() (driver.Value, error) {
	return l.Label, nil
}
