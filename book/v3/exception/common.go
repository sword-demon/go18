package exception

import (
	"fmt"
	"net/http"
)

func ErrServerInternal(format string, a ...any) *ApiException {
	return &ApiException{
		Code:     CodeServerError,
		Message:  fmt.Sprintf(format, a...),
		HttpCode: http.StatusInternalServerError,
	}
}

func ErrNotFound(format string, a ...any) *ApiException {
	return &ApiException{
		Code:     CodeNotFound,
		Message:  fmt.Sprintf(format, a...),
		HttpCode: http.StatusNotFound,
	}
}

func ErrParamsInvalid(format string, a ...any) *ApiException {
	return &ApiException{
		Code:     CodeParamInvalid,
		Message:  fmt.Sprintf(format, a...),
		HttpCode: http.StatusBadRequest,
	}
}
