package impl

import (
	"context"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/token"
	"github.com/sword-demon/go18/devcloud/mcenter/test"
)

var (
	svc token.Service
	ctx = context.Background()
)

func init() {
	test.DevelopmentSet()
	svc = token.GetService()
}
