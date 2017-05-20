package model

type Response struct {
	Token       string `json:"token"`
	IsMerchant  string `json:"is_merchant"`
	Code        string `json:"code"`
	Description string `json:"description"`
}
