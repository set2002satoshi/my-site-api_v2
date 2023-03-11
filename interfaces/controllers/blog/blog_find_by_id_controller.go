package blog

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/dto/request"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/dto/response"
)

type (
	FindByIdActiveBlogResponse struct {
		response.FindByIdActiveBlogResponse
	}
)

func (uc *BlogController) Find(ctx *gin.Context) {
	req := &request.BlogFindByIdRequest{}
	res := &FindByIdActiveBlogResponse{}

	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0001, err.Error()), res))
		return
	}

	acquiredBlog, err := uc.Interactor.FindById(ctx, req.Id)
	if err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0000, err.Error()), res))
		return
	}
	res.Result = response.ActiveBlogResult{Blog: uc.convertActiveBlogToDTO(acquiredBlog)}
	ctx.JSON(http.StatusOK, res)
}

func (c *FindByIdActiveBlogResponse) SetErr(err error) {

	h := make([]errors.ErrorInfo, 0)

	for k, v := range errors.ToMap(err) {
		e := errors.ErrorInfo{
			ErrCode: k,
			ErrMsg:  v,
		}
		h = append(h, e)
	}
	c.Errors = h
}
