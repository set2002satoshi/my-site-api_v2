package database

import (
	"github.com/set2002satoshi/my-site-api_v2/models"
	"github.com/set2002satoshi/my-site-api_v2/models/entities"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/types"
	"gorm.io/gorm"
)

type CategoryRepository struct{}

func (repo *CategoryRepository) FindAll(db *gorm.DB) ([]*models.ActiveCategoryModel, error) {
	var category []*entities.TBLCategoryEntity
	if err := db.Find(&category).Error; err != nil {
		return []*models.ActiveCategoryModel{}, err
	}
	return repo.toModels(category)
}

func (repo *CategoryRepository) Create(db *gorm.DB, obj *models.ActiveCategoryModel) (*models.ActiveCategoryModel, error) {
	ce, err := repo.toEntity(obj)
	if err != nil {
		return nil, errors.Wrap(errors.NewCustomError(), errors.REPO0001, err.Error())
	}
	if err := db.Create(ce).Error; err != nil {
		return new(models.ActiveCategoryModel), errors.Wrap(errors.NewCustomError(), errors.REPO0011, err.Error())
	}
	return repo.toModel(ce)
}

func (repo *CategoryRepository) toEntity(obj *models.ActiveCategoryModel) (*entities.TBLCategoryEntity, error) {
	return &entities.TBLCategoryEntity{
		CategoryId:   int(obj.GetCategoryId()),
		CategoryName: obj.GetCategoryName(),
		CreatedAt:    obj.GetAuditTrail().GetCreatedAt(),
		UpdatedAt:    obj.GetAuditTrail().GetUpdatedAt(),
	}, nil
}

func (repo *CategoryRepository) toModel(obj *entities.TBLCategoryEntity) (*models.ActiveCategoryModel, error) {
	return models.NewActiveCategoryModel(
		obj.CategoryId,
		obj.CategoryName,
		types.INITIAL_REVISION,
		obj.CreatedAt,
		obj.UpdatedAt,
	)
}

func (repo *CategoryRepository) toModels(obj []*entities.TBLCategoryEntity) ([]*models.ActiveCategoryModel, error) {
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