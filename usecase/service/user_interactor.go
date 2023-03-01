package service

import (
	c "github.com/set2002satoshi/my-site-api_v2/interfaces/controllers"

	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	service "github.com/set2002satoshi/my-site-api_v2/pkg/module/service/aws/s3"

	"github.com/set2002satoshi/my-site-api_v2/models"
	"github.com/set2002satoshi/my-site-api_v2/usecase"
	repo "github.com/set2002satoshi/my-site-api_v2/usecase/repository"
)

type UserInteractor struct {
	DB       usecase.DBRepository
	UserRepo repo.UserRepository
}

func (ui UserInteractor) Register(ctx c.Context, obj *models.ActiveUserModel) (*models.ActiveUserModel, error) {
	db := ui.DB.Connect()
	imgURL, err := service.UploadImage(obj.GetIcon().GetImgFile())
	if err != nil {
		return new(models.ActiveUserModel), err
	}
	um, err := models.NewActiveUserModel(
		int(obj.GetUserId()),
		obj.GetNickname(),
		obj.GetEmail(),
		obj.GetPassword(),
		nil,
		imgURL,
		true,
		string(obj.GetRoll()),
		int(obj.GetAuditTrail().GetRevision()),
		obj.GetAuditTrail().GetCreatedAt(),
		obj.GetAuditTrail().GetUpdatedAt(),
	)
	if err != nil {
		return new(models.ActiveUserModel), errors.NewCustomError()
	}

	created, err := ui.UserRepo.Create(db, um)
	if err != nil {
		txErr := service.DeleteImage(obj.GetIcon().GetImgURL())
		if txErr != nil {
			return new(models.ActiveUserModel), txErr
		}
		return new(models.ActiveUserModel), err
	}
	return created, nil
}