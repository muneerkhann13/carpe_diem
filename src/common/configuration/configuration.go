package configuration

import (
	"encoding/json"
	"io/ioutil"
	"model"

	"log"
)

//Declare global variable use to set global configuration once at time of project startup
var Config *model.Config

/*
   Author      : Mukesh Dutt
   CreatedDate : 23-September-2016
   Purpose     : Read json file and feed to structure
   Parameters  : {"input":"path", "output":""}
   Module      : Common
   LastUpdate  : { "name": "mukesh dutt", "date":"23rd november 2016" }
*/
func LoadConfig(path string) {

	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		//_, err = travelData.InsertErrorLogs(100, fmt.Sprint(errors.Errorf(err.Error()).ErrorStack()))
	}

	err = json.Unmarshal(file, &Config)
	if err != nil {
		log.Fatal(err)
		//_, err = travelData.InsertErrorLogs(100, fmt.Sprint(errors.Errorf(err.Error()).ErrorStack()))
	}
}
