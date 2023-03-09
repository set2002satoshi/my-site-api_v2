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
	UpdateActiveUserResponse struct {
		response.UpdateActiveUserResponse
	}
)

func (uc *UserController) Update(ctx *gin.Context) {
	req := &request.UserUpdateRequest{}
	res := &UpdateActiveUserResponse{}

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

	updatedUser, err := uc.Interactor.Update(ctx, reqModel)
	if err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0000, err.Error()), res))
		return
	}

	res.Result = response.ActiveUserResult{User: uc.convertActiveUserToDTO(updatedUser)}
	ctx.JSON(http.StatusOK, res)
}

func (uc *UserController) updateToModel(ctx *gin.Context, req *request.UserUpdateRequest) (*models.ActiveUserModel, error) {
	file, _, err := util.FormFile(ctx, "icon")
	if err != nil {
		return new(models.ActiveUserModel), err
	}

	return models.NewActiveUserModel(
		req.Id,
		req.Name,
		req.Email,
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

func (c *UpdateActiveUserResponse) SetErr(err error) {

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
