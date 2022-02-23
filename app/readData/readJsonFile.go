package readApi

import (
	"assignment_1/app/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func loadCountry() []structs.Country {
	file, err := os.Open("res/countries.json")
	var countries []structs.Country

	if err != nil {
		fmt.Println(err.Error())
		return countries
	}

	byteValue, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(byteValue, &countries)

	defer file.Close()

	return countries
}

func saveCountry(country []structs.Country, name string) {
	file, err := json.MarshalIndent(country, "", " ")

	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile("res/countries.json", file, 0644)

	f, err := os.OpenFile("res/countries.json", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	f.Write(file)
	defer f.Close()
}
