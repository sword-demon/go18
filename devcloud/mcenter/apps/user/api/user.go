package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/v2/http/restful/response"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/user"
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
