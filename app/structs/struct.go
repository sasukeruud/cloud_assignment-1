package structs

/*
Constructor for a country
Parameters in the constructor are:
Name, Cca2, Region, Capital, Borders, Language, Maps*/
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
Constructor for university information.
Parameter in the constructor is:
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
country []Country -> country slice needed for the function to do it's functions
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

/*
Function that will run a algorithm to remove duplicated universeties added to the slice.
u []UniInfo -> slice that will be check for duplications.
return u -> return the slice after it have been check and edited*/
func RemoveDup(u []UniInfo) []UniInfo {
	for i := 0; i < len(u); i++ {
		for j := i + 1; j < len(u); j++ {
			if u[i].NameUni == u[j].NameUni {
				u = append(u[:j], u[j+1:]...)
			}
		}
	}

	return u
}

/*
Function to remove an element from a slice
u []UniInfo -> slice that will be work on
i int -> number how wich element in the slice that will be removed
return []UniInfo -> returns a new slice*/
func RemoveElement(u []UniInfo, i int) []UniInfo {
	u[i] = u[len(u)-1]
	return u[:len(u)-1]
}
