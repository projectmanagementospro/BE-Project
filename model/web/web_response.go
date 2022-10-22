package web

type WebResponseSuccess struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type WebResponseError struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Errors interface{} `json:"errors"`
}

func NewWebSuccessResponse(code int, status string, data interface{}) WebResponseSuccess {
	return WebResponseSuccess{
		Code:   code,
		Status: status,
		Data:   data,
	}
}

func NewWebErrorResponse(code int, status string, err interface{}) WebResponseError {
	return WebResponseError{
		Code:   code,
		Status: status,
		Errors: err,
	}
}
