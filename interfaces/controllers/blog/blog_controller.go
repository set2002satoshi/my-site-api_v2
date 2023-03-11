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
			DB:              &config.DBRepository{DB: db},
			UserRepo:        &database.ActiveUserRepository{},
			BlogRepo:        &database.ActiveBlogRepository{},
			HistoryBlogRepo: &database.HistoryBlogRepository{},
		},
	}
}

func (bu *BlogController) convertActiveBlogToDTO(obj *models.ActiveBlogModel) response.ActiveBlogEntity {
	return response.ActiveBlogEntity{
		BlogId:   int(obj.GetBlogId()),
		UserId:   int(obj.GetUserId()),
		Nickname: obj.GetNickname(),
		Title:    obj.GetTitle(),
		Context:  obj.GetContext(),
		Option: response.Options{
			Revision:  int(obj.GetAuditTrail().GetRevision()),
			CreatedAt: obj.GetAuditTrail().GetCreatedAt(),
			UpdatedAt: obj.GetAuditTrail().GetUpdatedAt(),
		},
	}
}
