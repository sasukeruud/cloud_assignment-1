package structs

//"json"

type Page struct {
}

type Uni struct {
	webpage        string `json:"web_pages"`
	name           string `json:"name"`
	aplha_two_code string `json:"aplha_two_code`
	state_province string `json:"state-province"`
	domains        string `json:"domains"`
	country        string `json:"country"`
}

type Country struct {
	name    string `json:"name"`
	region  string `json:"region"`
	capital string `json:"captial"`
}
