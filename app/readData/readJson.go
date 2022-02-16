package readApi

import (
	constants "assignment_1/app"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

/*
Function to read information from the http://universities.hipolabs.com/ api.
the const READ_ALL_UNI is the URL for the api.
var search string -> varible for what to be search for.
return responseData -> returns a byte slice of information */
func ReadUniAPI(search string) []byte {
	response, err := http.Get(constants.READ_ALL_UNI + search)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}

/*
Function to read information from the https://restcountries.com/v3.1/ api.
The const READ_ALL_COUNTRIES is the URL for the api.
var search string -> varible for what to be search for
return responseData -> returns a byte slice of information */
func readCountriesAPI(search string) []byte {
	response, err := http.Get(constants.READ_ALL_COUNTRIES + search)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}
