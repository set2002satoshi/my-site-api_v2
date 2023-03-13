package database

import (
	"github.com/set2002satoshi/my-site-api_v2/models"
	"github.com/set2002satoshi/my-site-api_v2/models/entities"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/types"
	"gorm.io/gorm"
)

type ActiveBlogRepository struct{}

type BlogModel struct {
	BlogId     types.IDENTIFICATION
	UserId     types.IDENTIFICATION
	Nickname   string
	Title      string
	Context    string
	AuditTrail *types.AuditTrail
}

func (repo *ActiveBlogRepository) FindById(db *gorm.DB, id int) (*models.ActiveBlogModel, error) {
	var blog *entities.BlogWithNicknameEntity
	if err := db.Model(&entities.TBLBlogEntity{}).Select("tbl_blog_entities.*, tbl_user_entities.nickname").Joins("left join tbl_user_entities on tbl_blog_entities.user_id = tbl_user_entities.user_id").Scan(&blog).Error; err != nil {
		return new(models.ActiveBlogModel), nil
	}
	return repo.joinTBLToModel(blog)
}

func (repo *ActiveBlogRepository) FindAll(db *gorm.DB) ([]*models.ActiveBlogModel, error) {
	var blogs []entities.BlogWithNicknameEntity
	if err := db.Model(&entities.TBLBlogEntity{}).Select("tbl_blog_entities.*, tbl_user_entities.nickname").Joins("left join tbl_user_entities on tbl_blog_entities.user_id = tbl_user_entities.user_id").Scan(&blogs).Error; err != nil {
		return make([]*models.ActiveBlogModel, 0), nil
	}
	return repo.joinTBLToModels(blogs)
}

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

func (repo *ActiveBlogRepository) Update(tx *gorm.DB, obj *models.ActiveBlogModel) (*models.ActiveBlogModel, error) {
	be, err := repo.toEntity(obj)
	if err != nil {
		return new(models.ActiveBlogModel), errors.Wrap(errors.NewCustomError(), errors.REPO0001, err.Error())
	}
	if err := tx.Select("title", "context", "revision", "updated_at").Updates(&be).Error; err != nil {
		return nil, errors.Wrap(errors.NewCustomError(), errors.REPO0004, err.Error())
	}
	return repo.toModel(be)
}

func (repo *ActiveBlogRepository) DeleteById(tx *gorm.DB, id int) error {
	if err := tx.Unscoped().Delete(&entities.TBLBlogEntity{}, id).Error; err != nil {
		return err
	}
	return nil
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

func (repo *ActiveBlogRepository) joinTBLToModel(obj *entities.BlogWithNicknameEntity) (*models.ActiveBlogModel, error) {
	return models.NewActiveBlogModel(
		obj.BlogId,
		obj.UserId,
		obj.Nickname,
		obj.Title,
		obj.Context,
		obj.Revision,
		obj.CreatedAt,
		obj.UpdatedAt,
	)
}

func (repo *ActiveBlogRepository) joinTBLToModels(obj []entities.BlogWithNicknameEntity) ([]*models.ActiveBlogModel, error) {
	blogModels := make([]*models.ActiveBlogModel, len(obj))
	for i, obj := range obj {
		model, err := repo.joinTBLToModel(&obj)
		if err != nil {
			return nil, err
		}
		blogModels[i] = model
	}
	return blogModels, nil
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
