// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package impl_test

import (
	"testing"

	"github.com/sword-demon/go18/devcloud/mcenter/apps/user"
)

func TestQueryUser(t *testing.T) {
	req := user.NewQueryUserRequest()
	set, err := impl.QueryUser(ctx, req)
	if err != nil {
		t.Fatalf("QueryUser err: %v", err)
	}
	t.Log(set)
}

func TestCreateAdminUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.UserName = "admin"
	req.Password = "123456"
	req.EnabledApi = true
	req.IsAdmin = true
	createUser, err := impl.CreateUser(ctx, req)
	if err != nil {
		t.Fatalf("CreateUser err: %v", err)
	}
	t.Log(createUser)
}

func TestDescribeUser(t *testing.T) {
	req := user.NewDescribeUserRequestById("1")
	describeUser, err := impl.DescribeUser(ctx, req)
	if err != nil {
		t.Fatalf("DescribeUser err: %v", err)
	}
	t.Log(describeUser)
}
