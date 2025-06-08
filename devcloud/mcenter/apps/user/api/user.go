package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/v2/http/restful/response"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/user"
	"strconv"
)

// CreateUser 创建用户API
func (h *UserRestfulApiHandler) CreateUser(r *restful.Request, w *restful.Response) {
	req := user.NewCreateUserRequest()

	// 获取用户通过 body 传入的参数
	err := r.ReadEntity(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	// 执行逻辑
	userM, err := h.svc.CreateUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	// body 中返回 user 对象
	response.Success(w, userM)
}

// QueryUser 查询用户 api
// ?user_id=1&user_id=2&user_id=3
func (h *UserRestfulApiHandler) QueryUser(r *restful.Request, w *restful.Response) {
	req := user.NewQueryUserRequest()
	userIdsStr := r.QueryParameters("user_id")
	userIds := make([]uint64, 0, len(userIdsStr))
	for _, idStr := range userIdsStr {
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			response.Failed(w, err)
			return
		}

		userIds = append(userIds, id)
	}
	req.UserIds = userIds
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

	queryUser, err := h.svc.QueryUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, queryUser)
}
