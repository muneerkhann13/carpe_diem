package service_pkg


import (
	"encoding/json"
	"model"
	"utilityData"
	//"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	defer utility.CatchPanic(w, "handler")

	// var file *os.File
	var request model.Request
	var err error
	var response model.Response
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utilityData.ErrorLog("handler", "error", err.Error())
	}

	// raw request log in file
	utilityData.RawRequestLog("", string(body), "")

	//Insert raw request for trace every request
	err = utilityData.InsertRawRequest("", string(body), "")

	if err != nil {

		utilityData.ErrorInFile(err)
		response = utility.ErrorHandling(err, 101)
	} else {
		err = json.Unmarshal(body, &request)
		if err != nil {

			utilityData.ErrorLog("handler", "error", err.Error())

			response = utility.ErrorHandling(err, 101)
		} else {
			response = GetResponse(request, body)
		}
	}

	plainResponse, err := json.Marshal(response)
	if err != nil {
		utilityData.ErrorLog("handle", "error", err.Error())
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	io.WriteString(w, string(plainResponse))
}
