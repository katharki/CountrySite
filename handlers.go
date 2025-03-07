package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

var starTime = time.Now()

func countryInfoHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("You have requested URL: %s\n", r.URL.Path)

	//extract country name from the URL
	countryCode := strings.TrimPrefix(r.URL.Path, "/countryinfo/v1/info/")
	countryCode = strings.TrimSpace(countryCode)

	if countryCode == "" {
		http.Error(w, "Missing country code", http.StatusBadRequest)
		return
	}

	if len(countryCode) != 2 {
		http.Error(w, "Invalid country code. Use a 2-letter ISO code, please", http.StatusBadRequest)
		return
	}

	countryCode = strings.ToUpper(countryCode)

	// fetch country info
	country, err := fetchCountryInfo(countryCode)
	if err != nil {
		log.Printf("ERROR: Failed to fetch country info for %s: %v", countryCode, err)
		http.Error(w, "Country not found:(", http.StatusNotFound)
		return
	}

	/*
		limit := 0

		queryLimit := r.URL.Query().Get("limit")
		if queryLimit != "" {
			if
		}

	*/

	//extract country code from the URL

	/*countryCode := r.URL.Path[len("/countryinfo/v1/info/"):]
	if countryCode == "" {
		http.Error(w, "Missing country code", http.StatusBadRequest)
		return
	}
	/*	*country, ok := countries[countryCode]

		if !ok {
			http.Error(w, "Coun	try not found", http.StatusNotFound)

			return
		}



	if r.URL.Path == "/countryinfo/v1/info/" {
		http.Error(w, "Missing country code", http.StatusBadRequest)
		return
	}
	*/

	/**	code := strings.TrimPrefix(r.URL.Path, "/countryinfo/v1/info/")
	if code == "" {
		http.Error(w, "Missing country code", http.StatusBadRequest)
		return
	}

	data, err := fetchCountryInfo(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}


	*/

	//sned json response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(country)

}

// populationHandler for requests

func populationHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("You have requested URL: %s\n", r.URL.Path)
	/*
		pathParts := strings.Split(r.URL.Path, "/")

		if len(pathParts) < 3 || pathParts[1] != "population" {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

	*/

	countryCode := strings.TrimPrefix(r.URL.Path, "/countryinfo/v1/population/")
	countryCode = strings.TrimSpace(countryCode)

	if countryCode == "" || len(countryCode) != 2 {
		http.Error(w, "Invalid country code. Use a 2-letter ISO code, please", http.StatusBadRequest)
		return
	}

	countryCode = strings.ToUpper(countryCode)

	//extract the limit query parameter
	startYear, endYear := 0, 0 //default return all years
	limitParameter := r.URL.Query().Get("limit")
	if limitParameter != "" {
		fmt.Sscanf(limitParameter, "%d-%d", &startYear, &endYear)
	}

	//fetch population data
	population, err := fetchPopulation(countryCode, startYear, endYear)
	if err != nil {
		log.Printf("ERROR: Failed ot fetch popualation data for %s: %v", countryCode, err)
		http.Error(w, "Could not retrieve population data", http.StatusNotFound)
		return
	}

	/*	countryName := pathParts[2]
		population, err := fetchPopulation(countryName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

	*/

	//send json response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(population)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Status check requested")

	countriesNowResp, _ := http.Get(countriesNowAPI)
	restCountriesResp, _ := http.Get(restCountriesAPI + "no")

	//send json response
	statusResponse := map[string]string{
		"status":           "ok",
		"timestamp":        time.Now().Format(time.RFC3339),
		"countriesNowAPI":  countriesNowResp.Status,
		"restCountriesAPI": restCountriesResp.Status,
		"uptime":           time.Since(starTime).String(),
	}

	if countriesNowResp != nil {
		statusResponse["countriesNowAPI"] = countriesNowResp.Status
	} else {
		statusResponse["countriesNowAPI"] = "error"
	}

	if restCountriesResp != nil {
		statusResponse["restCountriesAPI"] = restCountriesResp.Status
	} else {
		statusResponse["restCountriesAPI"] = "error"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(statusResponse)
}
