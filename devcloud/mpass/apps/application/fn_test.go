package application

import (
	"github.com/infraboard/mcube/v2/ioc/config/application"
	"reflect"
	"testing"
)

func TestXxx(t *testing.T) {
	app := &application.Application{}
	tt := reflect.TypeOf(app)

	if tt.Kind() == reflect.Ptr {
		tt = tt.Elem()
	}

	fnName := tt.PkgPath() + "." + tt.Name()
	t.Log(fnName)
}
