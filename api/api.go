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

	w.WriteHeader(http.StatusOK)
}

// GetRankByITuneId and return as an list entries in each country
func GetRankByITuneId(w http.ResponseWriter, r *http.Request) {
	var pathPars = strings.Split(r.RequestURI, "/")

	// Has ITuneId in the path?
	if len(pathPars) < 3 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
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

	// Check if ID isn't
	if id < 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("[]"))
		return
	}

	entries, err := service.GetById(id)
	if err != nil {
		handleError(w, err)
		return
	}

	js, err := json.Marshal(entries)

	if err != nil {
		handleError(w, err)
		return
	}

	_, err = w.Write(js)

	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
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
