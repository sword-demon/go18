// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package impl_test

import (
	"github.com/sword-demon/go18/devcloud/mcenter/apps/token"
	"testing"
)

func TestIssueToken(t *testing.T) {
	req := token.NewIssueTokenRequest()
	req.IssueByPassword("admin", "123456")
	req.Source = token.SourceWeb
	set, err := svc.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestQueryToken(t *testing.T) {
	req := token.NewQueryTokenRequest()
	set, err := svc.QueryToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}
