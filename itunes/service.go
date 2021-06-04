package itunes

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eSlider/itunes-ranking-service/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Service struct {
}

// GetConnection to database
func (s *Service) GetConnection() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("data/podcast.db"), &gorm.Config{})
}

// Load country entries
func (s *Service) Load(country Country, limit int) (*Podcast, error) {
	var url = fmt.Sprintf("https://itunes.apple.com/%s/rss/topaudiopodcasts/limit=%d/json", country, limit)
	var c, err = utils.LoadOverHttp(url)
	var pc = &Podcast{}

	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, errors.New(fmt.Sprintf("Feed from '%s' is empty", url))
	}

	err = json.Unmarshal(*c, pc)

	if err != nil {
		return nil, err
	}

	var position = 0
	for _, entry := range pc.Feed.Entries {
		position++
		entry.Position = position
		entry.Country = country
	}

	return pc, err
}

// Update country specified top entries
func (s *Service) Update() (map[Country]*Podcast, error) {
	var limit = 100
	var db, err = s.GetConnection()
	//var path = "data/%s.json"

	if err != nil {
		return nil, err
	}

	var entry = &Entry{}

	// Migrate the schema
	err = db.AutoMigrate(entry)

	// Clean up
	db.Exec("DELETE FROM entries")

	var podcasts = map[Country]*Podcast{}

	// Get over all countries
	for _, country := range Countries {
		var pc, err = s.Load(country, limit)

		if err != nil {
			return nil, err
		}

		var position = 0
		for _, entry := range pc.Feed.Entries {
			position++
			entry.Position = position
			db.Create(&entry)
		}

		podcasts[country] = pc
	}
	return podcasts, nil
}

func (s *Service) GetById(id uint64) (*map[string]uint, error) {
	var db, err = s.GetConnection()

	if err != nil {
		return nil, err
	}

	// Get results from db
	var entries []RankResult
	db.Table("entries").Where("i_tune_id = ?", []uint64{id}).Find(&entries)

	// Group results by country
	var countryResult = map[string]uint{}
	for _, entry := range entries {
		countryResult[entry.Country] = entry.Position
	}

	return &countryResult, err
}
