// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package impl

import (
	"context"
	"github.com/infraboard/mcube/v2/types"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/endpoint"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/role"
)

func (i *RoleServiceImpl) QueryApiPermission(ctx context.Context, in *role.QueryApiPermissionRequest) ([]*role.ApiPermission, error) {
	//TODO implement me
	panic("implement me")
}

func (i *RoleServiceImpl) AddApiPermission(ctx context.Context, in *role.AddApiPermissionRequest) ([]*role.ApiPermission, error) {
	//TODO implement me
	panic("implement me")
}

func (i *RoleServiceImpl) RemoveApiPermission(ctx context.Context, in *role.RemoveApiPermissionRequest) ([]*role.ApiPermission, error) {
	//TODO implement me
	panic("implement me")
}

func (i *RoleServiceImpl) QueryMatchedEndpoint(ctx context.Context, in *role.QueryMatchedEndpointRequest) (*types.Set[*endpoint.Endpoint], error) {
	//TODO implement me
	panic("implement me")
}
