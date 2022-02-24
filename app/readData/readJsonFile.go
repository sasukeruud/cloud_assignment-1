package readApi

import (
	constants "assignment_1/app"
	"assignment_1/app/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

/*
Function to load in the json file.
returns countries -> a slice of []structs.Country made from json file*/
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

	defer file.Close() //makes sure that the file is closed

	return countries
}

/*
Function to save a country slice to a json file. Also runs a loop to change the
order of how countries are saved in the json file.
countries -> a slice of a []structs.country that will be saved.
name string -> variable of the country name to search for a spesific country.*/
func saveCountry(countries []structs.Country, name string) {
	loadCountry := loadCountry()

	if loadCountry != nil {
		if strings.EqualFold(loadCountry[0].Name.Common, name) {
			return //Returns out of function if the first element is a hit, not continuing the rest of the function
		}
		for i := 0; i < len(loadCountry); i++ {
			if strings.EqualFold(loadCountry[i].Name.Common, name) || strings.EqualFold(loadCountry[i].Cca2, name) {
				temp := loadCountry[i]
				for j := i - 1; j > 0; j-- { //Starts from where to country was located, and moves backwords
					countries[j] = countries[j-1] //moves the element to the right
					countries[j-1] = temp         //moves temp element to the left
				}
				break //Goes out of loop when it's finished
			}
		}
	}

	countries = structs.RemoveDup(append(countries, loadCountry...))

	if len(countries) > constants.JSON_LIMIT {
		for i := len(countries) - 1; i > constants.JSON_LIMIT; i-- { //Goes from last element of slice and removes an object
			structs.RemoveElementCountry(countries, i)
		}
	}

	file, err := json.MarshalIndent(countries, "", " ")

	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile("res/countries.json", file, 0644)

	f, err := os.OpenFile("res/countries.json", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	f.Write(file)
	defer f.Close() //makes sure that the file is closed
}
