package impl_test

import (
	"context"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/label"
	"github.com/sword-demon/go18/devcloud/mcenter/test"
)

var (
	svc label.Service
	ctx = context.Background()
)

func init() {
	test.DevelopmentSet()
	svc = label.GetService()
}
