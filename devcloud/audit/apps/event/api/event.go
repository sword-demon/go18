package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/v2/http/restful/response"
	"github.com/sword-demon/go18/devcloud/audit/apps/event"
	"strconv"
)

func (h *EventRestfulApiHandler) QueryEventApi(r *restful.Request, w *restful.Response) {
	req := event.NewQueryEventRequest()

	pageSizeStr := r.QueryParameter("page_size")
	pageSize, err := strconv.ParseUint(pageSizeStr, 10, 64)
	if err != nil {
		response.Failed(w, err)
		return
	}
	req.PageSize = pageSize

	pageNumberStr := r.QueryParameter("page_number")
	pageNumber, err := strconv.ParseUint(pageNumberStr, 10, 64)
	if err != nil {
		response.Failed(w, err)
		return
	}
	req.PageNumber = pageNumber

	// 执行逻辑
	userM, err := h.event.QueryEvent(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, userM)
}
