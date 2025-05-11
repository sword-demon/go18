// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

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
