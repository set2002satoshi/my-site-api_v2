package database

import (
	"time"

	"github.com/set2002satoshi/my-site-api_v2/models"
	"github.com/set2002satoshi/my-site-api_v2/models/entities"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"gorm.io/gorm"
)

type HistoryBlogRepository struct{}

func (hbr *HistoryBlogRepository) Create(tx *gorm.DB, obj *models.HistoryBlogModel) (*models.HistoryBlogModel, error) {
	hbe, err := hbr.toEntity(obj)
	if err != nil {
		return nil, errors.Wrap(errors.NewCustomError(), errors.REPO0001, err.Error())
	}
	if err := tx.Create(hbe).Error; err != nil {
		return new(models.HistoryBlogModel), errors.Wrap(errors.NewCustomError(), errors.REPO0002, err.Error())
	}
	return hbr.toModel(hbe)
}

func (hbr *HistoryBlogRepository) toModel(obj *entities.HistoryBlogEntity) (*models.HistoryBlogModel, error) {
	return models.NewHistoryBlogModel(
		obj.Id,
		obj.ActiveId,
		obj.UserId,
		obj.Title,
		obj.Context,
		[]*models.HistoryCategoryModel{},
		obj.Revision,
		obj.CreatedAt,
		obj.DeletedAt,
	)
}

func (hbr *HistoryBlogRepository) toEntity(obj *models.HistoryBlogModel) (*entities.HistoryBlogEntity, error) {
	return &entities.HistoryBlogEntity{
		Id:        int(obj.GetBlogId()),
		ActiveId:  int(obj.GetActiveId()),
		UserId:    int(obj.GetUserId()),
		Title:     obj.GetTitle(),
		Context:   obj.GetContext(),
		Revision:  int(obj.GetAuditTrail().GetRevision()),
		CreatedAt: obj.GetAuditTrail().GetCreatedAt(),
		DeletedAt: time.Now().UTC(),
	}, nil
}
