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
