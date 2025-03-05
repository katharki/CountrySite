package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// APIResponse - represents the response from the status endpoint API
// I am a bit unsure why i need alpha at the end
const (
	restCountriesAPI = "http://129.241.150.113:8080/v3.1/alpha/"
	//API for fetching current population of contires then
	countriesNowAPI = "http://129.241.150.113:3500/api/v0.1/countries/population"
)

// code string
func fetchCountryInfo(country string) (*CountryInfo, error) {
	url := fmt.Sprintf("%s%s", restCountriesAPI, country)
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
	// if err != nil {
	//	return nil, fmt.Errorf("Error. Could not read response body: %v", err) }

	//unmarshal the body into a CountryInfo struct
	//API returns an array of objects, so we need to unmarshal into a slice of CountryInfo
	var data []struct {
		Name       struct{ Common string } `json:"name"`
		Continents []string                `json:"Continent"`
		Population int                     `json:"population"`
		Languages  map[string]string       `json:"languages"`
		Borders    []string                `json:"borders"`
		Flag       struct{ PNG string }    `json:"flag"`
		Capital    []string                `json:"capital"`
		Cities     []string                `json:"cities"`
	}

	if err := json.Unmarshal(body, &data); err != nil || len(data) == 0 {
		return nil, fmt.Errorf("Error. Failed to unmarshal response body: %v", err)
	}

	//convert the data into a CountryInfo struct
	//we only need the first element of the array

	countryInfo := &CountryInfo{
		Name:       data[0].Name.Common,
		Continents: data[0].Continents[0],
		Population: data[0].Population,
		//jeg får også tips om bare "Languages: data[0].Languages"
		Languages: make([]string, 0, len(data[0].Languages)),
		Borders:   data[0].Borders,
		Flag:      data[0].Flag.PNG,
		Capital:   data[0].Capital[0],
	}

	for _, lang := range data[0].Languages {
		countryInfo.Languages = append(countryInfo.Languages, lang)
	}

	//handle case where capital is an empty string, unknown
	if len(data[0].Capital) > 0 {
		countryInfo.Capital = data[0].Capital[0]
	} else {
		countryInfo.Capital = "Unknown" //hvia vi ikkw vet hovedstaden
	}

	return countryInfo, nil

}

//ikke slett denne "}" over as. Ga meg sånn 15 problemer >_<

//HER SKAL VI FETCHE POPULATION FRA COUNTRIESNOWAPI
//DET ER HER ER DA EN EGEN DEL FRA RESTEN. slik at jeg vet og husker :p

func fetchPopulation(country string) (*PopulationData, error) {
	url := fmt.Sprintf("%s?country=%s", countriesNowAPI, country)
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("Error. Couldn't fetching population info: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error. Unexpected status code: %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("Error. Could not read response body: %v", err)
	}

	var result struct {
		Data struct {
			Country string `json:"country"`
			Yearly  []struct {
				Year  int `json:"year"`
				Value int `json:"value"`
			} `json:"Population"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("Error. Failed to unmarshal response body: %v", err)
	}

	//compute mean population
	total, count := 0, 0
	for _, p := range result.Data.Yearly {
		total += p.Value
		count++
	}
	mean := 0
	if count > 0 {
		mean = total / count
	}

	populationData := &PopulationData{
		Country: result.Data.Country,
		Mean:    mean,
		Yearly:  result.Data.Yearly,
	}

	return populationData, nil

}
