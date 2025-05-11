// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package exception_test

import (
	"github.com/sword-demon/go18/book/v3/exception"
	"testing"
)

func CheckIsError() error {
	return exception.ErrNotFound("book %d not found", 1)
}

func TestException(t *testing.T) {
	err := CheckIsError()
	t.Log(err)

	if v, ok := err.(*exception.ApiException); ok {
		t.Log(v.Code)
		t.Log(v.String())
	}

	t.Log(exception.IsApiException(err, exception.CodeNotFound))
}
