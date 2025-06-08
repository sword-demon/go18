// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/v2/http/response"
	"github.com/infraboard/mcube/v2/ioc/config/application"
	"github.com/sword-demon/go18/devcloud/mcenter/apps/token"
	"net/http"
	"net/url"
)

func (h *TokenRestfulApiHandler) Login(r *restful.Request, w *restful.Response) {
	// 获取用户的请求参数 参数在 body 里
	req := token.NewIssueTokenRequest()

	// 获取用户通过 body 传入的参数
	err := r.ReadEntity(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	// 设置当前调用者的 token
	// Private 用户自己的 token
	// 如果是 user/password 这种方式,token 直接放到 body
	switch req.Issuer {
	case token.IssuerPrivateToken:
		req.Parameter.SetAccessToken(token.GetAccessTokenFromHTTP(r.Request))
	}

	// 执行逻辑
	tk, err := h.svc.IssueToken(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	// access_token 通过 SetCookie 直接写到浏览器客户端 web
	http.SetCookie(w, &http.Cookie{
		Name:     token.AccessTokenCookieName,
		Value:    url.QueryEscape(tk.AccessToken),
		Path:     "/",
		Domain:   application.Get().Domain(),
		MaxAge:   0,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
	})
	// 在 header 头中也添加 token
	w.Header().Set(token.AccessTokenResponseHeaderName, tk.AccessToken)

	// body 中返回 token 对象
	response.Success(w, tk)
}

func (h *TokenRestfulApiHandler) Logout(r *restful.Request, w *restful.Response) {
	req := token.NewRevokeTokenRequest(
		token.GetAccessTokenFromHTTP(r.Request),
		token.GetRefreshTokenFromHTTP(r.Request),
	)

	tk, err := h.svc.RevokeToken(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	// access_token 通过SetCookie 直接写到浏览器客户端(Web)
	http.SetCookie(w, &http.Cookie{
		Name:     token.AccessTokenCookieName,
		Value:    "",
		MaxAge:   0,
		Path:     "/",
		Domain:   application.Get().Domain(),
		SameSite: http.SameSiteDefaultMode,
		Secure:   false,
		HttpOnly: true,
	})

	response.Success(w, tk)
}

func (h *TokenRestfulApiHandler) ValidateToken(r *restful.Request, w *restful.Response) {
	// 获取用户的请求参数,参数在 body里
	req := token.NewValidateTokenRequest("")
	err := r.ReadEntity(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	// 执行逻辑
	tk, err := h.svc.ValidateToken(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	// body 中返回 token 对象
	response.Success(w, tk)
}
