// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package secret

import (
	"context"
	"github.com/infraboard/mcube/v2/http/request"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/types"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/policy"
)

const (
	AppName   = "secret"
	SecretKey = "21312dqwdqwxqweqwdqwdqwd+dwqdqwdqwdwqdqwdqwdqw+="
)

func GetService() Service {
	return ioc.Controller().Get(AppName).(Service)
}

type Service interface {
	CreateSecret(context.Context, *CreateSecretRequest) (*Secret, error)           // CreateSecret 创建 secret 用于 secret 的后台管理
	QuerySecret(context.Context, *QuerySecretRequest) (*types.Set[*Secret], error) // QuerySecret 查询 secret
	DescribeSecret(context.Context, *DescribeSecretRequest) (*Secret, error)       // DescribeSecret 查询详情,api 层需要脱敏
}

type QuerySecretRequest struct {
	policy.ResourceScope
	*request.PageRequest
}

func NewQuerySecretRequest() *QuerySecretRequest {
	return &QuerySecretRequest{
		PageRequest: request.NewDefaultPageRequest(),
	}
}

type DescribeSecretRequest struct {
	policy.ResourceScope
	Id string `json:"id"`
}

func NewDescribeSecretRequest(id string) *DescribeSecretRequest {
	return &DescribeSecretRequest{Id: id}
}
