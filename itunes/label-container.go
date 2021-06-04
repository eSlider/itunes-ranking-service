package itunes

import "database/sql/driver"

type LabelContainer struct {
	Label string `json:"label,omitempty"`
}

func (l LabelContainer) Value() (driver.Value, error) {
	return l.Label, nil
}

// Representation as an string
// Related to https://stackoverflow.com/questions/13247644/tostring-function-in-go
func (l *LabelContainer) String() string {
	return l.Label
}
