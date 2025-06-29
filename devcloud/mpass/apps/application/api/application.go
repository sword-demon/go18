// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/gin-gonic/gin/binding"
	"github.com/infraboard/mcube/v2/http/restful/response"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/token"
	"github.com/sword-demon/go18/devcloud/mpass/apps/application"
)

func (h *ApplicationRestfulApiHandler) QueryApplication(r *restful.Request, w *restful.Response) {
	req := application.NewQueryApplicationRequest()
	if err := binding.Query.Bind(r.Request, &req.QueryApplicationRequestSpec); err != nil {
		response.Failed(w, err)
		return
	}

	tk := token.GetTokenFromCtx(r.Request.Context())
	//req.ResourceScope = tk
	log.L().Debug().Msgf("resource scope: %#v", tk)

	set, err := h.svc.QueryApplication(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}
