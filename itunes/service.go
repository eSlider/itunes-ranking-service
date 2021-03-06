package itunes

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eSlider/itunes-ranking-service/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/url"
	"sync"
)

func NewService() *Service {
	s := new(Service)
	return s
}

// Service client to use for sync and manage local data
// and the place to contain project business-logic
type Service struct {
	connection *gorm.DB
}

// GetConnection to database
func (s *Service) GetConnection() (*gorm.DB, error) {
	var err error
	if !s.HasCacheDbConnection() {

		var params = url.Values{}

		// Prevents any timeouts
		params.Add("_timeout", "0")

		// If shared-cache mode is enabled and a thread establishes multiple connections to the same database,
		// the connections share a single data and schema cache.
		// This can significantly reduce the quantity of memory and IO required by the system.
		params.Add("cache", "shared")

		// With synchronous OFF (0), SQLite continues without syncing as soon as
		// it has handed data off to the operating system.
		// If the application running SQLite crashes, the data will be safe,
		// but the database might become corrupted if the operating system crashes
		// or the computer loses power before that data has been written to the disk surface.
		// On the other hand, commits can be orders of magnitude faster with synchronous OFF.
		params.Add("_sync", "OFF")

		// No need to optimize database storage every time it's changes
		params.Add("_vacuum", "0")

		//params.Add("journal", "OFF")

		s.connection, err = gorm.Open(sqlite.Open(":memory:?"+params.Encode()), &gorm.Config{})
	}
	return s.connection, err
}

// HasCacheDbConnection bzw. is there an database?
func (s *Service) HasCacheDbConnection() bool {
	return s.connection != nil
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
func (s *Service) Update() (*map[Country]*Podcast, error) {
	var limit = 100
	var db, err = s.GetConnection()
	//var path = "data/%s.json"

	if err != nil {
		return nil, err
	}

	// Load new entries
	podcasts, _errors := s.LoadAll(limit)

	// Handle errors
	// TODO: implement global error handler
	if len(_errors) > 0 {
		jsb, _ := json.Marshal(_errors)
		return podcasts, errors.New(string(jsb))
	}

	// Start transaction
	var tx = db.Begin()

	// Migrate the schema
	var entry = &Entry{}
	err = tx.AutoMigrate(entry)

	if err != nil {
		return nil, err
	}

	// Clean up
	s.DeleteAll(tx)

	// Update podcasts
	for _, podcast := range *podcasts {
		for _, entry := range podcast.Feed.Entries {
			tx.Create(&entry)
		}
	}

	// Commit transaction
	tx.Commit()

	return podcasts, nil
}

// LoadAll podcasts with a wait group
func (s *Service) LoadAll(limit int) (*map[Country]*Podcast, []error) {
	var podcasts = map[Country]*Podcast{}
	var wg sync.WaitGroup
	var _errors []error

	// Load async over all countries
	for _, country := range Countries {
		wg.Add(1)
		go func(country Country) {
			defer wg.Done()
			var pc, err = s.Load(country, limit)
			if err != nil {
				_errors = append(_errors, err)
				return
			}
			podcasts[country] = pc
			log.Printf("Loaded %s country with %d entries\n", country, len(pc.Feed.Entries))
		}(country)
	}

	// Wait for all
	wg.Wait()

	return &podcasts, _errors
}

// DeleteAll entries
func (s *Service) DeleteAll(db *gorm.DB) *gorm.DB {
	return db.Exec("DELETE FROM entries")
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
