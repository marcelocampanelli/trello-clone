package utils

type CustomError struct {
	Status  int
	Message string
}

func NewCustomRerror(status int, message string) *CustomError {
	return &CustomError{
		Status:  status,
		Message: message,
	}
}
