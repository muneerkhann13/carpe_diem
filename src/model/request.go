package model

import "time"

type Request struct {
	Id          int64     `json:"request_id"`
	Service     string    `json:"service"`
	Mdn         int64     `json:"mdn"`
	Password    string    `json:"password,omitempty"`
	FirstName   string    `json:"first_name,omitempty"`
	MiddleName  string    `json:"middle_name,omitempty"`
	LastName    string    `json:"last_name,omitempty"`
	Address     string    `json:"address,omitempty"`
	DOB         string    `json:"dob,omitempty"`
	Email       string    `json:"email,omitempty"`
	Lat         string    `json:"lat,omitempty"`
	Long        string    `json:"long,omitempty"`
	ShopName    string    `json:"shop_name,omitempty"`
	Category    string    `json:"category,omitempty"`
	IsMerchant  bool      `json:"is_merchant,omitempty"`
	MerchantMdn int64     `json:"merchant_mdn,omitempty"`
	OfferId     int64     `json:"offer_id,omitempty"`
	Description string    `json:"desc,omitempty"`
	Title       string    `json:"title,omitempty"`
	Expire_at   time.Time `json:"expire_at,omitempty"`
}
