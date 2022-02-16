package structs

/*
Constructor for a Uni
Paremeters in the constructor are:
Webpage, Name, Aplha_two_code, Domains, Country, State_province*/
type Uni struct {
	Webpage        []string `json:"web_pages"`
	Name           string   `json:"name"`
	Aplha_two_code string   `json:"aplha_two_code"`
	Domains        []string `json:"domains"`
	Country        string   `json:"country"`
	State_province string   `json:"state-province,omitempty"`
}

/*
Constructor for a country
Parameters in the constructor are:
Name, Aplha_two_code, Region, Capital*/
type Country struct {
	Name struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	}
	//Name           string   `json:"name"`
	Aplha_two_code string   `json:"aplha2Code"`
	Region         string   `json:"region"`
	Capital        []string `json:"capital"`
}

/*
Constructor for Diag information
Parameters in the constructor are:
Universitiapi, Countriesapi, Version, Uptime*/
type Diag struct {
	Universitiapi int
	Counteriesapi int
	Version       string
	Uptime        float64
}
