package utilityData

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log4"
)

var Logg log4.LogSturct
var delimeter = string("::")

func RawRequestLog(encryptedText, plainText, header string) {

	txtbyte := []byte(plainText)
	txtbuff := bytes.Buffer{}
	err := json.Compact(&txtbuff, txtbyte)
	if err != nil {
		Logg.WriteLog("raw_request_log", "info", fmt.Sprint(delimeter, encryptedText, delimeter, plainText, delimeter, header))
	} else {
		Logg.WriteLog("raw_request_log", "info", fmt.Sprint(delimeter, encryptedText, delimeter, txtbuff.String(), delimeter, header))
	}
	return
}

func RequestLog(requestID, mdn, serviceName, requestText, dxn, methodName, partnerName, moduleName string) {

	txtbyte := []byte(requestText)
	txtbuff := bytes.Buffer{}
	err := json.Compact(&txtbuff, txtbyte)
	if err != nil {
		Logg.WriteLog("request_log", "info", fmt.Sprint(delimeter, requestID, delimeter, mdn, delimeter, serviceName, delimeter, requestText, delimeter, dxn, delimeter, methodName, delimeter, partnerName, delimeter, moduleName))
	} else {
		Logg.WriteLog("request_log", "info", fmt.Sprint(delimeter, requestID, delimeter, mdn, delimeter, serviceName, delimeter, txtbuff.String(), delimeter, dxn, delimeter, methodName, delimeter, partnerName, delimeter, moduleName))
	}
	return
}

func ResponseLog(requestID, responseText string) {
	txtbyte := []byte(responseText)
	txtbuff := bytes.Buffer{}
	err := json.Compact(&txtbuff, txtbyte)
	if err != nil {
		Logg.WriteLog("response_log", "info", fmt.Sprint(delimeter, requestID, delimeter, responseText))
	} else {
		Logg.WriteLog("response_log", "info", fmt.Sprint(delimeter, requestID, delimeter, txtbuff.String()))
	}
	return
}

func ErrorLog(requestID, lvl, errorMessage string) {
	Logg.WriteLog("error_log", lvl, fmt.Sprint(delimeter, requestID, delimeter, errorMessage))
	return
}

func PartnerRequestResponse(requestID, serviceName, requestURL, requestMethod, posted, responseText, partnerName, moduleName string, isSuccess bool) {
	txtbyte := []byte(responseText)
	txtbuff := bytes.Buffer{}
	err := json.Compact(&txtbuff, txtbyte)
	if err != nil {
		Logg.WriteLog("partner_log", "info", fmt.Sprint(delimeter, requestID, delimeter, serviceName, delimeter, requestURL, delimeter, requestMethod, delimeter, posted, delimeter, responseText, delimeter, partnerName, delimeter, moduleName, delimeter, isSuccess))
	} else {
		Logg.WriteLog("partner_log", "info", fmt.Sprint(delimeter, requestID, delimeter, serviceName, delimeter, requestURL, delimeter, requestMethod, delimeter, posted, delimeter, txtbuff.String(), delimeter, partnerName, delimeter, moduleName, delimeter, isSuccess))
	}
	return
}
