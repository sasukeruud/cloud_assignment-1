package readApi

import (
	constants "assignment_1/app"
	"assignment_1/app/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
Function to read information from the http://universities.hipolabs.com/ api.
the const READ_ALL_UNI is the URL for the api.
var search string -> varible for what to be search for.
return uniInfo -> returns slice information about the universety */
func ReadUniInfo(search string) []structs.UniInfo {
	var uniInfo []structs.UniInfo

	response, err := http.Get(constants.READ_ALL_UNI + search)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &uniInfo)
	for i := 0; i < len(uniInfo); i++ {
		countryInfo := ReadCountriesAPI(strings.ToLower(uniInfo[i].CountryName))

		json.Unmarshal(responseData, &countryInfo)
		uniInfo[i].SetCountryInfo(countryInfo)
	}

	return uniInfo
}

func ReadCountryUniInfo(searchCountry, searchUni string) []structs.UniInfo {
	countryInfo := ReadCountriesAPI(strings.ToLower(searchCountry))
	uniInfo := ReadUniInfo(searchUni + "&country=" + countryInfo[0].Name.Common)

	for i := 0; i < len(countryInfo[0].Borders); i++ {
		var borderCountry []structs.Country

		responseCountry, err := http.Get(constants.READ_ALL_COUNTRIES_APLHA + strings.ToLower(countryInfo[0].Borders[i]))

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseCountryData, err := ioutil.ReadAll(responseCountry.Body)
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(responseCountryData, &borderCountry)

		uniInfo2 := ReadUniInfo(searchUni + "&country=" + borderCountry[0].Name.Common)

		for i := 0; i < len(uniInfo2); i++ {
			if strings.EqualFold(uniInfo2[i].CountryName, borderCountry[0].Name.Common) {
				uniInfo2[i].SetCountryInfo(borderCountry)
			}
		}

		uniInfo = append(uniInfo, uniInfo2...)
	}

	return uniInfo
}

/*
Function to read information from the https://restcountries.com/v3.1/ api.
The const READ_ALL_COUNTRIES is the URL for the api.
var search string -> varible for what to be search for
return countryInfo -> returns a single country information in a slice. Returns in a slice because json saves information in slices*/
func ReadCountriesAPI(search string) []structs.Country {
	var countryInfo []structs.Country

	response, err := http.Get(constants.READ_ALL_COUNTRIES_NAME + search)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &countryInfo)

	return countryInfo
}
