package handler

import (
	"net/http"

	"demo2/internal/logic"
	"demo2/internal/svc"
	"demo2/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Demo2Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDemo2Logic(r.Context(), svcCtx)
		resp, err := l.Demo2(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
