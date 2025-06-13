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
	"github.com/sword-demon/go18/devcloud/mcenter/apps/endpoint"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/namespace"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/policy"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/role"
)

// QueryNamespace 查询用户可访问的空间
func (i *PolicyServiceImpl) QueryNamespace(ctx context.Context, in *policy.QueryNamespaceRequest) (*types.Set[*namespace.Namespace], error) {
	queryNamespaceReq := namespace.NewQueryNamespaceRequest()

	policies, err := i.QueryPolicy(ctx, policy.NewQueryPolicyRequest().SetSkipPage(true).SetUserId(in.UserId).
		SetExpired(false).SetEnabled(true))
	if err != nil {
		return nil, err
	}

	policies.ForEach(func(t *policy.Policy) {
		if t.NamespaceId != nil {
			queryNamespaceReq.AddNamespaceIds(*t.NamespaceId)
		}
	})

	return i.namespace.QueryNamespace(ctx, queryNamespaceReq)
}

func (i *PolicyServiceImpl) QueryMenu(ctx context.Context, in *policy.QueryMenuRequest) (*types.Set[*view.Menu], error) {
	//TODO implement me
	panic("implement me")
}

// QueryEndpoint 查询用户可以访问的 api 接口
// 找到用户可以访问的角色列表,然后再找出角色对于的 api 访问权限
func (i *PolicyServiceImpl) QueryEndpoint(ctx context.Context, in *policy.QueryEndpointRequest) (*types.Set[*endpoint.Endpoint], error) {
	set := types.New[*endpoint.Endpoint]()
	policies, err := i.QueryPolicy(ctx, policy.NewQueryPolicyRequest().
		SetSkipPage(true).
		SetNamespaceId(in.NamespaceId).
		SetUserId(in.UserId).
		SetExpired(false).
		SetEnabled(true))
	if err != nil {
		return nil, err
	}

	roleReq := role.NewQueryMatchedEndpointRequest()
	policies.ForEach(func(t *policy.Policy) {
		roleReq.Add(t.RoleId)
	})

	if policies.Len() > 0 {
		set, err = role.GetService().QueryMatchedEndpoint(ctx, roleReq)
		if err != nil {
			return nil, err
		}
	}

	return set, nil
}

func (i *PolicyServiceImpl) ValidatePagePermission(ctx context.Context, in *policy.ValidatePagePermissionRequest) (*policy.ValidatePagePermissionResponse, error) {
	//TODO implement me
	panic("implement me")
}

// ValidateEndpointPermission 校验 api 接口
func (i *PolicyServiceImpl) ValidateEndpointPermission(ctx context.Context, in *policy.ValidateEndpointPermissionRequest) (*policy.ValidateEndpointPermissionResponse, error) {
	resp := policy.NewValidateEndpointPermissionResponse(*in)

	ns, err := namespace.GetService().DescribeNamespace(ctx, namespace.NewDescribeNamespaceRequest().SetNamespaceId(in.NamespaceId))
	if err != nil {
		return nil, err
	}

	if ns.IsOwner(in.UserId) {
		resp.HasPermission = true
		return resp, nil
	}

	// 非空间拥有者 需要独立鉴权 查询用户可以访问的 api 列表
	endpointReq := policy.NewQueryEndpointRequest()
	endpointReq.UserId = in.UserId
	endpointReq.NamespaceId = in.NamespaceId
	endpointSet, err := i.QueryEndpoint(ctx, endpointReq)
	if err != nil {
		return nil, err
	}

	for _, item := range endpointSet.Items {
		if item.IsMatched(in.Service, in.Method, in.Path) {
			resp.HasPermission = true
			resp.Endpoint = item
			break
		}
	}

	return resp, nil
}
