package structs

/*
Constructor for a Uni*/
type Uni struct {
	Webpage        []string `json:"web_pages"`
	Name           string   `json:"name"`
	Aplha_two_code string   `json:"aplha_two_code"`
	Domains        []string `json:"domains"`
	Country        string   `json:"country"`
	State_province string   `json:"state-province,omitempty"`
}

/*
Constructor for a country */
type Country struct {
	Name           string `json:"name"`
	Aplha_two_code string `json:"aplha2Code"`
	Region         string `json:"region"`
	Capital        string `json:"captial"`
}

/*
Constructor for Diag information*/
type Diag struct {
	Universitiapi string
	Counteriesapi string
	Version       string
	Uptime        float64
}
