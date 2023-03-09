package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/dto/request"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/dto/response"
)

type (
	DeleteByIdActiveUserResponse struct {
		response.DeleteActiveUserResponse
	}
)

func (uc *UserController) DeleteById(ctx *gin.Context) {
	req := &request.UserDeleteByIdRequest{}
	res := &DeleteByIdActiveUserResponse{}

	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0001, err.Error()), res))
		return
	}

	err := uc.Interactor.DeleteById(ctx, req.Id)
	if err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0000, err.Error()), res))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *DeleteByIdActiveUserResponse) SetErr(err error) {

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
