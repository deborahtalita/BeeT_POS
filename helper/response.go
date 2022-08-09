package helper

import "strings"

//Response untuk return JSON
type Response struct{
	Status bool `json:"status"`
	Message string `json:"message"`
	Errors interface{} `json:"errors"`
	Data interface{} `json:"data"`
}

//EmptyObj untuk object yang kosong atau null
type EmptyObj struct{
	
}

//BuildResponse method yang memberikan nilai success dynamic response
func BuildResponse(status bool, message string, data interface{}) Response{
	res := Response{
		Status: status,
		Message: message,
		Errors: nil,
		Data: data,
	}

	return res 
}


//BuildErrorResponse method yang memberikan nilai failure dynamic repsonse
func BuildErrorResponse(message string, err string, data interface{})Response{
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status: false,
		Message: message,
		Errors : splittedError,
		Data: data,
	}
	return res
}
