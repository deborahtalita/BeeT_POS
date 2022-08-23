package helper

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"msg"`
	Errors  interface{} `json:"errors,omitempty"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

func BuildErrorResponse(message string, err string, data interface{}) Response {
	res := Response{
		Status:  false,
		Message: message,
		Errors:  err,
		Data:    data,
	}
	return res
}
