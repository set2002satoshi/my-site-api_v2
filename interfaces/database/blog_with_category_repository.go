package database

import (
	"fmt"

	"github.com/set2002satoshi/my-site-api_v2/models"
	"github.com/set2002satoshi/my-site-api_v2/models/entities"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"gorm.io/gorm"
)

type BlogWithCategoryRepository struct{}

func (repo *BlogWithCategoryRepository) FindById(db *gorm.DB, id int) (*models.ActiveBlogWithCategoryModel, error) {
	var acquisitionIds *entities.TBLBlogWithCategoriesEntity
	if err := db.First(&acquisitionIds, id).Error; err != nil {
		return new(models.ActiveBlogWithCategoryModel), err
	}
	return repo.toModel(acquisitionIds)
}

func (repo *BlogWithCategoryRepository) FindsByBlogId(db *gorm.DB, blogId int) ([]*models.ActiveBlogWithCategoryModel, error) {
	var acquisitionIds []*entities.TBLBlogWithCategoriesEntity
	if err := db.Where("blog_id = ?", blogId).Find(&acquisitionIds).Error; err != nil {
		return make([]*models.ActiveBlogWithCategoryModel, 0), err
	}
	fmt.Println(acquisitionIds)
	return repo.toModels(acquisitionIds)
}

func (repo *BlogWithCategoryRepository) Create(db *gorm.DB, obj *models.ActiveBlogWithCategoryModel) (*models.ActiveBlogWithCategoryModel, error) {
	bce, err := repo.toEntity(obj)
	if err != nil {
		return nil, errors.Wrap(errors.NewCustomError(), errors.REPO0001, err.Error())
	}
	if err := db.Create(&bce).Error; err != nil {
		return new(models.ActiveBlogWithCategoryModel), errors.Wrap(errors.NewCustomError(), errors.REPO0014, err.Error())
	}
	return repo.toModel(bce)
}

func (repo *BlogWithCategoryRepository) CreateAll(db *gorm.DB, objs []*models.ActiveBlogWithCategoryModel) ([]*models.ActiveBlogWithCategoryModel, error) {
	es, err := repo.toEntities(objs)
	if err != nil {
		return make([]*models.ActiveBlogWithCategoryModel, 0), errors.Wrap(errors.NewCustomError(), errors.REPO0001, err.Error())
	}
	if err := db.Create(&es).Error; err != nil {
		return make([]*models.ActiveBlogWithCategoryModel, 0), errors.Wrap(errors.NewCustomError(), errors.REPO0014, err.Error())
	}
	return repo.toModels(es)
}

func (repo *BlogWithCategoryRepository) DeleteById(db *gorm.DB, id string) error {
	if err := db.Where("composite_id = ?", id).Unscoped().Delete(&entities.TBLBlogWithCategoriesEntity{}).Error; err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.REPO0015, err.Error())
	}
	return nil
}

func (repo *BlogWithCategoryRepository) toEntity(obj *models.ActiveBlogWithCategoryModel) (*entities.TBLBlogWithCategoriesEntity, error) {
	return &entities.TBLBlogWithCategoriesEntity{
		CompositeId: obj.GetId(),
		CategoryId:  int(obj.GetCategoryId()),
		BlogId:      int(obj.GetBlogId()),
	}, nil
}

func (repo *BlogWithCategoryRepository) toEntities(obj []*models.ActiveBlogWithCategoryModel) ([]*entities.TBLBlogWithCategoriesEntity, error) {
	bes := make([]*entities.TBLBlogWithCategoriesEntity, len(obj))
	for i, obj := range obj {
		be, err := repo.toEntity(obj)
		if err != nil {
			return make([]*entities.TBLBlogWithCategoriesEntity, 0), err
		}
		bes[i] = be
	}
	return bes, nil
}

func (repo *BlogWithCategoryRepository) toModel(obj *entities.TBLBlogWithCategoriesEntity) (*models.ActiveBlogWithCategoryModel, error) {
	return models.NewActiveBlogWithCategoryModel(
		obj.CompositeId,
		obj.CategoryId,
		obj.BlogId,
	)
}

func (repo *BlogWithCategoryRepository) toModels(obj []*entities.TBLBlogWithCategoriesEntity) ([]*models.ActiveBlogWithCategoryModel, error) {
	models := make([]*models.ActiveBlogWithCategoryModel, len(obj))
	for i, obj := range obj {
		model, err := repo.toModel(obj)
		if err != nil {
			return nil, err
		}
		models[i] = model
	}
	return models, nil
}
