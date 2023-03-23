package blog

import (
	"github.com/set2002satoshi/my-site-api_v2/interfaces/database"
	"github.com/set2002satoshi/my-site-api_v2/interfaces/database/config"
	"github.com/set2002satoshi/my-site-api_v2/models"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/dto/response"
	"github.com/set2002satoshi/my-site-api_v2/usecase/service"
)

type BlogController struct {
	Interactor service.BlogInteractor
}

func NewBlogController(db config.DB) *BlogController {
	return &BlogController{
		Interactor: service.BlogInteractor{
			DB:                          &config.DBRepository{DB: db},
			UserRepo:                    &database.ActiveUserRepository{},
			BlogRepo:                    &database.ActiveBlogRepository{},
			HistoryBlogRepo:             &database.HistoryBlogRepository{},
			BlogWithCategoryRepo:        &database.BlogWithCategoryRepository{},
			HistoryBlogWithCategoryRepo: &database.HistoryBlogWithCategoryRepository{},
			CategoryRepo:                &database.ActiveCategoryRepository{},
			// HistoryCategoryRepo: &database.HistoryCategoryRepository{},
		},
	}
}

func (bu *BlogController) convertActiveBlogToDTO(obj *models.ActiveBlogModel) response.ActiveBlogEntity {
	categories := make([]*response.ActiveCategoryEntity, len(obj.GetCategoryIds()))
	for i, v := range obj.GetCategoryIds() {
		res := response.ActiveCategoryEntity{
			Id:           int(v.GetCategoryId()),
			CategoryName: v.GetCategoryName(),
		}
		categories[i] = &res
	}
	return response.ActiveBlogEntity{
		BlogId:     int(obj.GetBlogId()),
		UserId:     int(obj.GetUserId()),
		Nickname:   obj.GetNickname(),
		Title:      obj.GetTitle(),
		Context:    obj.GetContext(),
		Categories: categories,
		Option: response.Options{
			Revision:  int(obj.GetAuditTrail().GetRevision()),
			CreatedAt: obj.GetAuditTrail().GetCreatedAt(),
			UpdatedAt: obj.GetAuditTrail().GetUpdatedAt(),
		},
	}
}

func (bu *BlogController) convertActiveBlogToDTOs(obj []*models.ActiveBlogModel) []response.ActiveBlogEntity {
	BEs := make([]response.ActiveBlogEntity, len(obj))
	for i, v := range obj {
		BEs[i] = bu.convertActiveBlogToDTO(v)
	}
	return BEs
}
