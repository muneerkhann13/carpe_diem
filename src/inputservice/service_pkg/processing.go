package service_pkg

import (
	"fmt"
	"model"
	"strings"
	"utilityData"

	"common/configuration"

	"github.com/go-errors/errors"
)

func GetResponse(request model.Request, jsonbody []uint8) model.Response {

	response := model.Response{}

	//Validating request
	//hostCode, hostDescription, ok := validations(request)

	// insert request in file log

	if err != nil {
		utilityData.InsertErrorLogs(requestID, fmt.Sprint(errors.Errorf(err.Error()).ErrorStack()))
		hostCode = configuration.SomeThingWentWrong
		hostDescription = configuration.StatusCodes(configuration.SomeThingWentWrong)
	} else {
		if ok {

			switch strings.ToUpper(request.ServiceRequest.Serviceinfo.Category) {

			case "userservice":
				fmt.Println(request)
				var response1 model.Response
				response1.Code = "200"
				response1.Description = "aa gai"
				response1.IsMerchant = "yes"
				response1.Token = "token"
				response = response1
				//response = travelHotel.GetHotelResponse(requestID, request.Service, jsonbody)

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
			response.ServiceResponse.ResponseInfo.HostCode = "xxxx"
			response.ServiceResponse.ResponseInfo.HostDescription = "xxxx"
		}
	}

	// //Converting response object into rawtext
	// rawResponse, err := json.Marshal(response)
	// if err != nil {

	// 	travelData.ErrorLog(request.ServiceRequest.Transactioninfo.RequestID, "error", err.Error())
	// }

	// // responselog in file
	// travelData.ResponseLog(request.ServiceRequest.Transactioninfo.RequestID, string(rawResponse))

	// //Saving response of specific api against requestid
	// travelData.InsertResponse(requestID, strings.Replace(string(rawResponse), "  ", "", -1))
	// travelData.ResponseLog(request.ServiceRequest.Transactioninfo.RequestID, string(rawResponse))
	return response
}
