package echoutil

import (
	"strconv"

	"github.com/labstack/echo"
	restgostruct "github.com/incompetent-developer/restgo-struct"
)

// SuccessResponse is response with Status OK (200)
func SuccessResponse(echoContext echo.Context, code int, result interface{}) error {
	return echoContext.JSONPretty(code, result, "  ")
}

// ErrorResponse is response with Status not OK (!=200)
func ErrorResponse(echoContext echo.Context, code int, message, systemErrorCode string) error {
	responseError := restgostruct.ResponseError{
		Code:            code,
		Message:         message,
		SystemErrorCode: systemErrorCode,
	}
	return echoContext.JSONPretty(code, responseError, "  ")
}

// ConvertPageLimit is convert from string to integer
func ConvertPageLimit(pageString, limitString string) (page, limit int) {
	var err error
	page, err = strconv.Atoi(pageString)
	if err != nil {
		page = 1
	}
	limit, err = strconv.Atoi(limitString)
	if err != nil {
		limit = 10
	}
	return page, limit
}
