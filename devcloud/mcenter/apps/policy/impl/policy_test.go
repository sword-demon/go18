// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package impl_test

import (
	"github.com/sword-demon/go18/devcloud/mcenter/apps/policy"
	"testing"
)

func TestCreatePolicy(t *testing.T) {
	req := policy.NewCreatePolicyRequest()
	req.SetNamespaceId(1)
	req.UserId = 1
	req.RoleId = 1
	set, err := impl.CreatePolicy(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestQueryPolicy(t *testing.T) {
	req := policy.NewQueryPolicyRequest()
	req.WithUser = true
	req.WithRole = true
	req.WithNamespace = true
	set, err := impl.QueryPolicy(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}
