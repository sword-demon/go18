// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package impl

import (
	"context"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/infraboard/mcube/v2/types"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/role"
)

func (i *RoleServiceImpl) CreateRole(ctx context.Context, in *role.CreateRoleRequest) (*role.Role, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	ins := role.NewRole()
	ins.CreateRoleRequest = *in
	if err := datasource.DBFromCtx(ctx).Create(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *RoleServiceImpl) QueryRole(ctx context.Context, in *role.QueryRoleRequest) (*types.Set[*role.Role], error) {
	//TODO implement me
	panic("implement me")
}

func (i *RoleServiceImpl) DescribeRole(ctx context.Context, in *role.DescribeRoleRequest) (*role.Role, error) {
	//TODO implement me
	panic("implement me")
}

func (i *RoleServiceImpl) UpdateRole(ctx context.Context, in *role.UpdateRoleRequest) (*role.Role, error) {
	//TODO implement me
	panic("implement me")
}

func (i *RoleServiceImpl) DeleteRole(ctx context.Context, in *role.DeleteRoleRequest) (*role.Role, error) {
	//TODO implement me
	panic("implement me")
}
