package model

import "time"

type User struct {
	Mdn string `json:"mdn"`
}

type PartnerInfo struct {
	Code string `json:"code"`
}

type TransactionInfo struct {
	TimeStamp      time.Time `json:"time_stamp"`
	RequestID      string    `json:"request_id"`
	TxnDescription string    `json:"txn_description,omitempty"`
}

type Params struct {
	Param1 string `json:"param_1,omitempty"`
	Param2 string `json:"param_2,omitempty"`
	Param3 string `json:"param_3,omitempty"`
	Param4 string `json:"param_4,omitempty"`
	Param5 string `json:"param_5,omitempty"`
}
