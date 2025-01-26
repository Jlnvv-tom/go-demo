package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user/api/internal/logic/user"
	"user/api/internal/svc"
	"user/api/internal/types"
)

// 根据所有用户
func GetAllUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetListUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewGetAllUserLogic(r.Context(), svcCtx)
		resp, err := l.GetAllUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
