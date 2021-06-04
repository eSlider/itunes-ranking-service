package itunes

import (
	"github.com/eSlider/itunes-ranking-service/itunes/link"
)

type Feed struct {
	ID      *LabelContainer `json:"id,omitempty"`
	Title   *LabelContainer `json:"title,omitempty"`
	Author  *Author         `json:"author,omitempty"`
	Entries []*Entry        `json:"entry,omitempty"`
	Icon    *LabelContainer `json:"icon,omitempty"`
	Link    []link.Link     `json:"link,omitempty"`
	Rights  *LabelContainer `json:"rights,omitempty"`
	Updated *TimeStamp      `json:"updated,omitempty"`
}
