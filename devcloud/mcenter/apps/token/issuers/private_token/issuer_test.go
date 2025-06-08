// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package private_token_test

import (
	"context"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/token"
	"github.com/sword-demon/go18/devcloud/mcenter/test"
	"testing"
)

func TestPasswordIssuer(t *testing.T) {
	issuer := token.GetIssuer(token.IssuerPrivateToken)
	tk, err := issuer.IssueToken(context.Background(), token.NewIssueParameter().SetAccessToken("LccvuTwISJRheu8PtqAFTJBy").SetExpireTTL(24*60*60))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

func init() {
	test.DevelopmentSet()
}
