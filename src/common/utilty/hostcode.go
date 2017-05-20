package utility

const (
	Success            = "0"
	SomeThingWentWrong = "1"
)

var statusCodes = map[string]string{
	Success:            "Success",
	SomeThingWentWrong: "Something Went Wrong",
}

func StatusCodes(code string) string {
	return statusCodes[code]
}
