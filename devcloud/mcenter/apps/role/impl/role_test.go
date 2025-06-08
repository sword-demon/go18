package impl_test

import (
	"github.com/sword-demon/go18/devcloud/mcenter/apps/role"
	"testing"
)

func TestQueryRole(t *testing.T) {
	req := role.NewQueryRoleRequest()
	set, err := impl.QueryRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestDescribeRole(t *testing.T) {
	req := role.NewDescribeRoleRequest()
	req.SetId(1)
	ins, err := impl.DescribeRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestCreateAdminRole(t *testing.T) {
	req := role.NewCreateRoleRequest()
	req.Name = "admin"
	req.Description = "管理员"
	ins, err := impl.CreateRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
