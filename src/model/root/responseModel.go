package travelModel

type ResponseInfo struct {
	HostCode        string `json:"host_code"`
	HostDescription string `json:"host_description"`
}

type TransactionData struct {
	Amount           string `json:"amount"`
	CurrencyCode     string `json:"currency_code"`
	TransactionID    string `json:"transaction_id"`
	RemainingBalance string `json:"remaining_balance"`
	BookingID        string `json:"booking_id"`
}
