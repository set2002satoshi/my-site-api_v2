package models

import (
	"time"

	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/types"
)

type HistoryBlogModel struct {
	blogId      types.IDENTIFICATION
	activeId    types.IDENTIFICATION
	userId      types.IDENTIFICATION
	title       string
	context     string
	categoryIds []*HistoryCategoryModel
	auditTrail  *types.AuditTrail
}

func NewHistoryBlogModel(
	blogId int,
	activeId int,
	userId int,
	title string,
	context string,
	categoryIds []*HistoryCategoryModel,
	revision int,
	createdAt time.Time,
	updatedAt time.Time,
) (*HistoryBlogModel, error) {
	hbm := new(HistoryBlogModel)

	var err error
	err = errors.Combine(err, hbm.setBlogId(blogId))
	err = errors.Combine(err, hbm.setActiveId(activeId))
	err = errors.Combine(err, hbm.setUserId(userId))
	err = errors.Combine(err, hbm.setTitle(title))
	err = errors.Combine(err, hbm.setContext(context))
	err = errors.Combine(err, hbm.setCategoryIds(categoryIds))
	if err != nil {
		return new(HistoryBlogModel), err
	}
	com, err := types.NewAuditTrail(revision, createdAt, updatedAt)
	if err != nil {
		return new(HistoryBlogModel), err
	}
	err = errors.Combine(err, hbm.setAuditTrail(com))
	if err != nil {
		return new(HistoryBlogModel), err
	}

	if err := hbm.Validation(); err != nil {
		return new(HistoryBlogModel), err
	}

	return hbm, nil

}

func (hbm *HistoryBlogModel) GetBlogId() types.IDENTIFICATION {
	return hbm.blogId
}

func (hbm *HistoryBlogModel) GetUserId() types.IDENTIFICATION {
	return hbm.userId
}

func (hbm *HistoryBlogModel) GetActiveId() types.IDENTIFICATION {
	return hbm.activeId

}

func (hbm *HistoryBlogModel) GetTitle() string {
	return hbm.title
}

func (hbm *HistoryBlogModel) GetContext() string {
	return hbm.context
}

func (hbm *HistoryBlogModel) GetCategoryIds() []*HistoryCategoryModel {
	return hbm.categoryIds
}

func (hbm *HistoryBlogModel) GetAuditTrail() *types.AuditTrail {
	return hbm.auditTrail
}

func (hbm *HistoryBlogModel) setBlogId(blogId int) error {
	i, err := types.NewIDENTIFICATION(blogId)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0001, err.Error())
	}
	hbm.blogId = i
	return nil
}

func (hbm *HistoryBlogModel) setActiveId(activeId int) error {
	i, err := types.NewOneOrMoreIDENTIFICATION(activeId)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0001, err.Error())
	}
	hbm.activeId = i
	return nil
}

func (hbm *HistoryBlogModel) setUserId(UserId int) error {
	i, err := types.NewOneOrMoreIDENTIFICATION(UserId)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0001, err.Error())
	}
	hbm.userId = i
	return nil
}

func (hbm *HistoryBlogModel) setTitle(Title string) error {
	hbm.title = Title
	return nil
}

func (hbm *HistoryBlogModel) setContext(Context string) error {
	hbm.context = Context
	return nil
}

func (hbm *HistoryBlogModel) setCategoryIds(ids []*HistoryCategoryModel) error {
	hbm.categoryIds = ids
	return nil
}

func (hbm *HistoryBlogModel) setAuditTrail(auditTrail *types.AuditTrail) error {
	hbm.auditTrail = auditTrail
	return nil
}

func (hbm *HistoryBlogModel) Validation() error {
	return nil
}
