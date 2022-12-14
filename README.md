# assignment_1
Heroku instance is not running at the movement
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
Or use the URL: https://robinruassignment1.herokuapp.com/ where it is deployed. 

### How to use the api
There is three different URLs avaieble to use from this api.

- /unisearcher/v1/uniinfo/  
  This URL is used to search up different universeties from the URL. /unisearcher/v1/uniinfo/xxxx the xxxx is were you will use a keyword for the search. There is no limit for what you can use as a keyword. It`s not key sensetiv when it comes to the keyword writen.
```
                  GET request:
                    /unisearcher/v1/uniinfo/norway
                    Will search for all universities that have norway in it`s name and connect all the different countries 
                    information to it too
```
Body example:

```
[
  {
  "name": "Norway's Information Technology University College",
  "web_pages": [
  "http://www.nith.no/"
  ],
  "domains": [
    "nith.no"
  ],
  "country": "Norway",
  "cca2": "NO",
  "languages": {
    "nno": "Norwegian Nynorsk",
    "nob": "Norwegian Bokm??l",
    "smi": "Sami"
  },
  "maps": {
    "googleMaps": "https://goo.gl/maps/htWRrphA7vNgQNdSA",
    "openStreetMaps": "https://www.openstreetmap.org/relation/2978650"
  }
  },
  ...
]            
```                    
                      
- /unisearcher/v1/neighbourunis/  
  This URL is used to search for different universities in bordering countries. /unisearcher/v1/neighbourunis/xxxx/yyyy?limit=i xxxx is the main country that you will check the bordering countries for, and yyyy is for the keyword in the univerities name. You can also set a limit of how many universities will be listed.
  ```
                  GET request:
                  /unisearcher/v1/neighbourunis/norway/science
                  Will search for all universities that have science in there name in norway and in bordering countries
  ```

  ```
                  GET request:
                  /unisearcher/v1/neighbourunis/norway/science?limit=5
                  Will search for all universities that have science in there name in norway and in bordering countries. 
                  With using limit there will be a limit on how many universities from each bordring country will be 
                  listed, all univerities from the main country will be listed. 
  ```

Body example:
```
[
  {
      "name": "Norwegian University of Science and Technology",
      "web_pages": [
          "http://www.ntnu.no/"
      ],
      "domains": [
          "ntnu.no"
      ],
      "country": "Norway",
      "cca2": "NO",
      "languages": {
          "nno": "Norwegian Nynorsk",
          "nob": "Norwegian Bokm??l",
          "smi": "Sami"
      },
      "maps": {
          "googleMaps": "https://goo.gl/maps/htWRrphA7vNgQNdSA",
          "openStreetMaps": "https://www.openstreetmap.org/relation/2978650"
      }
  },
  ...
  {
      "name": "Central Ostrobothnia University of Applied Sciences",
      "web_pages": [
          "http://www.cou.fi/"
      ],
      "domains": [
          "cou.fi"
      ],
      "country": "Finland",
      "cca2": "FI",
      "languages": {
          "swe": "Swedish",
          "fin": "Finnish"
      },
      "maps": {
          "googleMaps": "https://goo.gl/maps/HjgWDCNKRAYHrkMn8",
          "openStreetMaps": "openstreetmap.org/relation/54224"
      }
  },
  ...
]  
```

- /unisearcher/v1/diag/  
  This URL will give out information about this api and if there is connection with the other apis

## Credits
Made by Robin Ruud Kristensen

There is code directery taken from other poeple to make the application better or there have been code that have inspired how to write a spesific part of the code.
Refrence code links:
- https://go.dev/doc/articles/wiki/
- https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2022/-/tree/main/02-JSON-demo

## License
                    GNU GENERAL PUBLIC LICENSE
                       Version 3, 29 June 2007

 Copyright (C) 2007 Free Software Foundation, Inc. <https://fsf.org/>
 Everyone is permitted to copy and distribute verbatim copies
 of this license document, but changing it is not allowed.
