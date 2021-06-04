package itunes

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eSlider/itunes-ranking-service/utils"
)

type Podcast struct {
	Feed *Feed `json:"feed,omitempty"`
}

// Load top podcast by language and limit
// TODO: check and validate rss possible languages
// 			 and maybe method refactor to own "Podcasts" manager component
func (p *Podcast) Load(country Country, limit int) error {
	var url = fmt.Sprintf("https://itunes.apple.com/%s/rss/topaudiopodcasts/limit=%d/json", country, limit)
	var c, err = utils.LoadOverHttp(url)

	if err != nil {
		return err
	}

	if c == nil {
		return errors.New(fmt.Sprintf("Feed from '%s' is empty", url))
	}

	return json.Unmarshal(*c, p)
}
