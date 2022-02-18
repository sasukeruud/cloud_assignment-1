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
		var countryInfo []structs.Country
		response, err := http.Get(constants.READ_ALL_COUNTRIES + strings.ToLower(uniInfo[i].CountryName))

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(responseData, &countryInfo)
		uniInfo[i].SetCountryInfo(countryInfo)
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
	response, err := http.Get(constants.READ_ALL_COUNTRIES + search)

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
