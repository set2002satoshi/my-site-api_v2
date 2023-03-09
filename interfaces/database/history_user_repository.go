package database

import (
	"time"

	"github.com/set2002satoshi/my-site-api_v2/models"
	"github.com/set2002satoshi/my-site-api_v2/models/entities"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"gorm.io/gorm"
)

type HistoryUserRepository struct{}

func (hur *HistoryUserRepository) FindById(db *gorm.DB, id int) (*models.HistoryUserModel, error) {
	var acquisitionUser *entities.HistoryUserEntity
	if err := db.Where("user_id = ?", id).First(&acquisitionUser); err != nil {
		return &models.HistoryUserModel{}, errors.Wrap(errors.NewCustomError(), errors.REPO0005, gorm.ErrRecordNotFound.Error())
	}
	return hur.toModel(acquisitionUser)
}

func (hur *HistoryUserRepository) Create(db *gorm.DB, user *models.HistoryUserModel) (*models.HistoryUserModel, error) {
	ue, err := hur.toEntity(user)
	if err != nil {
		return nil, errors.Wrap(errors.NewCustomError(), errors.REPO0001, err.Error())
	}
	if err := db.Create(ue).Error; err != nil {
		return new(models.HistoryUserModel), errors.Wrap(errors.NewCustomError(), errors.REPO0002, err.Error())
	}
	return hur.toModel(ue)
}

func (hur *HistoryUserRepository) DeleteById(db *gorm.DB, id int) error {
	if err := db.Unscoped().Delete(&entities.HistoryUserEntity{}, id).Error; err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.REPO0007, err.Error())
	}
	return nil
}

func (hur *HistoryUserRepository) toModel(obj *entities.HistoryUserEntity) (*models.HistoryUserModel, error) {
	return models.NewHistoryUserModel(
		obj.Id,
		obj.ActiveUserId,
		obj.Nick,
		obj.Email,
		string(obj.Password),
		obj.ImgURL,
		obj.ImgKey,
		obj.Roll,
		obj.Revision,
		obj.CreatedAt,
		time.Time{},
	)
}

func (hur *HistoryUserRepository) toEntity(obj *models.HistoryUserModel) (*entities.HistoryUserEntity, error) {
	return &entities.HistoryUserEntity{
		Id:           int(obj.GetId()),
		ActiveUserId: int(obj.GetActiveUserId()),
		Nick:         obj.GetNickname(),
		Email:        obj.GetEmail(),
		Password:     []byte(obj.GetPassword()),
		ImgURL:       obj.GetIcon().GetImgURL(),
		ImgKey:       obj.GetIcon().GetImgKey(),
		Roll:         string(obj.GetRoll()),
		Revision:     int(obj.GetAuditTrail().GetRevision()),
		CreatedAt:    obj.GetAuditTrail().GetCreatedAt(),
		DeletedAt:    time.Now().UTC(),
	}, nil
}
