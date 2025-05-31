package impl

import (
	"github.com/sword-demon/go18/devcloud/mcenter/apps/token"
	"github.com/sword-demon/go18/devcloud/mcenter/test"
	"testing"
)

func TestBearer(t *testing.T) {
	test.DevelopmentSet()
	t.Log(token.MakeBearer(24))
	t.Log(token.MakeBearer(24))
}

func TestIssueToken(t *testing.T) {
	req := token.NewIssueTokenRequest()
	req.IssueByPassword("admin", "123456")
	req.Source = token.SourceWeb
	set, err := svc.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestQueryToken(t *testing.T) {
	req := token.NewQueryTokenRequest()
	set, err := svc.QueryToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}
