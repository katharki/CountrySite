package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func countryInfoHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("You have requested URL: %s\n", r.URL.Path)

	//extract country name from the URL
	countryName := strings.TrimSpace(r.URL.Path[len("/countryinfo/v1/info/"):])
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
	w.Header().Set("Conent-Type", "application/json")
	json.NewEncoder(w).Encode(country)

}

func countryPopulationHandler(w http.ResponseWriter, r *http.Request) {

}
