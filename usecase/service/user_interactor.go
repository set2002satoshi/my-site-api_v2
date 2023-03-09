package service

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	service "github.com/set2002satoshi/my-site-api_v2/pkg/module/service/aws/s3"

	"github.com/set2002satoshi/my-site-api_v2/models"
	"github.com/set2002satoshi/my-site-api_v2/usecase"
	repo "github.com/set2002satoshi/my-site-api_v2/usecase/repository"
)

type UserInteractor struct {
	DB          usecase.DBRepository
	UserRepo    repo.UserRepository
	HistoryUserRepo repo.HistoryUserRepository
}

func (ui UserInteractor) FindById(ctx *gin.Context, id int) (*models.ActiveUserModel, error) {
	db := ui.DB.Connect()
	return ui.UserRepo.FindById(db, id)
}

func (ui *UserInteractor) FindAll(ctx *gin.Context) ([]*models.ActiveUserModel, error) {
	db := ui.DB.Connect()
	return ui.UserRepo.FindAll(db)
}

func (ui UserInteractor) Register(ctx *gin.Context, obj *models.ActiveUserModel) (*models.ActiveUserModel, error) {
	db := ui.DB.Connect()
	imgKey, imgURL, err := service.UploadUserImage("user", obj.GetNickname(), obj.GetEmail(), obj.GetIcon().GetImgFile())
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
		imgKey,
		true,
		string(obj.GetRoll()),
		int(obj.GetAuditTrail().GetRevision()),
		obj.GetAuditTrail().GetCreatedAt(),
		obj.GetAuditTrail().GetUpdatedAt(),
	)
	if err != nil {
		err = errors.Combine(err, service.DeleteUserImage(imgKey))
		return new(models.ActiveUserModel), err
	}

	created, err := ui.UserRepo.Create(db, um)
	if err != nil {
		err = errors.Combine(err, service.DeleteUserImage(imgKey))
		return new(models.ActiveUserModel), err
	}

	return created, nil
}

func (ui UserInteractor) Update(ctx *gin.Context, obj *models.ActiveUserModel) (*models.ActiveUserModel, error) {
	tx := ui.DB.Begin()
	currentUser, err := ui.UserRepo.FindById(tx, int(obj.GetUserId()))
	if err != nil {
		tx.Rollback()
		return new(models.ActiveUserModel), err
	}
	if err := currentUser.GetAuditTrail().CountUpRevision(obj.GetAuditTrail().GetRevision()); err != nil {
		tx.Rollback()
		return new(models.ActiveUserModel), err
	}
	imgKey, imgURL, err := service.UploadUserImage("user", obj.GetNickname(), obj.GetEmail(), obj.GetIcon().GetImgFile())
	if err != nil {
		tx.Rollback()
		return new(models.ActiveUserModel), err
	}
	um, err := models.NewActiveUserModel(
		int(currentUser.GetUserId()),
		obj.GetNickname(),
		obj.GetEmail(),
		obj.GetPassword(),
		nil,
		imgURL,
		imgKey,
		true,
		string(obj.GetRoll()),
		int(currentUser.GetAuditTrail().GetRevision()),
		currentUser.GetAuditTrail().GetCreatedAt(),
		time.Now(),
	)
	if err != nil {
		err = errors.Combine(err, service.DeleteUserImage(imgKey))
		tx.Rollback()
		return new(models.ActiveUserModel), err
	}
	updatedUser, err := ui.UserRepo.Update(tx, um)
	if err != nil {
		err = errors.Combine(err, service.DeleteUserImage(imgKey))
		tx.Rollback()
		return new(models.ActiveUserModel), err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return new(models.ActiveUserModel), err
	}
	return updatedUser, nil
}
