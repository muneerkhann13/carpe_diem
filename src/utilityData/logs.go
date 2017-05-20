package utilityData

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

/*-------------------------Insert raw request logs--------------------*/

func InsertRawRequest(encryptedText, plainText, headers string) error {
	stmt, err := DBtravel.Prepare("insert into rawrequest(encrypted_text, plain_text, headers, request_date) values(?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(encryptedText, plainText, headers, time.Now())
	if err != nil {
		return err
	}

	return err
}

/*-------------------------Insert error logs data--------------------*/

func InsertErrorLogs(requestId int64, errorDetails string) {

	stmt, err := DBtravel.Prepare("insert into errorlogs(request_id, error_details,error_date) values(?,?,?)")
	if err != nil {
		ErrorInFile(err)
	}
	_, err = stmt.Exec(requestId, errorDetails, time.Now())
	if err != nil {
		ErrorInFile(err)
	}

}

/*-------------------------Insert request logs--------------------*/

func InsertRequest(mobile, serviceName, text, description, methodName, partnerName, moduleName string) (int64, error) {

	stmt, err := DBtravel.Prepare("insert into request(mobile, service_name, _text, _description, created_date, method_name, partner_name, module_name) values(?,?,?,?,?,?,?,?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(mobile, serviceName, text, description, time.Now(), methodName, partnerName, moduleName)

	if err != nil {
		return 0, err
	}

	id, _ := res.LastInsertId()
	return id, err
}

/*-------------------------Insert response logs--------------------*/

func InsertResponse(requestId int64, text string) error {

	stmt, err := DBtravel.Prepare("insert into response(request_id, _text, _date) values(?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(requestId, text, time.Now())
	if err != nil {
		return err
	}
	return err
}

/*-------------------------Maintain partner's logs--------------------*/

func InsertPartnerLog(requestID, serviceName, requestURL, requestMethod, posted, responseText, partnerName, moduleName string, isSuccess bool) error {

	stmt, err := DBtravel.Prepare(`insert into partnerrequestresponse (request_id, service_name, request_url, request_method,
	posted, request_date, response_text, partner_name, module_name, is_success)
	values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)

	if err != nil {
		return err
	}

	res, err := stmt.Exec(requestID, serviceName, requestURL, requestMethod, posted, time.Now(), responseText, partnerName, moduleName)

	if err != nil {
		return err
	}

	res.LastInsertId()
	return err
}

/*-------------------------Checkerror method--------------------*/

func ErrorInFile(err error) {
	var file *os.File
	path, _ := filepath.Abs("./content/errors")
	path = path + "/errDB.txt"
	if _, pathErr := os.Stat(path); pathErr != nil {
		file, _ = os.Create(path)
		file.WriteString(err.Error() + "," + time.Now().String())
		file.WriteString("\n")
	} else {
		file, _ = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
		file.WriteString(err.Error())
		file.WriteString("\n")
	}
	file.Close()
}

/*
	Author : Shivendra Pratap Singh
	CreatedDate : 05-Dec-2016
	Purpose : Take Care of errors with corresponding error codes
	Parameters : {
	"input":"various Parameters,
	 "output":"Query Response"
							}
	Module : Data
	LastUpdate : { "name": "Shivendra Pratap Singh", "date":"22 December 2016" }
*/

func GetTempBooking(provBookingID, mobileNumber string) (string, error) {

	var queryData string
	err := DBtravel.QueryRow(`select request_data from tempbooking where booking_id='` + provBookingID + `' and mobile_no='` + mobileNumber + `' limit 1`).Scan(&queryData)
	return queryData, err
}

func SaveProvisionalData(bookingId, requestId, mobileNo, requestData, partnerName, moduleName, responseData, partnerInfo string) {

	query := `insert into tempbooking(booking_id, request_id, mobile_no, request_data, created_date, partner_name, module_name,
			  response_data, passenger_info) values(?, ?, ?, ?, curdate(), ?, ?, ?, ?)`

	stmt, err := DBtravel.Prepare(query)

	if err != nil {
		fmt.Println(err)
	}

	res, err := stmt.Exec(bookingId, requestId, mobileNo, requestData, partnerName, moduleName, responseData, partnerInfo)

	id, err := res.LastInsertId()

	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(id)
}

func UpdatingBooking(mobile, tempBookingId, confirmBookingId, mode, txnNumber, amount string, bookingStatus bool) {

	query := `update travelaggregator.bookings set txn_number=?, payment_via = ?, booking_status=1, booking_amount=?, payment_status= ?, confirm_booking_id = ?
			  where mobile=? and temp_booking_id = ?`
	fmt.Println(query)
}

func SaveBooking(mobile string, tempBookingId string, txn_number string, payment_via string, responseMessage string, booking_status string, booking_amount string, payment_status bool, confirm_booking_id string, partner string, module string) error {

	query := `insert into bookings(mobile, temp_booking_id,txn_number,payment_via,booking_message,booking_status,booking_date,booking_amount, payment_status, confirm_booking_id, partner, module) values(?, ?, ?, ?,?,?,current_date(), ?, ?,?,?,?)`
	stmt, err := DBtravel.Prepare(query)
	if err != nil {
		fmt.Println("err1", err)
		return err
	}
	_, err = stmt.Exec(mobile, tempBookingId, txn_number, payment_via, responseMessage, booking_status, booking_amount, payment_status, confirm_booking_id, partner, module)
	if err != nil {
		fmt.Println("err2", err)
		return err
	}
	return err
}

func GetBookingStatus(mobile, tempBookingId string) error {

	query := `select booking_status  from bookings where mobile = '9717212781' and temp_booking_id = '53894dfjk'`
	stmt, err := DBtravel.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(mobile, tempBookingId)
	if err != nil {
		return err
	}

	return err
}

func SaveFailureTxn(mobile, model, txnNumber, requestId, amount, bookingId, responseCode, responseMessage, remarks, partner, module string) error {

	query := `insert into failuretransaction(mobile, _mode, txn_number, request_id, amount, booking_id, response_code, response_message, remarks,
			  created_date, partner, module) values(?, ?, ?, ?, ?, ?, ?, ?, ?, current_date(), ?, ?)`

	stmt, err := DBtravel.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(mobile, model, txnNumber, requestId, amount, bookingId, responseCode, responseMessage, remarks, partner, module)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	return err
}
