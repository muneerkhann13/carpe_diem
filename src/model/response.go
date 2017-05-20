package model

type Response struct {
	Token       string       `json:"token,omitempty"`
	IsMerchant  bool         `json:"is_merchant,omitempty"`
	Code        string       `json:"code"`
	Description string       `json:"description"`
	FirstName   string       `json:"first_name,omitempty"`
	MiddleName  string       `json:"middle_name,omitempty"`
	LastName    string       `json:"last_name,omitempty"`
	Address     string       `json:"address,omitempty"`
	DOB         string       `json:"dob,omitempty"`
	Email       string       `json:"email,omitempty"`
	Lat         string       `json:"lat,omitempty"`
	Long        string       `json:"long,omitempty"`
	ShopName    string       `json:"shop_name,omitempty"`
	Category    string       `json:"category,omitempty"`
	Offers      []Offer      `json:"offers,omitempty"`
	CashCentres []CashCentre `json:"cash_centre,omitempty"`
	Merchantss  []Merchant   `json:"merchants,omitempty"`
}

type Offer struct {
	Id          string `json:"id"`
	Description bool   `json:"desc"`
	Title       int64  `json:"title"`
	ExpireAt    string `json:"expire_at"`
}

type CashCentre struct {
	Mdn        string `json:"mdn"`
	FirstName  string `json:"first_name,omitempty"`
	MiddleName string `json:"middle_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	Address    string `json:"address,omitempty"`
	DOB        string `json:"dob,omitempty"`
	Email      string `json:"email,omitempty"`
	Lat        string `json:"lat"`
	Long       string `json:"long"`
	ShopName   string `json:"shop_name"`
	Category   string `json:"category"`
	Features   string `json:"features"`
}
type Merchant struct {
	Mdn        string `json:"mdn"`
	FirstName  string `json:"first_name,omitempty"`
	MiddleName string `json:"middle_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	Address    string `json:"address,omitempty"`
	DOB        string `json:"dob,omitempty"`
	Email      string `json:"email,omitempty"`
	Lat        string `json:"lat"`
	Long       string `json:"long"`
	ShopName   string `json:"shop_name"`
	Category   string `json:"category"`
}
