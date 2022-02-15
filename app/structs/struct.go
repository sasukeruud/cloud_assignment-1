package structs

import "time"

//"json"

/*
* Struct for a page, this is how it will be saved in memory
* Code from: https://go.dev/doc/articles/wiki/	*/
type Page []struct {
	Title string
	Body  []byte
}

/*
* Constructor for a Uni*/
type Uni struct {
	Webpage        string `json:"web_pages"`
	Name           string `json:"name"`
	Aplha_two_code string `json:"aplha_two_code`
	State_province string `json:"state-province"`
	Domains        string `json:"domains"`
	Country        string `json:"country"`
}

/*
* Constructor for a country */
type Country struct {
	Name           string `json:"name"`
	Aplha_two_code string `json:"aplha2Code"`
	Region         string `json:"region"`
	Capital        string `json:"captial"`
}

/*
* Constructor for Diag information*/
type Diag struct {
	Universitiapi string
	Counteriesapi string
	Version       string
	Uptime        time.Timer
}
