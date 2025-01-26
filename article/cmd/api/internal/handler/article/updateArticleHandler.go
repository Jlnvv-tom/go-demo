package article

import (
	"net/http"

	"article/cmd/api/internal/logic/article"
	"article/cmd/api/internal/svc"
	"article/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 修改文章
func UpdateArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateArticleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := article.NewUpdateArticleLogic(r.Context(), svcCtx)
		resp, err := l.UpdateArticle(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
