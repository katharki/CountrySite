package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

// APIResponse - represents the response from the status endpoint API
// I am a bit unsure why i need alpha at the end
const (
	restCountriesAPI = "http://129.241.150.113:8080/v3.1/alpha/"
	//API for fetching current population of contires then
	countriesNowAPI = "http://129.241.150.113:3500/api/v0.1/countries/population"
)

//coyntryInfo struct

// code string
func fetchCountryInfo(countryCode string) (*CountryInfo, error) {

	countryCode = strings.ToUpper(countryCode)

	url := fmt.Sprintf("%s%s", restCountriesAPI, countryCode)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error. Could not fetch country info: %v", err)
	}
	defer resp.Body.Close()

	//check if status code is 200
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error. Unexected sstatus code: %v", resp.StatusCode)
	}

	//read the body of the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error. Could not read response body: %v", err)
	}

	fmt.Println("DEBUG: Raw API response: ", string(body))

	//unmarshal the body into a CountryInfo struct
	//API returns an array of objects, so we need to unmarshal into a slice of CountryInfo
	var data []struct {
		Name       struct{ Common string } `json:"name"`
		Continents []string                `json:"continents"`
		Population int                     `json:"population"`
		Languages  map[string]string       `json:"languages"`
		Borders    []string                `json:"borders"`
		Flag       struct{ PNG string }    `json:"flag"`
		Capital    []string                `json:"capital"`
		Cities     []string                `json:"cities"`
	}

	//(should be) safe unmarshal
	if err := json.Unmarshal(body, &data); err != nil || len(data) == 0 {
		fmt.Println("DEBUG: Raw API resposme before parsing error:", string(body))
		fmt.Println("DEBUG: Failed to parse API response for:", countryCode)
		return nil, fmt.Errorf("Error. Failed to unmarshal response body: %v", err)
	}

	//menes at det skal skje en "panic" hvis jeg prøver å accesse koden med en data[0]
	//dermed skal jeg bruke denne if'en for å unngå det
	if len(data) == 0 {
		fmt.Println("DEBUG: Empty response from API for country: ", countryCode)
		return nil, fmt.Errorf("Error. No country data found for : %s", countryCode)

	}

	//handling case where continent is an empty string, unknown, safe
	var continent string
	if len(data[0].Continents) > 0 {
		continent = data[0].Continents[0]
	} else {
		continent = "Unknown"
	}

	//handling case where capital is an empty string, unknown, safe
	var capital string
	if len(data[0].Capital) > 0 {
		capital = data[0].Capital[0]
	} else {
		capital = "Unknown"
	}

	//handling case where flag is an empty string, unknown, safe
	var flag string
	if data[0].Flag.PNG != "" {
		flag = data[0].Flag.PNG
	} else {
		flag = "Unknown"
	}

	//sort cities, aphabetically
	sort.Strings(data[0].Cities)

	//limit on cities, doest work, dont need
	/*
		if limit > 0 && limit < len(data[0].Cities) {
			data[0].Cities = data[0].Cities[:limit]
		}

	*/

	//convert the data into a CountryInfo struct
	//we only need the first element of the array

	countryInfo := &CountryInfo{
		Name:       data[0].Name.Common,
		Continents: continent,
		Population: data[0].Population,
		//jeg får også tips om bare "Languages: data[0].Languages"
		Languages: make([]string, 0, len(data[0].Languages)),
		Borders:   data[0].Borders,
		Flag:      flag,
		Capital:   capital,
		Cities:    data[0].Cities,
	}

	//copy languages into the countryInfo struct
	for _, lang := range data[0].Languages {
		countryInfo.Languages = append(countryInfo.Languages, lang)
	}

	/*
		//handle case where capital is an empty string, unknown
		if len(data[0].Capital) > 0 {
			countryInfo.Capital = data[0].Capital[0]
		} else {
			countryInfo.Capital = "Unknown" //hvis vi ikke vet hovedstaden
		}
	*/
	return countryInfo, nil

}

//ikke slett denne "}" over as. Ga meg sånn 15 problemer >_<

//HER SKAL VI FETCHE POPULATION FRA COUNTRIESNOWAPI
//DET ER HER ER DA EN EGEN DEL FRA RESTEN. slik at jeg vet og husker :p

func fetchPopulation(countryCode string) (*PopulationData, error) {

	countryCode = strings.ToUpper(countryCode)

	url := fmt.Sprintf("%s?country=%s", countriesNowAPI, countryCode)
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("Error. Couldn't fetching population info for %s: %v", countryCode, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error. Unexpected status code %d for country %s: %v", resp.StatusCode, countryCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("Error. Could not read response body for %s: %v", countryCode, err)
	}

	var result struct {
		Data struct {
			Country string `json:"country"`
			Yearly  []struct {
				Year  int `json:"year"`
				Value int `json:"value"`
			} `json:"yearly"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("DEBUG: Rawy API population response before parsing error:", string(body))
		fmt.Println("DEBUG: Failed to parse API response for:", countryCode)
		return nil, fmt.Errorf("Error. Failed to unmarshal response body %s: %v", countryCode, err)
	}

	//compute mean population
	total, count := 0, 0
	for _, p := range result.Data.Yearly {
		total += p.Value
		count++
	}
	mean := 0
	if count > 0 {
		mean = int(float64(total) / float64(count))
	}

	populationData := &PopulationData{
		Country: result.Data.Country,
		Mean:    mean,
		Yearly:  result.Data.Yearly,
	}

	return populationData, nil

}
