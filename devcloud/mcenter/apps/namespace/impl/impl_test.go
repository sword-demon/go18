package impl_test

import (
	"context"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/namespace"
	"github.com/sword-demon/go18/devcloud/mcenter/test"
)

var (
	impl namespace.Service
	ctx  = context.Background()
)

func init() {
	test.DevelopmentSet()
	impl = namespace.GetService()
}
