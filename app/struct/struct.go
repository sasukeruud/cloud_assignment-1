package structs

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
	webpage        string `json:"web_pages"`
	name           string `json:"name"`
	aplha_two_code string `json:"aplha_two_code`
	state_province string `json:"state-province"`
	domains        string `json:"domains"`
	country        string `json:"country"`
}

/*
* Constructor for a country */
type Country struct {
	name           string `json:"name"`
	aplha_two_code string `json:"aplha2Code"`
	region         string `json:"region"`
	capital        string `json:"captial"`
}
