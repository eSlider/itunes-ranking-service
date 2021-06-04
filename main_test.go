package main

import (
	"encoding/json"
	"github.com/eSlider/itunes-ranking-service/itunes"
	"github.com/eSlider/itunes-ranking-service/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestLoadAndUnmarshalPodcastFeed(t *testing.T) {
	//var country = itunes.USA
	var limit = 100
	var pc = &itunes.Podcast{}

	for _, country := range itunes.Countries {
		var err = pc.Load(country, limit)
		if err != nil {
			t.Fatalf(err.Error())
		}

		t.Logf("Podcast for '%s' has %d entries", country, len(pc.Feed.Entries))
	}
}

func TestUnmarshalPodcast(t *testing.T) {
	var path = "data/list.json"
	var c, err = utils.ReadFile(path)
	var pc = &itunes.Podcast{}

	if c == nil {
		t.Fatalf("No content in '%s'", path)
	}

	err = json.Unmarshal(*c, pc)

	if err != nil {
		t.Fatalf(err.Error())
	}

	db, err := gorm.Open(sqlite.Open("data/test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	//var entryBase = &itunes.Entry{}
	var entry = pc.Feed.Entries[0]

	// Migrate the schema
	err = db.AutoMigrate(&entry)

	// Clean up
	db.Exec("DELETE FROM entries")

	if err != nil {
		panic(err.Error())
	}

	// Save to database
	for _, entry = range pc.Feed.Entries {
		//entry.ID = entry.ITuneId.Uint()
		entry.Country = itunes.Deutschland
		db.Create(&entry)
	}
}

// http://localhost:8080/rank/1450994021 {"es":98,"it":46,"us":10}
