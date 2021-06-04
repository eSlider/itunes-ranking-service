package itunes

type Author struct {
	Name LabelContainer `json:"name,omitempty"`
	URI  LabelContainer `json:"uri,omitempty"`
}

// GetUri of the Author
func (a *Author) GetUri() string {
	return a.URI.Label
}

// String representation
func (a *Author) String() string {
	return a.Name.String()
}
