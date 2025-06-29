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
	"github.com/sword-demon/go18/devcloud/mpass/apps/application"
	"gorm.io/gorm"
)

func (s *ApplicationServiceImpl) CreateApplication(ctx context.Context, in *application.CreateApplicationRequest) (*application.Application, error) {
	ins, err := application.NewApplication(*in)
	if err != nil {
		return nil, err
	}

	if err := datasource.DBFromCtx(ctx).Save(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *ApplicationServiceImpl) QueryApplication(ctx context.Context, in *application.QueryApplicationRequest) (*types.Set[*application.Application], error) {
	set := types.New[*application.Application]()

	query := in.GormResourceFilter(datasource.DBFromCtx(ctx).Model(&application.Application{}))
	if in.Id != "" {
		query = query.Where("id = ?", in.Id)
	}

	if in.Name != "" {
		query = query.Where("name = ?", in.Name)
	}

	if in.Ready != nil {
		query = query.Where("ready = ?", *in.Ready)
	}

	err := query.Count(&set.Total).Error
	if err != nil {
		return nil, err
	}

	err = query.Order("created_at desc").Offset(int(in.ComputeOffset())).Limit(int(in.PageSize)).
		Find(&set.Items).Error
	if err != nil {
		return nil, err
	}
	return set, nil
}

func (s *ApplicationServiceImpl) UpdateApplication(ctx context.Context, in *application.UpdateApplicationRequest) (*application.Application, error) {
	ins, err := s.DescribeApplication(ctx, &in.DescribeApplicationRequest)
	if err != nil {
		return nil, err
	}

	in.CreateApplicationSpec = in.CreateApplicationSpec

	return ins, datasource.DBFromCtx(ctx).Where("id = ?", in.Id).Updates(ins).Error
}

func (s *ApplicationServiceImpl) DeleteApplication(ctx context.Context, in *application.DeleteApplicationRequest) (*application.Application, error) {
	ins, err := s.DescribeApplication(ctx, &in.DescribeApplicationRequest)
	if err != nil {
		return nil, err
	}

	return ins, datasource.DBFromCtx(ctx).
		Where("id = ?", in.Id).
		Delete(&application.Application{}).
		Error
}

func (s *ApplicationServiceImpl) DescribeApplication(ctx context.Context, in *application.DescribeApplicationRequest) (*application.Application, error) {
	query := in.GormResourceFilter(datasource.DBFromCtx(ctx).Model(&application.Application{}))

	ins := &application.Application{}
	if err := query.Where("id = ?", in.Id).First(ins).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewNotFound("app %s not found", in.Id)
		}
		return nil, err
	}

	return ins, nil
}
