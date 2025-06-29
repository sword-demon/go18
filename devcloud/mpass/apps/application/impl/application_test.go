// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

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
	req.SetNamespaceId(1)
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
