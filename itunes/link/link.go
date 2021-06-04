package link

import (
	"database/sql/driver"
)

// Link entry
type Link struct {
	Attributes struct {

		// Link relation?
		Rel Relation `json:"rel,omitempty"`

		// Content mime type
		Type ContentType `json:"type,omitempty"`

		// URL/URI
		Href string `json:"href,omitempty"`
	} `json:"attributes,omitempty"`
}

// Value Implement Valuer
func (l Link) Value() (driver.Value, error) {
	return l.GetUri(), nil
}

// GetUri URL
func (l Link) GetUri() string {
	return l.Attributes.Href
}
