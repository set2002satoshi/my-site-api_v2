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
	CreateActiveBlogResponse struct {
		response.CreateActiveBlogResponse
	}
)

func (uc *BlogController) Create(ctx *gin.Context) {
	req := &request.BlogCreateRequest{}
	res := &CreateActiveBlogResponse{}

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

	createdBlog, err := uc.Interactor.Register(ctx, reqModel)
	if err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0000, err.Error()), res))
		return
	}

	res.Result = response.ActiveBlogResult{Blog: uc.convertActiveBlogToDTO(createdBlog)}
	ctx.JSON(http.StatusOK, res)
}

func (uc *BlogController) createToModel(ctx *gin.Context, req *request.BlogCreateRequest) (*models.ActiveBlogModel, error) {

	// userSId, ok := ctx.Get("userID")
	// if !ok {
	// 	return &models.ActiveBlogModel{}, errors.Add(errors.NewCustomError(), errors.ERR0003)
	// }
	// userId, _ := strconv.Atoi(userSId.(string))

	userId := 1

	return models.NewActiveBlogModel(
		types.INITIAL_ID,
		userId,
		types.DEFAULT_NAME,
		req.Title,
		req.Context,
		types.INITIAL_REVISION,
		time.Time{},
		time.Time{},
	)
}

func (c *CreateActiveBlogResponse) SetErr(err error) {

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
