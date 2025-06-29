package impl

import (
	"context"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/infraboard/mcube/v2/types"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/label"
	"gorm.io/gorm"
)

func (s *LabelServiceImpl) CreateLabel(ctx context.Context, in *label.CreateLabelRequest) (*label.Label, error) {
	ins, err := label.NewLabel(in)
	if err != nil {
		return nil, err
	}

	if err := datasource.DBFromCtx(ctx).Create(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *LabelServiceImpl) UpdateLabel(ctx context.Context, in *label.UpdateLabelRequest) (*label.Label, error) {
	descReq := label.NewDescribeLabelRequest()
	descReq.SetId(in.Id)
	ins, err := s.DescribeLabel(ctx, descReq)
	if err != nil {
		return nil, err
	}

	ins.CreateLabelSpec = *in.Spec
	return ins, datasource.DBFromCtx(ctx).Where("id = ?", in.Id).Updates(ins).Error
}

func (s *LabelServiceImpl) DeleteLabel(ctx context.Context, in *label.DeleteLabelRequest) (*label.Label, error) {
	descReq := label.NewDescribeLabelRequest()
	descReq.SetId(in.Id)
	ins, err := s.DescribeLabel(ctx, descReq)
	if err != nil {
		return nil, err
	}

	return ins, datasource.DBFromCtx(ctx).Delete(ins).Error
}

func (s *LabelServiceImpl) DescribeLabel(ctx context.Context, in *label.DescribeLabelRequest) (*label.Label, error) {
	query := datasource.DBFromCtx(ctx).Model(&label.Label{})

	ins := &label.Label{}
	if err := query.Where("id = ?", in.Id).First(ins).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewNotFound("label %s not found", in.Id)
		}
		return nil, err
	}

	return ins, nil
}

func (s *LabelServiceImpl) QueryLabel(ctx context.Context, in *label.QueryLabelRequest) (*types.Set[*label.Label], error) {
	set := types.NewSet[*label.Label]()

	query := datasource.DBFromCtx(ctx).Model(&label.Label{})
	err := query.Count(&set.Total).Error
	if err != nil {
		return nil, err
	}

	err = query.Order("created_at desc").Offset(int(in.ComputeOffset())).Limit(int(in.PageSize)).Find(&set.Items).Error
	if err != nil {
		return nil, err
	}

	return set, nil
}
