package itunes

import "database/sql/driver"

type ContentType struct {
	Attributes struct {
		Term  string `json:"term,omitempty"`
		Label string `json:"label,omitempty"`
	} `json:"attributes,omitempty"`
}

func (c ContentType) Value() (driver.Value, error) {
	return c.Attributes.Term, nil
}

func (c ContentType) IsPodcast() bool {
	return c.Attributes.Term == "Podcast"
}
