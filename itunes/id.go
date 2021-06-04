package itunes

import (
	"database/sql/driver"
	"log"
	"strconv"
)

type Id struct {
	LabelContainer
	Attributes struct {
		LabelContainer
		Id     string `json:"im:id,omitempty"`
		Term   string `json:"term,omitempty"`
		Scheme string `json:"scheme,omitempty"`
	} `json:"attributes,omitempty"`
}

// Value Implement Valuer
func (i Id) Value() (driver.Value, error) {
	return i.Uint(), nil
}

func (i Id) Uint() int64 {
	if i.Attributes.Id == "" {
		return 0
	}

	id, err := strconv.ParseInt(i.Attributes.Id, 10, 32)
	if err != nil {
		log.Printf(err.Error())
	}
	return id
}
