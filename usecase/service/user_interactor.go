package service

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/types"
	service "github.com/set2002satoshi/my-site-api_v2/pkg/module/service/aws/s3"

	"github.com/set2002satoshi/my-site-api_v2/models"
	"github.com/set2002satoshi/my-site-api_v2/usecase"
	repo "github.com/set2002satoshi/my-site-api_v2/usecase/repository"
)

type UserInteractor struct {
	DB              usecase.DBRepository
	UserRepo        repo.UserRepository
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
	historyUser, err := ui.ActiveToHistory(currentUser)
	if err != nil {
		return new(models.ActiveUserModel), errors.Add(errors.NewCustomError(), errors.SE0005)
	}

	_, err = ui.HistoryUserRepo.Create(tx, historyUser)
	if err != nil {
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

func (ui *UserInteractor) DeleteById(ctx *gin.Context, id int) error {
	tx := ui.DB.Begin()
	currentUser, err := ui.UserRepo.FindById(tx, id)
	if err != nil {
		return err
	}
	historyUser, err := models.NewHistoryUserModel(
		types.INITIAL_ID,
		int(currentUser.GetUserId()),
		currentUser.GetNickname(),
		currentUser.GetEmail(),
		currentUser.GetPassword(),
		currentUser.GetIcon().GetImgKey(),
		currentUser.GetIcon().GetImgURL(),
		string(currentUser.GetRoll()),
		int(currentUser.GetAuditTrail().GetRevision()),
		currentUser.GetAuditTrail().GetCreatedAt(),
		time.Time{},
	)
	if err != nil {
		return errors.Add(errors.NewCustomError(), errors.SE0004)
	}
	_, err = ui.HistoryUserRepo.Create(tx, historyUser)
	if err != nil {
		tx.Rollback()
		return err
	}
	if err := ui.UserRepo.DeleteById(tx, id); err != nil {
		tx.Rollback()
		return err
	}
	if err := service.DeleteUserImage(currentUser.GetIcon().GetImgKey()); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (ui *UserInteractor) ActiveToHistory(obj *models.ActiveUserModel) (*models.HistoryUserModel, error) {
	return models.NewHistoryUserModel(
		types.INITIAL_ID,
		int(obj.GetUserId()),
		obj.GetNickname(),
		obj.GetEmail(),
		obj.GetPassword(),
		obj.GetIcon().GetImgURL(),
		obj.GetIcon().GetImgKey(),
		string(obj.GetRoll()),
		int(obj.GetAuditTrail().GetRevision()),
		obj.GetAuditTrail().GetCreatedAt(),
		obj.GetAuditTrail().GetUpdatedAt(),
	)
}
