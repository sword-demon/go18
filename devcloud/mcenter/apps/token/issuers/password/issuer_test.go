// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package password_test

import (
	"context"
	"testing"

	"github.com/sword-demon/go18/devcloud/mcenter/apps/token"
	"github.com/sword-demon/go18/devcloud/mcenter/test"
)

func TestPasswordIssuer(t *testing.T) {
	issuer := token.GetIssuer(token.IssuerPassword)
	tk, err := issuer.IssueToken(context.Background(), token.NewIssueParameter().SetUsername("admin").SetPassword("123456"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

func init() {
	test.DevelopmentSet()
}
