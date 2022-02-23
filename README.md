# assignment_1
## Assignment task description
First assignment from the cloud technologi class(prog2005-2022).
This program is "develop using REST web application for retriving information about universities that may be candidates for application based on their name, alongside useful contextual information pertaining to the country it is situated in. For this purpose, you will interrogate existing web services and return the result in a given output format." Is the task of the assignment. 

There are two API that are going to be used in the application:

- http://universities.hipolabs.com/
  - Documentation/Source under: https://github.com/Hipo/university-domains-list/

- https://restcountries.com/
  - Documentation/Source under: https://gitlab.com/amatos/rest-countries

The final web service should be deployed on Heroku. The initial development should occur on your local machine. For the submission, you will need to provide both a URL to the deployed Heroku service as well as your code repository.

In the following, you will find the specification for the REST API exposed to the user for interrogation/testing.

## How to run the application
Fastes way to run the application is to run "go run app\cmd\main.go" in the terminal in the main directory.
To build an .exe file(executibale file) you can run the command "go build app\cmd\main.go" in the terminal.

## How to use the api
There is three different URLs avaieble to use from this api.

- /unisearcher/v1/uniinfo/
  This URL is used to search up different universeties from the URL. /unisearcher/v1/uniinfo/xxxx the xxxx is were you will use a keyword for the search. 

- /unisearcher/v1/diag/
  This URL will give out information about this api and if the other apis used for it is availeble.

- /unisearcher/v1/

### How the api used in the application works
- http://universities.hipolabs.com/
  The api works that you need a keyword to use if you want to search for a universety. It can be a completed name or partiale completed. http://universities.hipolabs.com/search?name=xxxx is how the URL is build up and xxxx is where the keyword for searching need to be writen. 

- https://restcountries.com/v3.1/
  The api works that you need a keyword to search for a country or you can use "ALL" to get all countries. https://restcountries.com/v3.1/xxxx where xxxx is were you need to insert the keyword to search for a country

## Credits
Made by Robin Ruud Kristensen

There is code directery taken from other poeple to make the application better or there have been code that have inspired how to write a spesific part of the code.
Refrence code links:
- https://go.dev/doc/articles/wiki/
- https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2022/-/tree/main/02-JSON-demo
- 

## License
                    GNU GENERAL PUBLIC LICENSE
                       Version 3, 29 June 2007

 Copyright (C) 2007 Free Software Foundation, Inc. <https://fsf.org/>
 Everyone is permitted to copy and distribute verbatim copies
 of this license document, but changing it is not allowed.
