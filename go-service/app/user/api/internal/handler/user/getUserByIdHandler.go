package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user/api/internal/logic/user"
	"user/api/internal/svc"
	"user/api/internal/types"
)

// 根据id 获取用户
func GetUserByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewGetUserByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetUserById(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
