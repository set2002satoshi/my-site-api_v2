package database

import (
	"github.com/set2002satoshi/my-site-api_v2/models"
	"github.com/set2002satoshi/my-site-api_v2/models/entities"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/types"
	"gorm.io/gorm"
)

type ActiveCategoryRepository struct{}

func (repo *ActiveCategoryRepository) FindById(db *gorm.DB, id int) (*models.ActiveCategoryModel, error) {
	var acquisitionCategory *entities.TBLCategoryEntity
	if err := db.First(&acquisitionCategory, id).Error; err != nil {
		return &models.ActiveCategoryModel{}, err
	}
	return repo.toModel(acquisitionCategory)
}

func (repo *ActiveCategoryRepository) FindByIds(db *gorm.DB, ids []int) ([]*models.ActiveCategoryModel, error) {
	var categories []*entities.TBLCategoryEntity
	if err := db.Find(&categories, ids).Error; err != nil {
		return []*models.ActiveCategoryModel{}, err
	}
	return repo.toModels(categories)
}

func (repo *ActiveCategoryRepository) FindAll(db *gorm.DB) ([]*models.ActiveCategoryModel, error) {
	var categories []*entities.TBLCategoryEntity
	if err := db.Find(&categories).Error; err != nil {
		return []*models.ActiveCategoryModel{}, err
	}
	return repo.toModels(categories)
}

func (repo *ActiveCategoryRepository) Create(db *gorm.DB, obj *models.ActiveCategoryModel) (*models.ActiveCategoryModel, error) {
	ce, err := repo.toEntity(obj)
	if err != nil {
		return nil, errors.Wrap(errors.NewCustomError(), errors.REPO0001, err.Error())
	}
	if err := db.Create(ce).Error; err != nil {
		return new(models.ActiveCategoryModel), errors.Wrap(errors.NewCustomError(), errors.REPO0011, err.Error())
	}
	return repo.toModel(ce)
}

func (repo *ActiveCategoryRepository) DeleteById(db *gorm.DB, id int) error {
	if err := db.Unscoped().Delete(&entities.TBLCategoryEntity{}, id).Error; err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.REPO0001, err.Error())
	}
	return nil
}

func (repo *ActiveCategoryRepository) toEntity(obj *models.ActiveCategoryModel) (*entities.TBLCategoryEntity, error) {
	return &entities.TBLCategoryEntity{
		CategoryId:   int(obj.GetCategoryId()),
		CategoryName: obj.GetCategoryName(),
		CreatedAt:    obj.GetAuditTrail().GetCreatedAt(),
		UpdatedAt:    obj.GetAuditTrail().GetUpdatedAt(),
	}, nil
}

func (repo *ActiveCategoryRepository) toModel(obj *entities.TBLCategoryEntity) (*models.ActiveCategoryModel, error) {
	return models.NewActiveCategoryModel(
		obj.CategoryId,
		obj.CategoryName,
		types.INITIAL_REVISION,
		obj.CreatedAt,
		obj.UpdatedAt,
	)
}

func (repo *ActiveCategoryRepository) toModels(obj []*entities.TBLCategoryEntity) ([]*models.ActiveCategoryModel, error) {
	models := make([]*models.ActiveCategoryModel, len(obj))
	for i, e := range obj {
		model, err := repo.toModel(e)
		if err != nil {
			return nil, err
		}
		models[i] = model
	}
	return models, nil
}
