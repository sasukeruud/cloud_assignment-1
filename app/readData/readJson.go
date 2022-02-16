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
* Function to read information from the http://universities.hipolabs.com/ api.
* 	The api works that you need a keyword to search for and READ_ALL_UNI const
*	have the URL without the keyword used for searching, there is were the search
*	varible is used
* var search string -> varible for what to be search for
* return responseData -> returns a byte slice of information */
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
* */
func readCountriesAPI(search string) []byte {
	//https://restcountries.com/v3.1/
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
