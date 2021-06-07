package itunes

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eSlider/itunes-ranking-service/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Service client to use for sync and manage local data
// and the place to contain project business-logic
type Service struct {
}

// Share one connection between all service instances (singleton)
var connection *gorm.DB

// GetConnection to database
func (s *Service) GetConnection() (*gorm.DB, error) {
	var err error
	if s.HasCacheDbConnection() {
		connection, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	}
	return connection, err
}

// HasCacheDbConnection bzw. is there an database?
func (s *Service) HasCacheDbConnection() bool {
	return connection == nil
}

// Load country specified entries
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

	for position, entry := range pc.Feed.Entries {
		entry.Position = position + 1
		entry.Country = country
	}

	return pc, err
}

// Update country specified top entries from official sources
// Clean's entries table before update
// TODO: improve with incremental update?
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

		for _, entry := range pc.Feed.Entries {
			db.Create(&entry)
		}

		podcasts[country] = pc
	}
	return podcasts, nil
}

// GetRankedEntriesByITuneId ordered by land as an `map[(string)land-string](uint)ranking-position-number
// Example: {es:2}
func (s *Service) GetRankedEntriesByITuneId(id uint64) (*[]EntryRankedResult, error) {
	var db, err = s.GetConnection()
	if err != nil {
		return nil, err
	}

	// Get results from db
	var entries []EntryRankedResult
	db.Table("entries").Where("i_tune_id = ?", []uint64{id}).Find(&entries)
	return &entries, err
}

// GroupEntriesByCountry as result
func (s *Service) GroupEntriesByCountry(entries *[]EntryRankedResult) *map[string]uint {
	var countryResult = &map[string]uint{}
	for _, entry := range *entries {
		(*countryResult)[entry.Country] = entry.Position
	}
	return countryResult
}
