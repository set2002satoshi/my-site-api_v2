package models

import (
	"time"

	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/types"
)

type HistoryBlogModel struct {
	blogId     types.IDENTIFICATION
	activeId   types.IDENTIFICATION
	userId     types.IDENTIFICATION
	title      string
	context    string
	auditTrail *types.AuditTrail
}

func NewHistoryBlogModel(
	blogId int,
	activeId int,
	userId int,
	title string,
	context string,
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

func (abm *HistoryBlogModel) GetBlogId() types.IDENTIFICATION {
	return abm.blogId
}

func (abm *HistoryBlogModel) GetUserId() types.IDENTIFICATION {
	return abm.userId
}

func (abm *HistoryBlogModel) GetActiveId() types.IDENTIFICATION {
	return abm.activeId

}

func (abm *HistoryBlogModel) GetTitle() string {
	return abm.title
}

func (abm *HistoryBlogModel) GetContext() string {
	return abm.context
}

func (abm *HistoryBlogModel) GetAuditTrail() *types.AuditTrail {
	return abm.auditTrail
}

func (abm *HistoryBlogModel) setBlogId(blogId int) error {
	i, err := types.NewIDENTIFICATION(blogId)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0001, err.Error())
	}
	abm.blogId = i
	return nil
}

func (abm *HistoryBlogModel) setActiveId(activeId int) error {
	i, err := types.NewOneOrMoreIDENTIFICATION(activeId)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0001, err.Error())
	}
	abm.activeId = i
	return nil
}

func (abm *HistoryBlogModel) setUserId(UserId int) error {
	i, err := types.NewOneOrMoreIDENTIFICATION(UserId)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0001, err.Error())
	}
	abm.userId = i
	return nil
}

func (abm *HistoryBlogModel) setTitle(Title string) error {
	abm.title = Title
	return nil
}

func (abm *HistoryBlogModel) setContext(Context string) error {
	abm.context = Context
	return nil
}

func (abm *HistoryBlogModel) setAuditTrail(auditTrail *types.AuditTrail) error {
	abm.auditTrail = auditTrail
	return nil
}

func (abm *HistoryBlogModel) Validation() error {
	return nil
}
