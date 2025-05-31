// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package impl

import (
	"context"
	"errors"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/infraboard/mcube/v2/types"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/user"
	"gorm.io/gorm"
)

func (i *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.User, error) {
	// 校验用户参数
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// 生成一个 user 对象
	entity := user.NewUser(req)

	if err := datasource.DBFromCtx(ctx).Create(entity).Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (i *UserServiceImpl) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.User, error) {
	// 查找用户是否存在
	u, err := i.DescribeUser(ctx, user.NewDescribeUserRequestById(req.Id))
	if err != nil {
		return nil, err
	}

	return u, datasource.DBFromCtx(ctx).Where("id = ?", req.Id).Delete(&user.User{}).Error
}

func (i *UserServiceImpl) DescribeUser(ctx context.Context, req *user.DescribeUserRequest) (*user.User, error) {
	query := datasource.DBFromCtx(ctx)

	// 构造查询条件
	switch req.DescribeBy {
	case user.DescribeByID:
		query = query.Where("id = ?", req.DescribeValue)
	case user.DescribeByUsername:
		query = query.Where("user_name = ?", req.DescribeValue)
	}

	entity := &user.User{}
	if err := query.First(entity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.NewNotFound("user %s not found", req.DescribeValue)
		}
		return nil, err
	}

	// 数据库里面存储的是 hash
	entity.SetIsHashed()
	return entity, nil
}

func (i *UserServiceImpl) QueryUser(ctx context.Context, req *user.QueryUserRequest) (*types.Set[*user.User], error) {
	set := types.New[*user.User]()
	query := datasource.DBFromCtx(ctx).Model(&user.User{})

	// 查询总量
	err := query.Count(&set.Total).Error
	if err != nil {
		return nil, err
	}

	err = query.
		Order("created_at desc").
		Offset(int(req.ComputeOffset())).
		Limit(int(req.PageSize)).
		Find(&set.Items).Error
	if err != nil {
		return nil, err
	}

	return set, nil
}
