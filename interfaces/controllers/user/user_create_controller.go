package user

import (
	"fmt"
	"net/http"
	"time"

	c "github.com/set2002satoshi/my-site-api_v2/interfaces/controllers"
	"github.com/set2002satoshi/my-site-api_v2/models"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/types"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/dto/request"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/dto/response"
)

type (
	CreateActiveUserResponse struct {
		response.CreateActiveUserResponse
	}
)

func (uc *UserController) Create(ctx c.Context) {
	req := &request.UserCreateRequest{}
	res := &CreateActiveUserResponse{}

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0001, err.Error()), res))
		return
	}
	fmt.Println(req)
	// skip Validation

	reqModel, err := uc.cToModel(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0001, err.Error()), res))
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

func (uc *UserController) cToModel(ctx c.Context, req *request.UserCreateRequest) (*models.ActiveUserModel, error) {
	file, err := ctx.FormFile("icon")
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
