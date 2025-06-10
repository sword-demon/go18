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

func TestQueryApiPermission(t *testing.T) {
	req := role.NewQueryApiPermissionRequest()
	req.AddRoleId(1)
	set, err := impl.QueryApiPermission(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(set)
}

func TestAddApiPermission(t *testing.T) {
	req := role.NewAddApiPermissionRequest(1)
	req.Add(role.NewResourceActionApiPermissionSpec("devcloud", "user", "list"))
	set, err := impl.AddApiPermission(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestQueryMatchedEndpoint(t *testing.T) {
	req := role.NewQueryMatchedEndpointRequest()
	req.Add(1)
	set, err := impl.QueryMatchedEndpoint(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestRemoveApiPermission(t *testing.T) {
	req := role.NewRemoveApiPermissionRequest(2)
	req.Add(2)
	set, err := impl.RemoveApiPermission(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(set)
}
