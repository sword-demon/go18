// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package impl

import (
	"context"
	"github.com/infraboard/mcube/v2/types"
	"github.com/infraboard/modules/iam/apps/view"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/role"
)

func (i *RoleServiceImpl) QueryViewPermission(ctx context.Context, in *role.QueryViewPermissionRequest) ([]*role.ViewPermission, error) {
	//TODO implement me
	panic("implement me")
}

func (i *RoleServiceImpl) AddViewPermission(ctx context.Context, in *role.AddViewPermissionRequest) ([]*role.ViewPermission, error) {
	//TODO implement me
	panic("implement me")
}

func (i *RoleServiceImpl) RemoveViewPermission(ctx context.Context, in *role.RemoveViewPermissionRequest) ([]*role.ViewPermission, error) {
	//TODO implement me
	panic("implement me")
}

func (i *RoleServiceImpl) QueryMatchedPage(ctx context.Context, in *role.QueryMatchedPageRequest) (*types.Set[*view.Menu], error) {
	//TODO implement me
	panic("implement me")
}
