package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

func countryInfoHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("You have requested URL: %s\n", r.URL.Path)

	//extract country name from the URL
	countryName := strings.TrimPrefix(r.URL.Path, "/countryinfo/v1/info/")
	countryName = strings.TrimSpace(countryName)
	if countryName == "" {
		http.Error(w, "Missing country code", http.StatusBadRequest)
		return
	}

	// fetch country info
	country, err := fetchCountryInfo(countryName)
	if err != nil {
		http.Error(w, "Country not found:(", http.StatusNotFound)
		return
	}

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

	countryName := strings.TrimPrefix(r.URL.Path, "/countryinfo/v1/population/")
	countryName = strings.TrimSpace(countryName)
	if countryName == "" {
		http.Error(w, "Missing country code", http.StatusBadRequest)
		return
	}

	//fetch population data
	population, err := fetchPopulation(countryName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
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

	response := map[string]string{
		"status":    "ok",
		"timestamp": time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
