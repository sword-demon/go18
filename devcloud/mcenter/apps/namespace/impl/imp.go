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
	"github.com/sword-demon/go18/devcloud/mcenter/apps/namespace"
	"gorm.io/gorm"
)

func (i *NamespaceServiceImpl) CreateNamespace(ctx context.Context, in *namespace.CreateNamespaceRequest) (*namespace.Namespace, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	ins := namespace.NewNamespace()
	ins.CreateNamespaceRequest = *in

	if err := datasource.DBFromCtx(ctx).Create(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *NamespaceServiceImpl) QueryNamespace(ctx context.Context, in *namespace.QueryNamespaceRequest) (*types.Set[*namespace.Namespace], error) {
	set := types.New[*namespace.Namespace]()

	query := datasource.DBFromCtx(ctx).Model(&namespace.Namespace{})
	err := query.Count(&set.Total).Error
	if err != nil {
		return nil, err
	}

	err = query.Order("created_at desc").
		Offset(int(in.ComputeOffset())).
		Limit(int(in.PageSize)).
		Find(&set.Items).Error
	if err != nil {
		return nil, err
	}

	return set, nil
}

func (i *NamespaceServiceImpl) DescribeNamespace(ctx context.Context, in *namespace.DescribeNamespaceRequest) (*namespace.Namespace, error) {
	query := datasource.DBFromCtx(ctx)

	ins := namespace.NewNamespace()
	if err := query.Where("id = ?", in.Id).First(ins).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewNotFound("namespace %d not found", in.Id)
		}
		return nil, err
	}

	return ins, nil
}

func (i *NamespaceServiceImpl) UpdateNamespace(ctx context.Context, in *namespace.UpdateNamespaceRequest) (*namespace.Namespace, error) {
	desReq := namespace.NewDescribeNamespaceRequest()
	desReq.SetId(in.Id)

	ins, err := i.DescribeNamespace(ctx, desReq)
	if err != nil {
		return nil, err
	}

	ins.CreateNamespaceRequest = in.CreateNamespaceRequest
	return ins, datasource.DBFromCtx(ctx).Where("id = ?", in.Id).Updates(ins).Error
}

func (i *NamespaceServiceImpl) DeleteNamespace(ctx context.Context, in *namespace.DeleteNamespaceRequest) (*namespace.Namespace, error) {
	desReq := namespace.NewDescribeNamespaceRequest()
	desReq.SetId(in.Id)

	ins, err := i.DescribeNamespace(ctx, desReq)
	if err != nil {
		return nil, err
	}

	return ins, datasource.DBFromCtx(ctx).Where("id = ?", in.Id).Delete(&namespace.Namespace{}).Error
}
