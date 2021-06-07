package itunes

import "database/sql/driver"

type Category struct {
	Id
	// unnecessary properties that are transported to Id type:
	//Attributes struct {
	//	LabelContainer
	//	Term   string `json:"term,omitempty"`
	//	Scheme string `json:"scheme,omitempty"`
	//} `json:"attributes,omitempty"`
}

func (c Category) Value() (driver.Value, error) {
	return c.Id.Uint()
}
