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

		fmt.Println("welcome to travel aggregator")

		// Oxiface Exteneded Handler
		http.HandleFunc("/OxifaceExtended-Plain", apiHandler.OxifaceExtendedPlain)

		// Oxiface Extended Handler with encryption
		http.HandleFunc("/OxifaceExtended", apiHandler.OxifaceExtended)
		err = http.ListenAndServe(":8080", nil)

		if err != nil {

			utilityData.ErrorLog("main", "error", err.Error())

		}
	}
}
