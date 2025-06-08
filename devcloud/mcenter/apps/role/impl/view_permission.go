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
	"github.com/infraboard/modules/iam/apps/view"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/role"
	"gorm.io/gorm"
)

func (i *RoleServiceImpl) QueryViewPermission(ctx context.Context, in *role.QueryViewPermissionRequest) ([]*role.ViewPermission, error) {
	query := datasource.DBFromCtx(ctx).Model(&role.ViewPermission{})
	if len(in.RoleIds) > 0 {
		query = query.Where("role_id in ?", in.RoleIds)
	}
	if len(in.ViewPermissionIds) > 0 {
		query = query.Where("id in ?", in.ViewPermissionIds)
	}

	var perms []*role.ViewPermission
	if err := query.Order("created_at desc").Find(&perms).Error; err != nil {
		return nil, err
	}

	return perms, nil
}

func (i *RoleServiceImpl) AddViewPermission(ctx context.Context, in *role.AddViewPermissionRequest) ([]*role.ViewPermission, error) {
	if err := in.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate add view permission error, %s", err)
	}

	var perms []*role.ViewPermission
	if err := datasource.DBFromCtx(ctx).Transaction(func(tx *gorm.DB) error {
		for i := range in.Items {
			item := in.Items[i]
			perm := role.NewViewPermission(in.RoleId, item)
			tx.Save(perm)
			perms = append(perms, perm)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return perms, nil
}

func (i *RoleServiceImpl) RemoveViewPermission(ctx context.Context, in *role.RemoveViewPermissionRequest) ([]*role.ViewPermission, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	perms, err := i.QueryViewPermission(ctx, role.NewQueryViewPermissionRequest().AddRoleId(in.RoleId).AddPermissionId(in.ViewPermissionIds...))
	if err != nil {
		return nil, err
	}

	if err := datasource.DBFromCtx(ctx).Where("role_id = ?", in.RoleId).
		Where("id in ?", in.ViewPermissionIds).Delete(&role.ViewPermission{}).Error; err != nil {
		return nil, err
	}

	return perms, err
}

func (i *RoleServiceImpl) QueryMatchedPage(ctx context.Context, in *role.QueryMatchedPageRequest) (*types.Set[*view.Menu], error) {
	// 现在就一个查询参数是 id 先这么查询 得知道这个 id 是什么的 id
	// set := types.New[*view.Menu]()

	return nil, nil
}
