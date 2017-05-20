package main

import (
	"common/configuration"
	"fmt"
	"utilityData"

	//"TravelAggregator/bus/data"

	"log4"
	"net/http"
)

func init() {

	// Configuration Setting Globally
	configuration.LoadConfig("./config.json")
}

func main() {

	// var file *os.File
	var err error

	utilityData.Logg, err = log4.LoadConfiguration("logConfig.json")
	if err != nil {
		fmt.Println("err")
	}

	err = utilityData.StartDB()

	if err != nil {
		utilityData.ErrorLog("Main", "error", err.Error())
	} else {

		fmt.Println("welcome to carpediem")

		// Oxiface Exteneded Handler
		http.HandleFunc("/test", apiHandler.Handle)
		err = http.ListenAndServe(":7000", nil)

		if err != nil {
			utilityData.ErrorLog("main", "error", err.Error())
		}
	}
}
