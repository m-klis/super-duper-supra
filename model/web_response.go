package model

type WebResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message interface{} `json:"message"`
}

func NewWebResponse() WebResponse {
	return WebResponse{
		Data:    "empty",
		Message: "as expected",
	}
}
