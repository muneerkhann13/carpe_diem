package utility

const (
	Success                       = "0"
	SomeThingWentWrong            = "1"
	ServiceCategoryDoesntExists   = "3"
	InvalidRequesthaveSent        = "3002"
	ServiceInfoObjectNotInRequest = "3003"
	ServiceCategoryCannotBeBlank  = "3004"
	InstrumentCannotBeBlank       = "3005"
	RequesterCannotBeBlank        = "3006"
	MobileNumberCannotBeBlank     = "3007"
	InvalidServiceName            = "3008"
)

var statusCodes = map[string]string{
	Success:                     "Success",
	SomeThingWentWrong:          "Something Went Wrong",
	ServiceCategoryDoesntExists: "Service Category Doesnt Exists",
	"3002": "Invalid Request have sent",
	"3003": "Service info object not in request",
	"3004": "Service Category cannot be blank",
	"3005": "Instrument cannot be blank",
	"3006": "Requester cannot be blank",
	"3007": "Mobile number cannot be blank",
	"3008": "Invalid Service name",
}

func StatusCodes(code string) string {
	return statusCodes[code]
}
