package impl

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/label"
)

func init() {
	ioc.Controller().Registry(&LabelServiceImpl{})
}

var _ label.Service = (*LabelServiceImpl)(nil)

type LabelServiceImpl struct {
	ioc.ObjectImpl
}

func (s *LabelServiceImpl) Name() string {
	return label.AppName
}

func (s *LabelServiceImpl) Init() error {
	if datasource.Get().AutoMigrate {
		err := datasource.DB().AutoMigrate(&label.Label{})
		if err != nil {
			return err
		}
	}
	return nil
}
