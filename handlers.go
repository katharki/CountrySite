package main

import (
	"fmt"
	"net/http"
)

func countryInfoHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("You have requested URL: %s\n", r.URL.Path)

	w.Header().Set("Content-Type", "application/json")

	//extract country code from the URL

	/*countryCode := r.URL.Path[len("/countryinfo/v1/info/"):]
	if countryCode == "" {
		http.Error(w, "Missing country code", http.StatusBadRequest)
		return
	}
	/*	*country, ok := countries[countryCode]

		if !ok {
			http.Error(w, "Coun	try not found", http.StatusNotFound)
	 */
			return
		}

	*/

	if r.URL.Path == "/countryinfo/v1/info/" {
		http.Error(w, "Missing country code", http.StatusBadRequest)
		return
	}

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

	w.Header().Set("Conent-Type", "application/json")
	json.NewEncoder(w).Encode(data)

	*/

}

func countryPopulationHandler(w http.ResponseWriter, r *http.Request) {

}
