package itunes

import "github.com/eSlider/itunes-ranking-service/itunes/link"

type Entry struct {
	ID uint `gorm:"primaryKey"`
	//UpdatedAt int64  `gorm:"autoUpdateTime"` // Use unix seconds as creating time
	//CreatedAt int64  `gorm:"autoCreateTime"` // Use unix seconds as creating time

	ITuneId Id              `json:"id,omitempty"`
	Title   *LabelContainer `json:"title,omitempty"`
	Summary *LabelContainer `json:"summary,omitempty"`
	Name    *LabelContainer `json:"im:name,omitempty"`
	Rights  *LabelContainer `json:"rights,omitempty"`
	Artist  *Artist         `json:"im:artist,omitempty"`
	Price   *Price          `json:"im:price,omitempty"`
	Link    *link.Link      `json:"link,omitempty"`

	ReleasedAt  *TimeStamp   `json:"im:releaseDate,omitempty"`
	Category    *Category    `json:"category,omitempty"`
	ContentType *ContentType `json:"im:contentType,omitempty"`

	Images []Image `json:"im:image,omitempty" gorm:"-"`

	Country Country

	Position int
}

//
//func (e *Entry) IsPodcast() bool {
//	return e.ContentType.IsPodcast()
//}
