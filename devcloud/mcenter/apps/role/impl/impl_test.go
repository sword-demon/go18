package impl_test

import (
	"context"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/role"
	"github.com/sword-demon/go18/devcloud/mcenter/test"
)

var (
	impl role.Service
	ctx  = context.Background()
)

func init() {
	test.DevelopmentSet()
	impl = role.GetService()
}
