package response 

type CorrectResponse struct {
	// http status code 
	StatusCode int `json:"status_code"`
	Data any `json:"data"`
}

type ErrorResponse struct{
	StatusCode int `json:"status_code"`
	Data any `json:"data"`
}