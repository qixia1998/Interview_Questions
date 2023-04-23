package handler

import (
	"net/http"

	"Interview_Questions/Go-Zero/go-zero-courseware/user/api/internal/logic"
	"Interview_Questions/Go-Zero/go-zero-courseware/user/api/internal/svc"
	"Interview_Questions/Go-Zero/go-zero-courseware/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func userRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserRegisterLogic(r.Context(), svcCtx)
		resp, err := l.UserRegister(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
