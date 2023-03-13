package category

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/set2002satoshi/my-site-api_v2/models"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/types"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/dto/request"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/dto/response"
)

type (
	CreateActiveCategoryResponse struct {
		response.CreateCategoryResponse
	}
)

func (uc *CategoryController) Create(ctx *gin.Context) {
	req := &request.CategoryCreateRequest{}
	res := &CreateActiveCategoryResponse{}

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0001, err.Error()), res))
		return
	}
	// skip Validation

	reqModel, err := uc.createToModel(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0002, err.Error()), res))
		return
	}

	createdCategory, err := uc.Interactor.Register(ctx, reqModel)
	if err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0000, err.Error()), res))
		return
	}

	res.Result = response.ActiveCategoryResult{Category: uc.convertActiveCategoryToDTO(createdCategory)}
	ctx.JSON(http.StatusOK, res)
}

func (uc *CategoryController) createToModel(ctx *gin.Context, req *request.CategoryCreateRequest) (*models.ActiveCategoryModel, error) {

	return models.NewActiveCategoryModel(
		types.INITIAL_ID,
		req.CategoryName,
		types.INITIAL_REVISION,
		time.Time{},
		time.Time{},
	)
}

func (c *CreateActiveCategoryResponse) SetErr(err error) {

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
