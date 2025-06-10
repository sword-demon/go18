// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package impl_test

import (
	"github.com/sword-demon/go18/devcloud/mcenter/apps/role"
	"testing"
)

func TestQueryRole(t *testing.T) {
	req := role.NewQueryRoleRequest()
	set, err := impl.QueryRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestDescribeRole(t *testing.T) {
	req := role.NewDescribeRoleRequest()
	req.SetId(1)
	ins, err := impl.DescribeRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestCreateAdminRole(t *testing.T) {
	req := role.NewCreateRoleRequest()
	req.Name = "admin"
	req.Description = "管理员"
	ins, err := impl.CreateRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
