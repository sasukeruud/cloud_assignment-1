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
