package main

import (
	"common/configuration"
	"fmt"
	"inputservice/service_pkg"
	"log4"
	"net/http"
	"utilityData"
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
		http.HandleFunc("/Handle", service_pkg.Handle)
		err = http.ListenAndServe(":8001", nil)

		if err != nil {
			utilityData.ErrorLog("main", "error", err.Error())
		}
	}
}
