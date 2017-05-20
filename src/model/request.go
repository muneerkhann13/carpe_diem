package model

type Request struct {
	Service  string `json:"service"`
	Mdn      string `json:"mdn"`
	Password string `json:"password"`
}
