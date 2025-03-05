package main

//Countryinfo - represents data from teh REST API

type CountryInfo struct {
	Name       string   `json:"name"`
	Continents string   `json:"continent"`
	Population int      `json:"population"`
	Languages  []string `json:"languages"`
	Borders    []string `json:"borders"`
	Flag       string   `json:"flag"`
	Capital    string   `json:"capital"`
	Cities     []string `json:"cities"`
}

// PopulationData - represents data for population history data
type PopulationData struct {
	Country string `json:"country"`
	Mean    int    `json:"mean"`
	Yearly  []struct {
		Year  int `json:"year"`
		Value int `json:"value"`
	} `json:"yearly"`
}

// APIResponse - represents the response from the status endpoint API
type APIResponse struct {
	CountriesNowAPI  string `json:"countriesNowAPI"`
	RestCountriesAPI string `json:"restCountriesAPI"`
	Version          string `json:"version"`
	Uptime           string `json:"uptime"

`
}

var countries = map[string]*CountryInfo{
	"NO": &CountryInfo{
		Name:       "Norway",
		Continents: "Europe",
		Population: 5367580,
		Languages:  []string{"Norwegian", "Sami"},
		Borders:    []string{"Sweden", "Finland", "Russia"},
		Flag:       "https://restcountries.com/data/nor.svg",
		Capital:    "Oslo",
		Cities:     []string{"Bergen", "Stavanger", "Trondheim"},
	},
	"SE": &CountryInfo{
		Name:       "Sweden",
		Continents: "Europe",
		Population: 10367232,
		Languages:  []string{"Swedish"},
		Borders:    []string{"Norway", "Finland"},
		Flag:       "https://restcountries.com/data/swe.svg",
		Capital:    "Stockholm",
		Cities:     []string{"Gothenburg", "Malmö", "Uppsala"},
	},
}
