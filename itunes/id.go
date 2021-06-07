package itunes

import (
	"database/sql/driver"
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
	return i.Uint()
}

// Uint ID as int64
func (i Id) Uint() (int64, error) {
	if i.Attributes.Id == "" {
		return 0, nil
	}
	return strconv.ParseInt(i.Attributes.Id, 10, 32)
}
