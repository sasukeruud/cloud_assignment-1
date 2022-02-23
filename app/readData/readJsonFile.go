package readApi

import (
	"assignment_1/app/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
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
	loadCountry := loadCountry()

	if loadCountry != nil {
		for i := 0; i < len(loadCountry); i++ {
			if strings.EqualFold(loadCountry[i].Name.Common, name) || strings.EqualFold(loadCountry[i].Cca2, name) {
				temp := loadCountry[i]
				for j := i - 1; j > 0; j-- {
					country[j] = country[j-1]
					country[j-1] = temp
				}
				break
			}
		}
		loadCountry = structs.RemoveDup(loadCountry)
		country = structs.RemoveDup(append(country, loadCountry...))
	}

	if len(country) > 50 {
		for i := len(country) - 1; i > 30; i-- {
			structs.RemoveElementCountry(country, i)
		}
		fmt.Println("1")
	}

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
