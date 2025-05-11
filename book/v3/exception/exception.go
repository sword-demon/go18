// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package exception

import (
	"errors"
	"github.com/infraboard/mcube/v2/tools/pretty"
)

type ApiException struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	HttpCode int    `json:"-"`
}

func NewApiException(code int, message string) *ApiException {
	return &ApiException{Code: code, Message: message}
}

func (e *ApiException) Error() string {
	return e.Message
}

func (e *ApiException) String() string {
	return pretty.ToJSON(e)
}

func (e *ApiException) WithMessage(msg string) *ApiException {
	e.Message = msg
	return e
}

func (e *ApiException) WithHttpCode(code int) *ApiException {
	e.HttpCode = code
	return e
}

// IsApiException 异常的对比，通过业务码进行对比
func IsApiException(err error, code int) bool {
	var apiErr *ApiException
	if errors.As(err, &apiErr) {
		return apiErr.Code == code
	}
	return false
}
