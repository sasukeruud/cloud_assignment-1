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
		if country[0].Name.Common != loadCountry[0].Name.Common {
			for i := 0; i < len(loadCountry); i++ {
				if strings.EqualFold(loadCountry[i].Name.Common, name) {
					temp := loadCountry[i]
					for i := len(loadCountry) - 1; i > 0; i-- {
						country[i] = country[i-1]
						country[i-1] = temp
					}
				}
			}
		} else {
			for i := 0; i < len(loadCountry); i++ {
				if strings.EqualFold(loadCountry[i].Name.Common, name) {
					temp := loadCountry[i]
					for i := len(loadCountry) - 1; i > 0; i-- {
						country[i] = country[i-1]
						country[i-1] = temp
					}
				}
			}
		}
		loadCountry = structs.RemoveDup(loadCountry)
		country = append(country, loadCountry...)
		country = structs.RemoveDup(country)
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
