package controller

// Response is a common response content for ticktok application
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

// ResponseError creates a common response content represents error
func ResponseError(err error) Response {
	return Response{
		StatusCode: 1,
		StatusMsg:  err.Error(),
	}
}

// ResponseOK creates a common response content represents success
func ResponseOK() Response {
	return Response{
		StatusCode: 0,
	}
}
