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
	"gorm.io/gorm"
)

// RegistryEndpoint 注册 api 接口
// 批量注册 一次添加多个记录
// 需要保证事务 同时成功,或者同时失败
func (s *EndpointServiceImpl) RegistryEndpoint(ctx context.Context, in *endpoint.RegistryEndpointRequest) (*types.Set[*endpoint.Endpoint], error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	set := types.New[*endpoint.Endpoint]()
	err := datasource.DBFromCtx(ctx).Transaction(func(tx *gorm.DB) error {
		for i := range in.Items {
			item := in.Items[i].BuildUUID()
			ins := endpoint.NewEndpoint().SetRouteEntry(*item)

			oldEndpoint := endpoint.NewEndpoint()
			if err := tx.Where("uuid = ?", item.UUID).Take(oldEndpoint).Error; err != nil {
				if err != gorm.ErrRecordNotFound {
					return err
				}

				// 需要创建
				if err := tx.Save(ins).Error; err != nil {
					return err
				}
			} else {
				// 需要更新
				ins.Id = oldEndpoint.Id
				if err := tx.Where("uuid = ?", item.UUID).Updates(ins).Error; err != nil {
					return err
				}
			}
			set.Add(ins)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return set, nil
}

func (s *EndpointServiceImpl) QueryEndpoint(ctx context.Context, in *endpoint.QueryEndpointRequest) (*types.Set[*endpoint.Endpoint], error) {
	set := types.New[*endpoint.Endpoint]()

	query := datasource.DBFromCtx(ctx).Model(&endpoint.Endpoint{})
	if len(in.Services) > 0 && !in.IsMatchAllService() {
		query = query.Where("service in ?", in.Services)
	}

	err := query.Count(&set.Total).Error
	if err != nil {
		return nil, err
	}

	err = query.Order("created_at desc").Find(&set.Items).Error
	if err != nil {
		return nil, err
	}

	return set, nil
}

func (s *EndpointServiceImpl) DescribeEndpoint(ctx context.Context, in *endpoint.DescribeEndpointRequest) (*endpoint.Endpoint, error) {
	query := datasource.DBFromCtx(ctx)

	ins := endpoint.NewEndpoint()
	if err := query.Where("id = ?", in.Id).First(ins).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewNotFound("endpoint %s not found", in)
		}
		return nil, err
	}

	return ins, nil
}
