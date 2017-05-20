package utility

import (
	"fmt"
	"model"
	"net/http"
	"runtime"
	"strings"
	"utilityData"

	"io"

	"github.com/go-errors/errors"
)

//Function will be called when there is error to log error and to send travelModel.Response
func ErrorHandling(err error, requestId int64) model.Response {

	//Variable declaration
	var Response model.Response

	if err != nil {
		Response.Code = SomeThingWentWrong
		Response.Description = StatusCodes(SomeThingWentWrong)
	}

	//Inserting error in DB
	utilityData.InsertErrorLogs(requestId, fmt.Sprint(errors.Errorf(err.Error()).ErrorStack()))
	if err != nil {
		Response.Code = SomeThingWentWrong
		Response.Description = StatusCodes(SomeThingWentWrong)

	}
	return Response
}

//Function will be called when there is success to send response

func SuccessResult(Response travelModel.Response) travelModel.Response {
	Response.Code = Success
	Response.Description = StatusCodes(Success)

	return Response
}

//To check the string is available in list or not

func StringInSlice(requestID, str string, list []string) bool {

	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func CatchPanic(w http.ResponseWriter, requestID string) {

	if x := recover(); x != nil {
		err, ok := x.(error)
		if ok == true {
			src := ""
			for i := 1; ; i++ {
				pc, _, lineno, _ := runtime.Caller(i)
				if strings.Contains(runtime.FuncForPC(pc).Name(), "Carpediem") {
					src = fmt.Sprint(runtime.FuncForPC(pc).Name(), ":", lineno)
					break
				}
			}
			if src != "" {
				utilityData.ErrorLog(requestID, "error", err.Error()+"["+src+"]")
			} else {
				utilityData.ErrorLog(requestID, "error", "Some panic has occured.")
			}
			io.WriteString(w, "Some thing went wrong. Please try again")
		}
	}
}
