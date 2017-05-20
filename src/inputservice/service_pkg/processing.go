package service_pkg

import (
	"fmt"
	"model"
	"strings"
	"utilityData"

	"common/utility"

	"github.com/go-errors/errors"
)

func GetResponse(request model.Request, jsonbody []uint8) model.Response {

	response := model.Response{}

	//Validating request
	//hostCode, hostDescription, ok := validations(request)

	hostCode := "0"
	hostDescription := "SomeThingWentWrong"
	var err error
	err = nil
	// insert request in file log
	requestID := request.Id
	ok := true
	//Saving request with parameters
	//	requestID, err := utilityData.InsertRequest(request.Mdn, request.Service, strings.Replace(string(jsonbody), "  ", "", -1), "", "", request.ServiceRequest.Partnerinfo.Code, request.ServiceRequest.Serviceinfo.Category)
	//utilityData.RequestLog(request.ServiceRequest.Transactioninfo.RequestID, request.ServiceRequest.User.Mdn, request.Service, strings.Replace(string(jsonbody), "  ", "", -1), "", "", request.ServiceRequest.Partnerinfo.Code, request.ServiceRequest.Serviceinfo.Category)
	// When comes error during insert request

	if err != nil {
		utilityData.InsertErrorLogs(requestID, fmt.Sprint(errors.Errorf(err.Error()).ErrorStack()))
		hostCode = utility.SomeThingWentWrong
		hostDescription = utility.StatusCodes(utility.SomeThingWentWrong)
	} else {
		if ok {

			switch strings.ToUpper(request.Category) {

			case "userservice":
				fmt.Println(request)
				//var response1 model.Response
				response.Code = "200"
				response.Description = "aa gai"
				response.IsMerchant = true
				response.Token = "token"
				//response = response1
				//response = travelHotel.GetHotelResponse(requestID, request.Service, jsonbody)

			case "nearbyservice":
				//response = travelFlight.GetFlightResponse(request.ServiceRequest.Transactioninfo.RequestID, requestID, request.Service, jsonbody)

			case "offerservice":
				//response = travelBus.GetBusResponse(request.Service, jsonbody)

			default:
				hostCode = utility.ServiceCategoryDoesntExists
				hostDescription = utility.StatusCodes(utility.ServiceCategoryDoesntExists)
			}

			// When validation failed
		} else {
			hostCode = utility.Success
			hostDescription = utility.StatusCodes(utility.Success)
		}
	}

	response.Code = hostCode
	response.Description = hostDescription
	return response
}
