package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/dto/request"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/dto/response"
)

type (
	FindByIdActiveCategoryResponse struct {
		response.FindByIdCategoryResponse
	}
)

func (uc *CategoryController) Find(ctx *gin.Context) {
	req := &request.CategoryFindByIdRequest{}
	res := &FindByIdActiveCategoryResponse{}

	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0001, err.Error()), res))
		return
	}

	acquiredCategory, err := uc.Interactor.FindById(ctx, req.Id)
	if err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0000, err.Error()), res))
		return
	}
	res.Result = response.ActiveCategoryResult{Category: uc.convertActiveCategoryToDTO(acquiredCategory)}
	ctx.JSON(http.StatusOK, res)
}

func (c *FindByIdActiveCategoryResponse) SetErr(err error) {

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
