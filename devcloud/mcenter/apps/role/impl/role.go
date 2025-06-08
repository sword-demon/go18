// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package impl

import (
	"context"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/infraboard/mcube/v2/types"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/role"
	"gorm.io/gorm"
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
	set := types.New[*role.Role]()

	query := datasource.DBFromCtx(ctx).Model(&role.Role{})
	if len(in.RoleIds) > 0 {
		query = query.Where("id in ?", in.RoleIds)
		in.PageSize = uint64(len(in.RoleIds))
	}
	if err := query.Count(&set.Total).Error; err != nil {
		return nil, err
	}

	if err := query.Order("created_at desc").
		Offset(int(in.ComputeOffset())).
		Limit(int(in.PageSize)).
		Find(&set.Items).Error; err != nil {
		return nil, err
	}

	return set, nil
}

func (i *RoleServiceImpl) DescribeRole(ctx context.Context, in *role.DescribeRoleRequest) (*role.Role, error) {
	query := datasource.DBFromCtx(ctx)

	ins := role.NewRole()
	if err := query.Where("id = ?", in.Id).First(ins).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewNotFound("role %d not found", in.Id)
		}
		return nil, err
	}

	// 查询出当前角色关联的 api 权限
	permission, err := i.QueryApiPermission(ctx, role.NewQueryApiPermissionRequest().AddRoleId(in.Id))
	if err != nil {
		return nil, err
	}
	ins.ApiPermissions = permission

	return ins, err
}

func (i *RoleServiceImpl) UpdateRole(ctx context.Context, in *role.UpdateRoleRequest) (*role.Role, error) {
	descReq := role.NewDescribeRoleRequest()
	descReq.SetId(in.Id)
	ins, err := i.DescribeRole(ctx, descReq)
	if err != nil {
		return nil, err
	}

	ins.CreateRoleRequest = in.CreateRoleRequest
	return ins, datasource.DBFromCtx(ctx).Where("id = ?", in.Id).Updates(ins).Error
}

func (i *RoleServiceImpl) DeleteRole(ctx context.Context, in *role.DeleteRoleRequest) (*role.Role, error) {
	descReq := role.NewDescribeRoleRequest()
	descReq.SetId(in.Id)
	ins, err := i.DescribeRole(ctx, descReq)
	if err != nil {
		return nil, err
	}

	return ins, datasource.DBFromCtx(ctx).Where("id = ?", in.Id).Delete(&role.Role{}).Error
}
