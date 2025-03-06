package main

import (
	"fmt"
	"net/http"
)

//homepage server

func homepage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `
		<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Country Info API</title>
	<style>
		body { font-family: Arial, sans-serif; text-align: center; padding: 50px; }
	h1 { color: #333; }
	p { font-size: 18px; }
	a { color: blue; text-decoration: none; font-weight: bold; }
	</style>
	</head>
	<body>
	<h1>🌍 Country Info API</h1>
	<p>Welcome to the Country Info API! You can retrieve country information and population data using the following endpoints:</p>
	<ul>
	<li><a href="/countryinfo/v1/info/no">/countryinfo/v1/info/{country_code}</a> - Get country details</li>
	<li><a href="/countryinfo/v1/population/no">/countryinfo/v1/population/{country_code}</a> - Get population data</li>
	<li><a href="/countryinfo/v1/status">/countryinfo/v1/status</a> - Check API status</li>
	</ul>
	<p>Example: <a href="/countryinfo/v1/info/no">Get Norway's Info</a></p>
	<p>Try adding a query filter for population data: <strong>?limit=2010-2015</strong></p>
	</body>
	</html>
	`)

	//random linje
}
