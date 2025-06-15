package comptroller

import (
	"context"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/v2/ioc/config/application"
	"github.com/sword-demon/go18/devcloud/audit/apps/event"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/token"
)

// SendEvent 审计日志的发送逻辑
func (sender *EventSender) SendEvent() restful.FilterFunction {
	return func(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
		sr := req.SelectedRoute()
		md := NewMetaData(sr.Metadata())

		// 开启审计日志的开关是否配置
		// 是否在 api 里配置了 Metadata(comptroller.Enable(true))
		if md.GetBool(MetaAuditKey) {
			// 获取当前审计信息
			e := event.NewEvent()

			// 获取用户信息
			tk := token.GetTokenFromCtx(req.Request.Context())
			if tk != nil {
				e.Who = tk.UserName
				e.Namespace = tk.NamespaceName
			}

			// ioc 里面获取当前应用的名称
			e.Service = application.Get().AppName // 每次 init 定义的 AppName
			e.UserAgent = req.Request.UserAgent()
			e.Extras["method"] = sr.Method()
			e.Extras["path"] = sr.Path()
			e.Extras["operation"] = sr.Operation()

			// 补充处理后的数据
			e.StatusCode = resp.StatusCode()
			err := sender.writer.WriteMessages(context.Background(), e.ToKafkaMessage())
			if err != nil {
				sender.log.Error().Msgf("failed to send event, %s", err)
			} else {
				sender.log.Debug().Msgf("send audit event ok, who: %s, resource: %s, action: %s",
					e.Who, e.ResourceType, e.Action)
			}
		}

		// 继续执行后续的过滤器
		chain.ProcessFilter(req, resp)
	}
}

func NewMetaData(data map[string]any) *MetaData {
	return &MetaData{
		data: data,
	}
}

type MetaData struct {
	data map[string]any
}

func (m *MetaData) GetString(key string) string {
	if v, ok := m.data[key]; ok {
		return v.(string)
	}
	return ""
}

func (m *MetaData) GetBool(key string) bool {
	if v, ok := m.data[key]; ok {
		return v.(bool)
	}
	return false
}
