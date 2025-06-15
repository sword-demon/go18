package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/v2/http/restful/response"
	"github.com/sword-demon/go18/devcloud/audit/apps/event"
)

func (h *EventRestfulApiHandler) QueryEventApi(r *restful.Request, w *restful.Response) {
	req := event.NewQueryEventRequest()

	// 执行逻辑
	userM, err := h.event.QueryEvent(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, userM)
}
