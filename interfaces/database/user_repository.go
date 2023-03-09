package database

import (
	"github.com/set2002satoshi/my-site-api_v2/models"
	"github.com/set2002satoshi/my-site-api_v2/models/entities"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"gorm.io/gorm"
)

type ActiveUserRepository struct{}

func (r *ActiveUserRepository) FindById(db *gorm.DB, id int) (*models.ActiveUserModel, error) {
	var acquisitionUser *entities.TBLUserEntity
	if err := db.Where("user_id = ?", id).First(&acquisitionUser).Error; err != nil {
		return &models.ActiveUserModel{}, errors.Wrap(errors.NewCustomError(), errors.REPO0003, gorm.ErrRecordNotFound.Error())
	}
	return r.toModel(acquisitionUser)
}

func (r *ActiveUserRepository) FindAll(db *gorm.DB) ([]*models.ActiveUserModel, error) {
	var acquisitionUsers []*entities.TBLUserEntity
	if err := db.Find(&acquisitionUsers).Error; err != nil {
		return []*models.ActiveUserModel{}, err
	}
	return r.toModels(acquisitionUsers)
}

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

func (r *ActiveUserRepository) Update(tx *gorm.DB, obj *models.ActiveUserModel) (*models.ActiveUserModel, error) {
	ue, err := r.toEntity(obj)
	if err != nil {
		return nil, errors.Wrap(errors.NewCustomError(), errors.REPO0001, err.Error())
	}
	if err := tx.Select("email", "nickname", "password", "roll", "revision", "image_key", "icon_url").Updates(&ue).Error; err != nil {
		return nil, errors.Wrap(errors.NewCustomError(), errors.REPO0004, err.Error())
	}
	return r.toModel(ue)
}

func (r *ActiveUserRepository) DeleteById(tx *gorm.DB, id int) error {
	if err := tx.Unscoped().Delete(&entities.TBLUserEntity{}, id).Error; err != nil {
		return err
	}
	return nil
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
func (r *ActiveUserRepository) toModels(en []*entities.TBLUserEntity) ([]*models.ActiveUserModel, error) {
	models := make([]*models.ActiveUserModel, len(en))
	for i, e := range en {
		model, err := r.toModel(e)
		if err != nil {
			return nil, err
		}
		models[i] = model
	}
	return models, nil
}

func (r *ActiveUserRepository) toEntity(obj *models.ActiveUserModel) (*entities.TBLUserEntity, error) {
	return &entities.TBLUserEntity{
		UserId:    int(obj.GetUserId()),
		Nickname:  obj.GetNickname(),
		Email:     obj.GetEmail(),
		Password:  []byte(obj.GetPassword()),
		IconURL:   string(obj.GetIcon().GetImgURL()),
		ImageKey:  string(obj.GetIcon().GetImgKey()),
		Roll:      string(obj.GetRoll()),
		Revision:  int(obj.GetAuditTrail().GetRevision()),
		CreatedAt: obj.GetAuditTrail().GetCreatedAt(),
		UpdatedAt: obj.GetAuditTrail().GetUpdatedAt(),
	}, nil
}

func (r *ActiveUserRepository) toEntities(obj []*models.ActiveUserModel) ([]*entities.TBLUserEntity, error) {
	UEs := make([]*entities.TBLUserEntity, len(obj))
	for i, v := range obj {
		obj, err := r.toEntity(v)
		if err != nil {
			return nil, err
		}
		UEs[i] = obj
	}
	return UEs, nil
}
