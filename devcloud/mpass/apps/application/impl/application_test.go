package impl_test

import (
	"github.com/sword-demon/go18/devcloud/mpass/apps/application"
	"testing"
)

func TestCreateApplication(t *testing.T) {
	req := application.NewCreateApplicationRequest()
	req.Name = "devcloud"
	req.Description = "应用研发云"
	req.Type = application.TypeSourceCode
	req.CodeRepository = application.CodeRepository{
		SshUrl: "git@github.com:sword-demon/go18.git",
	}
	req.SetLabel("team", "dev01.web_developer")

	ins, err := svc.CreateApplication(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ins)
}

func TestQueryApplication(t *testing.T) {
	req := application.NewQueryApplicationRequest()
	req.SetScope("team", []string{"%"})

	ins, err := svc.QueryApplication(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
