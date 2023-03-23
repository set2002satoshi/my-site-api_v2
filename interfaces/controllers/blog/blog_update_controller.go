package blog

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
	UpdateActiveBlogResponse struct {
		response.UpdateActiveBlogResponse
	}
)

func (uc *BlogController) Update(ctx *gin.Context) {
	req := &request.BlogUpdateRequest{}
	res := &UpdateActiveBlogResponse{}

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0001, err.Error()), res))
		return
	}
	// skip Validation

	reqModel, err := uc.updateToModel(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0002, err.Error()), res))
		return
	}

	UpdatedBlog, err := uc.Interactor.Update(ctx, reqModel)
	if err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0000, err.Error()), res))
		return
	}

	res.Result = response.ActiveBlogResult{Blog: uc.convertActiveBlogToDTO(UpdatedBlog)}
	ctx.JSON(http.StatusOK, res)
}

func (uc *BlogController) updateToModel(ctx *gin.Context, req *request.BlogUpdateRequest) (*models.ActiveBlogModel, error) {

	// userSId, ok := ctx.Get("userID")
	// if !ok {
	// 	return &models.ActiveBlogModel{}, errors.Add(errors.NewCustomError(), errors.ERR0003)
	// }
	// userId, _ := strconv.Atoi(userSId.(string))

	categoryIds := make([]*models.ActiveCategoryModel, len(req.CategoryIds))
	for i, category := range req.CategoryIds {
		cm, err := models.NewActiveCategoryModel(
			category.Id,
			types.DEFAULT_NAME,
			types.INITIAL_REVISION,
			time.Time{},
			time.Time{},
		)
		if err != nil {
			return &models.ActiveBlogModel{}, errors.Add(errors.NewCustomError(), errors.ERR0004)
		}
		categoryIds[i] = cm
	}

	userId := 1

	return models.NewActiveBlogModel(
		req.BlogId,
		userId,
		types.DEFAULT_NAME,
		req.Title,
		req.Context,
		categoryIds,
		req.Revision,
		time.Time{},
		time.Time{},
	)
}

func (c *UpdateActiveBlogResponse) SetErr(err error) {

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
