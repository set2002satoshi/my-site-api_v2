package database

import (
	"github.com/set2002satoshi/my-site-api_v2/models"
	"github.com/set2002satoshi/my-site-api_v2/models/entities"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"gorm.io/gorm"
)

type ActiveUserRepository struct{}

func (r *ActiveUserRepository) Create(db *gorm.DB, user *models.ActiveUserModel) (*models.ActiveUserModel, error) {
	ue, err := r.toEntity(user)
	if err != nil {
		return nil, errors.Wrap(errors.NewCustomError(), errors.REPO0001, err.Error())
	}
	if err := db.Create(ue).Error; err != nil {
		return new(models.ActiveUserModel), errors.Wrap(errors.NewCustomError(), errors.REPO0002, err.Error())
	}
	return r.toModel(ue)
}






func (r *ActiveUserRepository) toModel(obj *entities.TBLUserEntity) (*models.ActiveUserModel, error) {
	return models.NewActiveUserModel(
		obj.UserId,
		obj.Nickname,
		obj.Email,
		string(obj.Password),
		nil,
		obj.IconURL,
		obj.ImageKey,
		true,
		obj.Roll,
		obj.Revision,
		obj.CreatedAt,
		obj.UpdatedAt,
	)
}

func (r *ActiveUserRepository) toEntity(obj *models.ActiveUserModel) (*entities.TBLUserEntity, error) {
	return &entities.TBLUserEntity{
		UserId:    int(obj.GetUserId()),
		Nickname:  obj.GetNickname(),
		Email:     obj.GetEmail(),
		Password:  []byte(obj.GetPassword()),
		IconURL:   string(obj.GetIcon().GetImgURL()),
		ImageKey: string(obj.GetIcon().GetImgKey()),
		Roll:      string(obj.GetRoll()),
		Revision:  int(obj.GetAuditTrail().GetRevision()),
		CreatedAt: obj.GetAuditTrail().GetCreatedAt(),
		UpdatedAt: obj.GetAuditTrail().GetUpdatedAt(),
	}, nil

}
