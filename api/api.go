package api

import (
	"encoding/json"
	"github.com/eSlider/itunes-ranking-service/itunes"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var service = &itunes.Service{}

// GetUpdate database by getting land specified top 100 entries
func GetUpdate(w http.ResponseWriter, r *http.Request) {
	var podcasts, err = service.Update()

	if err != nil {
		handleError(w, err)
		return
	}

	// TODO: return better statics
	js, err := json.Marshal(podcasts)

	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_, err = w.Write(js)

	if err != nil {
		handleError(w, err)
		return
	}
}

// GetRankByITuneId and return as an list entries in each country
func GetRankByITuneId(w http.ResponseWriter, r *http.Request) {
	var headers = w.Header()
	for k, v := range map[string]string{
		"Content-Type":  "application/json; charset=UTF-8",
		"Cache-Control": "no-store, max-age=0",
	} {
		headers.Set(k, v)
	}

	var pathPars = strings.Split(r.RequestURI, "/")
	// Has ITuneId in the path?
	if len(pathPars) < 3 {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("[]"))
		return
	}

	// Check if ID isn't number
	var id, err = strconv.ParseUint(pathPars[2], 10, 32)
	if err != nil {
		handleError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	// Check if ID isn't:
	// No need to communicate database
	if id < 0 {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("[]"))
		if err != nil {
			handleError(w, err)
			return
		}
		return
	}

	// Sync before rank
	if !service.HasCacheDbConnection() {
		_, err := service.Update()
		if err != nil {
			handleError(w, err)
			return
		}
	}

	entries, err := service.GetRankedEntriesByITuneId(id)
	var result = service.GroupEntriesByCountry(entries)

	if err != nil {
		handleError(w, err)
		return
	}

	js, err := json.Marshal(result)

	if err != nil {
		handleError(w, err)
		return
	}

	_, err = w.Write(js)

	if err != nil {
		handleError(w, err)
		return
	}
}

// Handle error
func handleError(w http.ResponseWriter, err error) {
	log.Printf(err.Error())
	_, err = w.Write([]byte(err.Error()))
	if err != nil {
		log.Printf(err.Error())
		w.Write([]byte("," + err.Error()))
	}
	w.WriteHeader(http.StatusInternalServerError)
}
