package impl

import (
	"context"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/infraboard/mcube/v2/types"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/namespace"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/policy"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/role"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/user"
	"gorm.io/gorm"
)

func (i *PolicyServiceImpl) CreatePolicy(ctx context.Context, in *policy.CreatePolicyRequest) (*policy.Policy, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	ins := policy.NewPolicy()
	ins.CreatePolicyRequest = *in

	if err := datasource.DBFromCtx(ctx).Create(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *PolicyServiceImpl) QueryPolicy(ctx context.Context, in *policy.QueryPolicyRequest) (*types.Set[*policy.Policy], error) {
	set := types.New[*policy.Policy]()

	query := datasource.DBFromCtx(ctx).Model(&policy.Policy{}).Order("created_at desc")
	if in.UserId != nil {
		query = query.Where("user_id = ?", in.UserId)
	}
	if in.NamespaceId != nil {
		query = query.Where("namespace_id = ?", in.NamespaceId)
	}
	if in.Enabled != nil {
		query = query.Where("enabled = ?", in.Enabled)
	}

	err := query.Count(&set.Total).Error
	if err != nil {
		return nil, err
	}

	if !in.SkipPage {
		query = query.Offset(int(in.ComputeOffset())).Limit(int(in.PageSize))
	}

	if err = query.Find(&set.Items).Error; err != nil {
		return nil, err
	}

	if in.WithUser {
		userReq := user.NewQueryUserRequest()
		set.ForEach(func(t *policy.Policy) {
			userReq.AddUser(t.UserId)
		})
		userSet, err := user.GetService().QueryUser(ctx, userReq)
		if err != nil {
			return nil, err
		}
		set.ForEach(func(p *policy.Policy) {
			p.User = userSet.Filter(func(t *user.User) bool {
				return p.UserId == t.Id
			}).First()
		})
	}

	if in.WithRole {
		roleReq := role.NewQueryRoleRequest()
		set.ForEach(func(t *policy.Policy) {
			roleReq.AddRoleId(t.RoleId)
		})
		roleSet, err := role.GetService().QueryRole(ctx, roleReq)
		if err != nil {
			return nil, err
		}
		set.ForEach(func(p *policy.Policy) {
			p.Role = roleSet.Filter(func(t *role.Role) bool {
				return p.RoleId == t.Id
			}).First()
		})
	}

	if in.WithNamespace {
		nsReq := namespace.NewQueryNamespaceRequest()
		set.ForEach(func(t *policy.Policy) {
			if t.NamespaceId != nil {
				nsReq.AddNamespaceIds(*t.NamespaceId)
			}
		})
		nsSet, err := namespace.GetService().QueryNamespace(ctx, nsReq)
		if err != nil {
			return nil, err
		}
		set.ForEach(func(p *policy.Policy) {
			if p.NamespaceId != nil {
				p.Namespace = nsSet.Filter(func(t *namespace.Namespace) bool {
					return *p.NamespaceId == t.Id
				}).First()
			}
		})
	}

	return set, nil
}

func (i *PolicyServiceImpl) DescribePolicy(ctx context.Context, in *policy.DescribePolicyRequest) (*policy.Policy, error) {
	query := datasource.DBFromCtx(ctx)

	ins := &policy.Policy{}
	if err := query.Where("id = ?", in.Id).First(ins).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewNotFound("policy %d not found", in.Id)
		}
		return nil, err
	}

	return ins, nil
}

func (i *PolicyServiceImpl) UpdatePolicy(ctx context.Context, in *policy.UpdatePolicyRequest) (*policy.Policy, error) {
	descReq := policy.NewDescribePolicyRequest()
	descReq.SetId(in.Id)
	ins, err := i.DescribePolicy(ctx, descReq)
	if err != nil {
		return nil, err
	}

	ins.CreatePolicyRequest = in.CreatePolicyRequest
	return ins, datasource.DBFromCtx(ctx).Where("id = ?", in.Id).Updates(ins).Error
}

func (i *PolicyServiceImpl) DeletePolicy(ctx context.Context, in *policy.DeletePolicyRequest) (*policy.Policy, error) {
	descReq := policy.NewDescribePolicyRequest()
	descReq.SetId(in.Id)
	ins, err := i.DescribePolicy(ctx, descReq)
	if err != nil {
		return nil, err
	}

	return ins, datasource.DBFromCtx(ctx).Where("id = ?", in.Id).Delete(&policy.Policy{}).Error
}
