package main

import (
	"fmt"
	"net/http"
)

// HomepageHandler - Displays basic API information
func homepageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, `Welcome to the Country Info API

Available Endpoints:
- /countryinfo/v1/info/{country_code}       - Get general info about a country
- /countryinfo/v1/population/{country_code} - Get historical population data
- /countryinfo/v1/status                    - Check API status

Example:
GET /countryinfo/v1/info/no     -> Returns information about Norway
GET /countryinfo/v1/population/no?limit=2010-2015 -> Returns population data for 2010-2015

Replace {country_code} with a valid two-letter ISO code (e.g., 'us' for USA, 'fr' for France).

`)
}
