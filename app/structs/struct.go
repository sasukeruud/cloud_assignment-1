package structs

/*
Constructor for a Uni
Paremeters in the constructor are:
Webpage, Name, Aplha_two_code, Domains, Country, State_province*/
type Uni struct {
	Webpage        []string `json:"web_pages"`
	Name           string   `json:"name"`
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
	} `json:"name"`
	Cca2     string   `json:"cca2"`
	Region   string   `json:"region"`
	Capital  []string `json:"capital"`
	Language struct {
		Nno string `json:"nno"`
		Nob string `json:"nob"`
		Smi string `json:"smi"`
	} `json:"languages,omitempty"`
	Maps struct {
		OpenStreetMaps string `json:"openStreetMaps"`
	} `json:"maps"`
}

type UniInfo struct {
	Webpage     []string `json:"web_pages"`
	NameUni     string   `json:"name"`
	Domains     []string `json:"domains"`
	CountryName string   `json:"country"`
	Cca2        string   `json:"cca2"`
	Capital     []string `json:"capital"`
	Language    struct {
		Nno string `json:"nno"`
		Nob string `json:"nob"`
		Smi string `json:"smi"`
	} `json:"languages,omitempty"`
	Maps struct {
		OpenStreetMaps string `json:"openStreetMaps"`
	} `json:"maps"`
}

func (o *UniInfo) SetCca2(country []Country) {
	o.Cca2 = country[0].Cca2
	o.Capital = country[0].Capital
	o.Language = country[0].Language
	o.Maps = country[0].Maps
}

type UniInfo2 struct {
	UniDetails     []Uni
	CountryDetails []Country
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
