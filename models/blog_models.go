package models

import (
	"time"

	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/types"
)

type ActiveBlogModel struct {
	blogId     types.IDENTIFICATION
	userId     types.IDENTIFICATION
	title      string
	context    string
	auditTrail *types.AuditTrail
}

func NewActiveBlogModel(
	blogId int,
	userId int,
	title string,
	context string,
	revision int,
	createdAt time.Time,
	updatedAt time.Time,
) (*ActiveBlogModel, error) {
	abm := new(ActiveBlogModel)

	var err error
	err = errors.Combine(err, abm.setBlogId(blogId))
	err = errors.Combine(err, abm.setUserId(userId))
	err = errors.Combine(err, abm.setTitle(title))
	err = errors.Combine(err, abm.setContext(context))
	if err != nil {
		return new(ActiveBlogModel), err
	}
	com, err := types.NewAuditTrail(revision, createdAt, updatedAt)
	if err != nil {
		return new(ActiveBlogModel), err
	}
	err = errors.Combine(err, abm.setAuditTrail(com))
	if err != nil {
		return new(ActiveBlogModel), err
	}

	if err := abm.Validation(); err != nil {
		return new(ActiveBlogModel), err
	}

	return abm, nil

}

func (abm *ActiveBlogModel) GetBlogId() types.IDENTIFICATION {
	return abm.blogId
}

func (abm *ActiveBlogModel) GetUserId() types.IDENTIFICATION {
	return abm.userId
}

func (abm *ActiveBlogModel) GetTitle() string {
	return abm.title
}

func (abm *ActiveBlogModel) GetContext() string {
	return abm.context
}

func (abm *ActiveBlogModel) GetAuditTrail() *types.AuditTrail {
	return abm.auditTrail
}

func (abm *ActiveBlogModel) setBlogId(blogId int) error {
	i, err := types.NewIDENTIFICATION(blogId)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0001, err.Error())
	}
	abm.blogId = i
	return nil
}

func (abm *ActiveBlogModel) setUserId(userId int) error {
	i, err := types.NewOneOrMoreIDENTIFICATION(userId)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0001, err.Error())
	}
	abm.userId = i
	return nil
}

func (abm *ActiveBlogModel) setTitle(Title string) error {
	abm.title = Title
	return nil
}

func (abm *ActiveBlogModel) setContext(Context string) error {
	abm.context = Context
	return nil
}

func (abm *ActiveBlogModel) setAuditTrail(auditTrail *types.AuditTrail) error {
	abm.auditTrail = auditTrail
	return nil
}

func (abm *ActiveBlogModel) Validation() error {
	return nil
}
