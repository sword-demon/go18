// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package user

import (
	"context"
	"github.com/infraboard/mcube/v2/http/request"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/types"
	"slices"
)

const (
	AppName = "user"
)

func GetService() Service {
	return ioc.Controller().Get(AppName).(Service)
}

type Service interface {
	// CreateUser 创建用户
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	// DeleteUser 删除用户
	DeleteUser(context.Context, *DeleteUserRequest) (*User, error)
	// DescribeUser 查询用户详情
	DescribeUser(context.Context, *DescribeUserRequest) (*User, error)
	// QueryUser 查询用户列表
	QueryUser(context.Context, *QueryUserRequest) (*types.Set[*User], error)
}

type DeleteUserRequest struct {
	Id string `json:"id"`
}

func NewDeleteUserRequest(id string) *DeleteUserRequest {
	return &DeleteUserRequest{
		Id: id,
	}
}

type DescribeUserRequest struct {
	DescribeBy    DescribeBy `query:"describe_by"`
	DescribeValue string     `json:"describe_value"`
}

func NewDescribeUserRequestById(id string) *DescribeUserRequest {
	return &DescribeUserRequest{
		DescribeValue: id,
	}
}

func NewDescribeUserRequestByUserName(username string) *DescribeUserRequest {
	return &DescribeUserRequest{
		DescribeBy:    DescribeByUsername,
		DescribeValue: username,
	}
}

type QueryUserRequest struct {
	*request.PageRequest
	UserIds []uint64 `form:"user" json:"user"`
}

func NewQueryUserRequest() *QueryUserRequest {
	return &QueryUserRequest{
		PageRequest: request.NewDefaultPageRequest(),
		UserIds:     []uint64{},
	}
}

func (r *QueryUserRequest) AddUser(userIds ...uint64) *QueryUserRequest {
	for _, uid := range userIds {
		if !slices.Contains(r.UserIds, uid) {
			r.UserIds = append(r.UserIds, uid)
		}
	}
	return r
}
