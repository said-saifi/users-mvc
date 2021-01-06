package errors

type RestErr struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Status:  400,
		Code:    "bad_request",
		Message: message,
	}
}
