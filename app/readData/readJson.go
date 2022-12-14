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
		countryInfo := ReadCountriesAPI(strings.ToLower(uniInfo[i].CountryName), constants.SEARCH_NAME)
		json.Unmarshal(responseData, &countryInfo)
		uniInfo[i].SetCountryInfo(countryInfo)
	}

	//return structs.RemoveDup(uniInfo)
	return uniInfo
}

/*
Function to read information from both of the apis.
It uses ReadUniInfo and ReadCountriesApi.
searchCountry string -> keyword for wich country to search after
searchUni string -> keyword for the name of the universiety filled or only partialy
return uniInfo -> returns a slice of data of the different universities*/
func ReadCountryUniInfo(searchCountry, searchUni string, remove int) []structs.UniInfo {
	countryInfo := ReadCountriesAPI(strings.ToLower(searchCountry), constants.SEARCH_NAME)
	uniInfo := ReadUniInfo(searchUni + "&country=" + countryInfo[0].Name.Common)

	for i := 0; i < len(countryInfo[0].Borders); i++ {
		borderCountry := ReadCountriesAPI(strings.ToLower(countryInfo[0].Borders[i]), constants.SEARCH_APLHA)
		uniInfo2 := ReadUniInfo(searchUni + "&country=" + borderCountry[0].Name.Common)

		for i := 0; i < len(uniInfo2); i++ {
			if strings.EqualFold(uniInfo2[i].CountryName, borderCountry[0].Name.Common) {
				uniInfo2[i].SetCountryInfo(borderCountry)
			}
		}

		for i := 0; i < len(uniInfo2)+1; i++ {
			if len(uniInfo2) >= remove && remove != 0 {
				uniInfo2 = structs.RemoveElement(uniInfo2, 1)
			}
		}
		//fmt.Print(uniInfo2[0].Languages)
		uniInfo = append(uniInfo, uniInfo2...)
	}

	return uniInfo
}

/*
Function to read information from the https://restcountries.com/v3.1/ api.
The const READ_ALL_COUNTRIES is the URL for the api.
var search string -> varible for what to be search for
return countryInfo -> returns a single country information in a slice. Returns in a slice because json saves information in slices*/
func ReadCountriesAPI(search, option string) []structs.Country {
	countryInfo := loadCountry()
	for i := 0; i < len(countryInfo); i++ {
		if strings.EqualFold(countryInfo[i].Name.Common, search) {
			saveCountry(countryInfo, search)
			return countryInfo
		}
	}

	countryInfo = nil

	switch option {
	case constants.SEARCH_NAME:
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
	case constants.SEARCH_APLHA:
		response, err := http.Get(constants.READ_ALL_COUNTRIES_APLHA + search)

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(responseData, &countryInfo)
	}

	saveCountry(countryInfo, search)

	return countryInfo
}
