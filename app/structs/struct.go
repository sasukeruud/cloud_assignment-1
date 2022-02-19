package structs

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
	Borders  []string `json:"borders"`
	Language struct {
		Nno string `json:"nno"`
		Nob string `json:"nob"`
		Smi string `json:"smi"`
	} `json:"languages,omitempty"`
	Maps struct {
		OpenStreetMaps string `json:"openStreetMaps"`
	} `json:"maps"`
}

/*
Constructor for universety information.
Parameter in the constructor are:
Webpage, NameUni, Domains, CountryName, Cca2, Capital, Maps
Parameter: Language will be changed to be more fitting at a later date
*/
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

/*
Setter function.
The fields it will edit is:
Cca2, Capital, Language, Maps
o *UniInfo -> says that it is a UniInfo object that the function will work on
country []Country -> country slice needed for the function to do it\s funtions
*/
func (o *UniInfo) SetCountryInfo(country []Country) {
	o.Cca2 = country[0].Cca2
	o.Capital = country[0].Capital
	o.Language = country[0].Language
	o.Maps = country[0].Maps
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
