package user

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/set2002satoshi/my-site-api_v2/models"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/types"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/dto/request"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/dto/response"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/service/util"
)

type (
	CreateActiveUserResponse struct {
		response.CreateActiveUserResponse
	}
)

func (uc *UserController) Create(ctx *gin.Context) {
	req := &request.UserCreateRequest{}
	res := &CreateActiveUserResponse{}

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

	createdUser, err := uc.Interactor.Register(ctx, reqModel)
	if err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0000, err.Error()), res))
		return
	}

	res.Result = response.ActiveUserResult{User: uc.convertActiveUserToDTO(createdUser)}
	ctx.JSON(http.StatusOK, res)
}

func (uc *UserController) createToModel(ctx *gin.Context, req *request.UserCreateRequest) (*models.ActiveUserModel, error) {
	file, _, err := util.FormFile(ctx, "icon")
	if err != nil {
		return new(models.ActiveUserModel), err
	}

	return models.NewActiveUserModel(
		types.INITIAL_ID,
		req.Email,
		req.Name,
		req.Password,
		file,
		types.DEFAULT_URL,
		types.DEFAULT_KEY,
		false,
		req.Roll,
		types.INITIAL_REVISION,
		time.Time{},
		time.Time{},
	)
}

func (c *CreateActiveUserResponse) SetErr(err error) {

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
