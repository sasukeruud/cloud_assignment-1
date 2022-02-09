package main

import (
	"fmt"
	"os"
	//hipo "https://github.com/Hipo/university-domains-list.git"
)

/*
* Struct for a page, this is how it will be saved in memory
* Code from: https://go.dev/doc/articles/wiki/	*/
type Page struct {
	Title string
	Body  []byte
}

/*
* Function to save the body of the website in a .txt file named after the "Title"
* p --> Pointer to a object of a page
* return --> it will return an error (return value "0600" is for read write permissons for current user)
* Code from: https://go.dev/doc/articles/wiki/	*/
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

/*
* Function to get a spesific port to run the application on.
* It's needed for the application to run on heroku where to application is run on.*/
func getport() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

/*
* The function loadPage constructs the file name from the title parameter,
* reads the file's contents into a new variable body, and returns a pointer
* to a Page literal constructed with the proper title and body values.
* title --> Is the name of the spesific file that is going to be loaded in
*			the application.
* return --> If there is an error while the function is trying to load the
*			file it will return an error, else it will return the file with
*			the parameters for a webpage.
* Code from: https://go.dev/doc/articles/wiki/	*/
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	fmt.Println("hello world")
}
