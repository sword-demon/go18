package impl_test

import (
	"github.com/sword-demon/go18/devcloud/mcenter/apps/policy"
	"testing"
)

func TestQueryNamespace(t *testing.T) {
	req := policy.NewQueryNamespaceRequest()
	req.UserId = 1
	set, err := impl.QueryNamespace(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestQueryEndpoint(t *testing.T) {
	req := policy.NewQueryEndpointRequest()
	req.UserId = 2
	req.NamespaceId = 1
	set, err := impl.QueryEndpoint(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestValidateEndpointPermission(t *testing.T) {
	req := policy.NewValidateEndpointPermissionRequest()
	req.UserId = 1
	req.NamespaceId = 1
	req.Service = "devcloud"
	req.Method = "GET"
	req.Path = "/api/devcloud/v1/users/"
	set, err := impl.ValidateEndpointPermission(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}
