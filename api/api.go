package api

import (
	"github.com/eSlider/itunes-ranking-service/itunes"
	"net/http"
	"strconv"
	"strings"
)

var service = &itunes.Service{}

type controlMethodRegister map[string]func(p *PathMethod, r *http.Request) interface{}

var methods = controlMethodRegister{

	// GetUpdates database by getting land specified top 100 entries
	"GetUpdate": func(p *PathMethod, r *http.Request) interface{} {
		var podcasts, err = service.Update()
		if err != nil {
			return err
		}
		return podcasts
	},

	// GetRankByITuneId and return as an list entries in each country
	"GetRankiTuneId": func(p *PathMethod, r *http.Request) interface{} {
		var pathPars = strings.Split(r.RequestURI, "/")

		// Has ITuneId in the path?
		if len(pathPars) < 3 {
			return []string{}
		}

		// Check if ID isn't number
		var id, err = strconv.ParseUint(pathPars[2], 10, 32)
		if err != nil {
			return err
		}

		// Check if ID isn't:
		// No need to communicate database
		if id < 0 {
			return []string{}
		}

		// Sync before rank
		if !service.HasCacheDbConnection() {
			_, err := service.Update()
			if err != nil {
				return err
			}
		}

		entries, err := service.GetRankedEntriesByITuneId(id)

		if err != nil {
			return err
		}

		return service.GroupEntriesByCountry(entries)
	},
}
