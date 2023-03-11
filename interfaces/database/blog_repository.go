package database

import (
	"github.com/set2002satoshi/my-site-api_v2/models"
	"github.com/set2002satoshi/my-site-api_v2/models/entities"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/types"
	"gorm.io/gorm"
)

type ActiveBlogRepository struct{}

func (repo *ActiveBlogRepository) Create(db *gorm.DB, obj *models.ActiveBlogModel) (*models.ActiveBlogModel, error) {
	be, err := repo.toEntity(obj)
	if err != nil {
		return new(models.ActiveBlogModel), errors.Wrap(errors.NewCustomError(), errors.REPO0001, err.Error())
	}
	if err := db.Create(be).Error; err != nil {
		return new(models.ActiveBlogModel), errors.Wrap(errors.NewCustomError(), errors.REPO0002, err.Error())
	}
	return repo.toModel(be)
}




func (repo *ActiveBlogRepository) toModel(obj *entities.TBLBlogEntity) (*models.ActiveBlogModel, error) {
	return models.NewActiveBlogModel(
		obj.BlogId,
		obj.UserId,
		types.DEFAULT_NAME,
		obj.Title,
		obj.Context,
		obj.Revision,
		obj.CreatedAt,
		obj.UpdatedAt,
	)
}

func (repo *ActiveBlogRepository) toEntity(obj *models.ActiveBlogModel) (*entities.TBLBlogEntity, error) {
	return &entities.TBLBlogEntity{
		BlogId:    int(obj.GetBlogId()),
		UserId:    int(obj.GetUserId()),
		Title:     obj.GetTitle(),
		Context:   obj.GetContext(),
		Revision:  int(obj.GetAuditTrail().GetRevision()),
		CreatedAt: obj.GetAuditTrail().GetCreatedAt(),
		UpdatedAt: obj.GetAuditTrail().GetUpdatedAt(),
	}, nil
}
