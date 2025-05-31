package user

import (
	"context"
	"github.com/infraboard/mcube/v2/types"
)

const (
	AppName = "user"
)

func init() {
}

type Service interface {
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*User, error)
	DescribeUser(context.Context, *DescribeUserRequest) (*User, error)
	QueryUser(context.Context, *QueryUserRequest) (*types.Set[*User], error)
}

type CreateUserRequest struct {
}

type DeleteUserRequest struct {
}

type DescribeUserRequest struct {
}

type QueryUserRequest struct {
}
