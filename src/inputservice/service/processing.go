package apiHandler

import (
	"encoding/json"
	"fmt"
	"model"
	"strings"
	"utilityData"

	"common/configuration"

	"github.com/go-errors/errors"
)

func GetResponse(request model.Request, jsonbody []uint8) model.Response {

	response := travelModel.Response{}

	//Validating request
	hostCode, hostDescription, ok := validations(request)

	// insert request in file log

	if err != nil {
		utilityData.InsertErrorLogs(requestID, fmt.Sprint(errors.Errorf(err.Error()).ErrorStack()))
		hostCode = configuration.SomeThingWentWrong
		hostDescription = configuration.StatusCodes(configuration.SomeThingWentWrong)
	} else {
		if ok {

			switch strings.ToUpper(request.ServiceRequest.Serviceinfo.Category) {

			case "userservice":
				response = travelHotel.GetHotelResponse(requestID, request.Service, jsonbody)

			case "nearbyservice":
				//response = travelFlight.GetFlightResponse(request.ServiceRequest.Transactioninfo.RequestID, requestID, request.Service, jsonbody)

			case "offerservice":
				//response = travelBus.GetBusResponse(request.Service, jsonbody)

			default:
				response.ServiceResponse.ResponseInfo.HostCode = configuration.ServiceCategoryDoesntExists
				response.ServiceResponse.ResponseInfo.HostDescription = configuration.StatusCodes(configuration.ServiceCategoryDoesntExists)
			}

			// When validation failed
		} else {
			response.ServiceResponse.ResponseInfo.HostCode = hostCode
			response.ServiceResponse.ResponseInfo.HostDescription = hostDescription
		}
	}

	//Set global parameters values
	response.Service = request.Service
	response.Version = request.Version
	response.ServiceResponse.User.Mdn = request.ServiceRequest.User.Mdn

	//Converting response object into rawtext
	rawResponse, err := json.Marshal(response)
	if err != nil {

		travelData.ErrorLog(request.ServiceRequest.Transactioninfo.RequestID, "error", err.Error())
	}

	// responselog in file
	travelData.ResponseLog(request.ServiceRequest.Transactioninfo.RequestID, string(rawResponse))

	//Saving response of specific api against requestid
	travelData.InsertResponse(requestID, strings.Replace(string(rawResponse), "  ", "", -1))
	travelData.ResponseLog(request.ServiceRequest.Transactioninfo.RequestID, string(rawResponse))
	return response
}
