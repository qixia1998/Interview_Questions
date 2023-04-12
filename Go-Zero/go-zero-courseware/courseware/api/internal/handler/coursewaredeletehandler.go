package handler

import (
	"net/http"

	"Interview_Questions/Go-Zero/go-zero-courseware/courseware/api/internal/logic"
	"Interview_Questions/Go-Zero/go-zero-courseware/courseware/api/internal/svc"
	"Interview_Questions/Go-Zero/go-zero-courseware/courseware/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func coursewareDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCoursewareDeleteLogic(r.Context(), svcCtx)
		resp, err := l.CoursewareDelete(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
