package main

import (
	"fmt"
	"log"
	"net/http"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	//	/countryinfo/v1/info/
	//	/countryinfo/v1/population/
	//	/countryinfo/v1/status/

	router := http.NewServeMux()

	//homepage
	router.HandleFunc("/", homepage)

	router.HandleFunc("/countryinfo/v1/info/", countryInfoHandler)
	router.HandleFunc("/countryinfo/v1/population/", populationHandler)
	router.HandleFunc("/countryinfo/v1/status/", statusHandler)

	//start server on port 8080
	port := "8080"
	fmt.Printf("Server running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
