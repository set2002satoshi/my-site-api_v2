package user

import (
	"github.com/set2002satoshi/my-site-api_v2/interfaces/database"
	"github.com/set2002satoshi/my-site-api_v2/interfaces/database/config"
	"github.com/set2002satoshi/my-site-api_v2/models"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/dto/response"

	// "github.com/set2002satoshi/my-site-api_v2/models"
	// "github.com/set2002satoshi/my-site-api_v2/pkg/module/dto/response"
	"github.com/set2002satoshi/my-site-api_v2/usecase/service"
)

type UserController struct {
	Interactor service.UserInteractor
}

func NewUserController(db config.DB) *UserController {
	return &UserController{
		Interactor: service.UserInteractor{
			DB:       &config.DBRepository{DB: db},
			UserRepo: &database.ActiveUserRepository{},
		},
	}
}

func (uc *UserController) convertActiveUserToDTO(obj *models.ActiveUserModel) response.ActiveUserEntity {
	return response.ActiveUserEntity{
		UserId:    int(obj.GetUserId()),
		Nickname:  obj.GetNickname(),
		Email:     obj.GetEmail(),
		Password:  obj.GetPassword(),
		IconURL:   obj.GetIcon().GetImgURL(),
		Roll:      string(obj.GetRoll()),
		Revision:  int(obj.GetAuditTrail().GetRevision()),
		CreatedAt: obj.GetAuditTrail().GetCreatedAt(),
		UpdatedAt: obj.GetAuditTrail().GetUpdatedAt(),
	}

}
