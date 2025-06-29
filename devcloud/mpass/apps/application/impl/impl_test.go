package impl_test

import (
	"context"
	"github.com/sword-demon/go18/devcloud/mpass/apps/application"
	"github.com/sword-demon/go18/devcloud/mpass/test"
)

var (
	svc application.Service
	ctx = context.Background()
)

func init() {
	test.DevelopmentSetup()
	svc = application.GetService()
}
