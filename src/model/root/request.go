package travelModel

type Request struct {
	Service        string         `json:"_service"`
	Version        string         `json:"_version"`
	ServiceRequest ServiceRequest `json:"service_request"`
}

type ServiceRequest struct {
	Channelinfo     ChannelInfo     `json:"channel_info"`
	Deviceinfo      DeviceInfo      `json:"device_info"`
	Transactioninfo TransactionInfo `json:"transaction_info"`
	User            User            `json:"user"`
	Partnerinfo     PartnerInfo     `json:"partner_info"`
	Serviceinfo     ServiceInfo     `json:"service_info"`
	Params          Params          `json:"params"`
}
