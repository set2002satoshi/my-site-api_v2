package database

import (
	"github.com/set2002satoshi/my-site-api_v2/models"
	"github.com/set2002satoshi/my-site-api_v2/models/entities"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"gorm.io/gorm"
)

type HistoryBlogWithCategoryRepository struct{}

func (repo HistoryBlogWithCategoryRepository) Create(db *gorm.DB, obj *models.HistoryBlogWithCategoryModel) (*models.HistoryBlogWithCategoryModel, error) {
	es, err := repo.toEntity(obj)
	if err != nil {
		return new(models.HistoryBlogWithCategoryModel), errors.Wrap(errors.NewCustomError(), errors.REPO0001, err.Error())
	}
	if err := db.Create(&es).Error; err != nil {
		return new(models.HistoryBlogWithCategoryModel), errors.Wrap(errors.NewCustomError(), errors.REPO0016, err.Error())
	}
	return repo.toModel(es)
}

func (repo *HistoryBlogWithCategoryRepository) toEntity(obj *models.HistoryBlogWithCategoryModel) (*entities.HistoryBlogWithCategoriesEntity, error) {
	return &entities.HistoryBlogWithCategoriesEntity{
		ActiveId: obj.GetActiveId(),
		CategoryId: int(obj.GetCategoryId()),
		BlogId:     int(obj.GetBlogId()),
	}, nil
}

func (repo *HistoryBlogWithCategoryRepository) toEntities(objs []*models.HistoryBlogWithCategoryModel) ([]*entities.HistoryBlogWithCategoriesEntity, error) {
	entities := make([]*entities.HistoryBlogWithCategoriesEntity, len(objs))
	for i, obj := range objs {
		entity, err := repo.toEntity(obj)
		if err != nil {
			return nil, err
		}
		entities[i] = entity
	}
	return entities, nil
}

func (repo *HistoryBlogWithCategoryRepository) toModel(obj *entities.HistoryBlogWithCategoriesEntity) (*models.HistoryBlogWithCategoryModel, error) {
	return models.NewHistoryBlogWithCategoryModel(
		obj.Id,
		obj.ActiveId,
		obj.CategoryId,
		obj.BlogId,
	)
}

func (repo *HistoryBlogWithCategoryRepository) toModels(objs []*entities.HistoryBlogWithCategoriesEntity) ([]*models.HistoryBlogWithCategoryModel, error) {
	models := make([]*models.HistoryBlogWithCategoryModel, len(objs))
	for i, v := range objs {
		model, err := repo.toModel(v)
		if err != nil {
			return nil, err
		}
		models[i] = model
	}
	return models, nil
}
