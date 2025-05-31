package password

import (
	"context"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/token"
	"github.com/sword-demon/go18/devcloud/mcenter/test"
	"testing"
)

func TestPasswordIssuer(t *testing.T) {
	issuer := token.GetIssuer(token.IssuerPassword)
	tk, err := issuer.IssueToken(context.Background(), token.NewIssueParameter().SetUsername("admin").SetPassword("123456"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

func init() {
	test.DevelopmentSet()
}
