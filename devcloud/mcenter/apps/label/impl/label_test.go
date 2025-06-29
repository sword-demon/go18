package impl_test

import (
	"github.com/sword-demon/go18/devcloud/mcenter/apps/label"
	"testing"
)

func TestCreateLabel(t *testing.T) {
	req := label.NewCreateLabelRequest()
	req.Key = "team"
	req.KeyDesc = "小组"
	req.ValueType = label.ValueTypeEnum
	req.AddEnumOption(&label.EnumOption{
		Label: "开发一组",
		Value: "dev01",
		Children: []*label.EnumOption{
			{
				Label: "后端开发",
				Value: "dev01.backend_developer",
			},
			{
				Label: "前端开发",
				Value: "dev01.frontend_developer",
			},
		},
	})

	ins, err := svc.CreateLabel(ctx, req)
	if err != nil {
		t.Fatalf("create label error, %s", err)
	}
	t.Log(ins)
}

func TestQueryLabel(t *testing.T) {
	req := label.NewQueryLabelRequest()
	set, err := svc.QueryLabel(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}
