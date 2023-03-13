package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/dto/request"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/dto/response"
)

type (
	DeleteActiveCategoryResponse struct {
		response.DeleteByIdActiveBlogResponse
	}
)

func (uc *CategoryController) Delete(ctx *gin.Context) {
	req := &request.CategoryDeleteByIdRequest{}
	res := &DeleteActiveCategoryResponse{}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0001, err.Error()), res))
		return
	}

	if err := uc.Interactor.DeleteById(ctx, req.Id);err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0000, err.Error()), res))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *DeleteActiveCategoryResponse) SetErr(err error) {

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
