package sendandrecive

type SendDataSucusses struct {
	// status code
	StatusCode int `json:"code"`
	// header info
	HeaderInfo string `json:"hdinfo"`
	// messsage
	AnyMessage string `json:"msg"`
	// data
	Data any `json:"mydata"`
}

type SendDataFail struct {
	// status code
	StatusCode int
	// header info
	HeaderInfo string
	// messsage
	AnyMessage string
	// data
	Data any
}
