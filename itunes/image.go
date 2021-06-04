package itunes

type Image struct {
	LabelContainer
	Attributes struct {
		Height string `json:"height,omitempty"`
	} `json:"attributes,omitempty"`
}
