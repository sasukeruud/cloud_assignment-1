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
)

func ReadUniAPI() {
	//http://universities.hipolabs.com/search?name=
	response, err := http.Get(constants.READ_ALL_UNI + "Marywood")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

	uni := structs.Uni{}
	json.Unmarshal(responseData, &uni)
	fmt.Println(uni.Name)
}

func readCountriesAPI() {
	//https://restcountries.com/v3.1/
	response, err := http.Get(constants.READ_ALL_COUNTRIES + "")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

	country := structs.Country{}
	json.Unmarshal(responseData, &country)
	fmt.Println(country.Name)
}
