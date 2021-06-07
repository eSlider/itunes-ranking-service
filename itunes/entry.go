package itunes

import "github.com/eSlider/itunes-ranking-service/itunes/link"

type Entry struct {
	ID uint `gorm:"primaryKey"`
	//UpdatedAt int64  `gorm:"autoUpdateTime"` // Use unix seconds as creating time
	//CreatedAt int64  `gorm:"autoCreateTime"` // Use unix seconds as creating time

	ITuneId Id              `json:"id,omitempty"`
	Title   *LabelContainer `json:"title,omitempty" gorm:"-"`
	Summary *LabelContainer `json:"summary,omitempty" gorm:"-"`
	Name    *LabelContainer `json:"im:name,omitempty" gorm:"-"`
	Rights  *LabelContainer `json:"rights,omitempty" gorm:"-"`
	Artist  *Artist         `json:"im:artist,omitempty" gorm:"-"`
	Price   *Price          `json:"im:price,omitempty" gorm:"-"`
	Link    *link.Link      `json:"link,omitempty" gorm:"-"`

	ReleasedAt  *TimeStamp   `json:"im:releaseDate,omitempty" gorm:"-"`
	Category    *Category    `json:"category,omitempty" gorm:"-"`
	ContentType *ContentType `json:"im:contentType,omitempty" gorm:"-"`

	Images []Image `json:"im:image,omitempty" gorm:"-"`

	Country Country

	Position int
}

//
//func (e *Entry) IsPodcast() bool {
//	return e.ContentType.IsPodcast()
//}
