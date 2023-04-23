package handler

import (
	"net/http"

	"Interview_Questions/Go-Zero/go-zero-courseware/user/api/internal/logic"
	"Interview_Questions/Go-Zero/go-zero-courseware/user/api/internal/svc"
	"Interview_Questions/Go-Zero/go-zero-courseware/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func userLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserLoginLogic(r.Context(), svcCtx)
		resp, err := l.UserLogin(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
