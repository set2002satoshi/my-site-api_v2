package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/dto/response"
)

type (
	FindAllActiveUserResponse struct {
		response.FindAllActiveUserResponse
	}
)

func (uc *UserController) FindAll(ctx *gin.Context) {

	res := &FindAllActiveUserResponse{}

	acquiredUser, err := uc.Interactor.FindAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0000, err.Error()), res))
		return
	}

	res.Results = response.ActiveUserResults{Users: uc.convertActiveUserToDTOs(acquiredUser)}
	ctx.JSON(http.StatusOK, res)
}

func (c *FindAllActiveUserResponse) SetErr(err error) {

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
