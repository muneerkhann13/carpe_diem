package configuration

import (
	"encoding/json"
	"io/ioutil"
	"model"

	"log"
)

//Declare global variable use to set global configuration once at time of project startup
var Config *model.Config

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
