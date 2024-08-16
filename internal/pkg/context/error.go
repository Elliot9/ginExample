package context

type CustomError interface {
	BusinessCode() int
	HTTPCode() int
	Message() string
}

type customError struct {
	httpCode     int
	businessCode int
	message      string
}

func (err *customError) BusinessCode() int {
	return err.businessCode
}

func (err *customError) HTTPCode() int {
	return err.httpCode
}

func (err *customError) Message() string {
	return err.message
}

func Error(httpCode, businessCode int, message string) CustomError {
	return &customError{
		httpCode:     httpCode,
		businessCode: businessCode,
		message:      message,
	}
}
