package impl_test

import (
	"github.com/sword-demon/go18/devcloud/mcenter/apps/namespace"
	"testing"
)

func TestQueryNamespace(t *testing.T) {
	req := namespace.NewQueryNamespaceRequest()
	set, err := impl.QueryNamespace(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestCreateNamespace(t *testing.T) {
	req := namespace.NewCreateNamespaceRequest()
	req.Name = namespace.DefaultNamespace
	req.Description = "默认空间"
	req.OwnerUserId = 1
	set, err := impl.CreateNamespace(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}
