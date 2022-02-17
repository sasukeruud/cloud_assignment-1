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

var uniInfo []structs.UniInfo

//var CountryInfo structs.Country

/*
Function to read information from the http://universities.hipolabs.com/ api.
the const READ_ALL_UNI is the URL for the api.
var search string -> varible for what to be search for.
return responseData -> returns a byte slice of information */
func ReadUniAPI(search string) []structs.UniInfo {
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
		//fmt.Println(uniInfo[i].CountryName)
		ReadCountriesAPI(strings.ToLower(uniInfo[i].CountryName))
		uniInfo[i].SetCca2(ReadCountriesAPI(strings.ToLower(uniInfo[i].CountryName)))
		//&uniInfo[i].Cca2 = &CountryInfo.Cca2
		//CountryInfo.Cca2 = "NO"
		//fmt.Print(CountryInfo.Cca2)
		//uniInfo[i].SetCca2(CountryInfo.Cca2)
		//&uniInfo[i].Cca2 = CountryInfo.Cca2
		//CountryInfo.Cca2 = &uniInfo[i].Cca2
		//fmt.Println((uniInfo[i].Cca2))
		//uniInfo[i].Country = &CountryInfo
	}
	//json.Unmarshal(ReadCountriesAPI(uniInfo[0].Country), &uniInfo[0])

	return uniInfo
}

/*
Function to read information from the https://restcountries.com/v3.1/ api.
The const READ_ALL_COUNTRIES is the URL for the api.
var search string -> varible for what to be search for
return responseData -> returns a byte slice of information */
func ReadCountriesAPI(search string) []structs.Country {

	var CountryInfo []structs.Country
	response, err := http.Get(constants.READ_ALL_COUNTRIES + search)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &CountryInfo)

	fmt.Println(CountryInfo[0].Name.Common)

	return CountryInfo
}
