package model

type Request struct {
	Id       int64  `json:"request_id"`
	Service  string `json:"service"`
	Mdn      string `json:"mdn"`
	Password string `json:"password"`
}
