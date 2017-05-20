package travelModel

type Response struct {
	Service         string          `json:"_service"`
	Version         string          `json:"_version"`
	ServiceResponse ServiceResponse `json:"service_response"`
}

type ServiceResponse struct {
	ResponseInfo    ResponseInfo    `json:"response_info"`
	TransactionInfo TransactionInfo `json:"transaction_info"`
	User            User            `json:"user"`
	Bus             interface{}     `json:"bus,omitempty"`
	Hotel           interface{}     `json:"hotel,omitempty"`
	Flight          interface{}     `json:"flight,omitempty"`
	PartnerInfo     PartnerInfo     `json:"partner_info"`
	Params          Params          `json:"params,omitempty"`
}
