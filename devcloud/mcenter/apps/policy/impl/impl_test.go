package impl_test

import (
	"context"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/policy"
	"github.com/sword-demon/go18/devcloud/mcenter/test"
)

var (
	impl policy.Service
	ctx  = context.Background()
)

func init() {
	test.DevelopmentSet()
	impl = policy.GetService()
}
