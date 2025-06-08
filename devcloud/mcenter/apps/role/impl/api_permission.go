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
	"github.com/sword-demon/go18/devcloud/mcenter/apps/endpoint"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/role"
	"gorm.io/gorm"
)

func (i *RoleServiceImpl) QueryApiPermission(ctx context.Context, in *role.QueryApiPermissionRequest) ([]*role.ApiPermission, error) {
	query := datasource.DBFromCtx(ctx).Model(&role.ApiPermission{})
	if len(in.RoleIds) > 0 {
		query = query.Where("role_id in ?", in.RoleIds)
	}

	var perms []*role.ApiPermission
	if err := query.Order("created_at desc").Find(&perms).Error; err != nil {
		return nil, err
	}

	return perms, nil
}

func (i *RoleServiceImpl) AddApiPermission(ctx context.Context, in *role.AddApiPermissionRequest) ([]*role.ApiPermission, error) {
	if err := in.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate add api permission error, %v", err)
	}

	var perms []*role.ApiPermission
	if err := datasource.DBFromCtx(ctx).Transaction(func(tx *gorm.DB) error {
		for i := range in.Items {
			item := in.Items[i]
			perm := role.NewApiPermission(in.RoleId, item)
			tx.Save(perm)
			perms = append(perms, perm)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return perms, nil
}

func (i *RoleServiceImpl) RemoveApiPermission(ctx context.Context, in *role.RemoveApiPermissionRequest) ([]*role.ApiPermission, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	perms, err := i.QueryApiPermission(ctx, role.NewQueryApiPermissionRequest().AddRoleId(in.RoleId).AddPermissionId(in.ApiPermissionIds...))
	if err != nil {
		return nil, err
	}

	if err := datasource.DBFromCtx(ctx).Where("role_id = ?", in.RoleId).
		Where("id in ?", in.ApiPermissionIds).
		Delete(&role.ApiPermission{}).Error; err != nil {
		return nil, err
	}

	return perms, nil
}

func (i *RoleServiceImpl) QueryMatchedEndpoint(ctx context.Context, in *role.QueryMatchedEndpointRequest) (*types.Set[*endpoint.Endpoint], error) {
	set := types.New[*endpoint.Endpoint]()

	// 查询角色的权限
	perms, err := i.QueryApiPermission(ctx, role.NewQueryApiPermissionRequest().AddRoleId(in.RoleIds...))
	if err != nil {
		return nil, err
	}

	// 查询服务的 endpoint 列表
	endpointReq := endpoint.NewQueryEndpointRequest()
	for _, perm := range perms {
		endpointReq.WithService(perm.Service)
	}

	endpoints, err := endpoint.GetService().QueryEndpoint(ctx, endpointReq)
	if err != nil {
		return nil, err
	}

	// 找出匹配的 API
	endpoints.ForEach(func(t *endpoint.Endpoint) {
		for _, perm := range perms {
			if perm.IsMatch(t) {
				if !endpoint.IsEndpointExist(set, t) {
					set.Add(t)
				}
			}
		}
	})

	return set, nil
}
