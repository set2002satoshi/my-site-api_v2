package models

import (
	"time"

	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/types"
)

type ActiveCategoryModel struct {
	categoryId   types.IDENTIFICATION
	categoryName string
	auditTrail   *types.AuditTrail
}

func NewActiveCategoryModel(
	categoryId int,
	categoryName string,
	revision int,
	createdAt time.Time,
	updatedAt time.Time,
) (*ActiveCategoryModel, error) {
	ac := new(ActiveCategoryModel)

	var err error
	err = errors.Combine(err, ac.setCategoryId(categoryId))
	err = errors.Combine(err, ac.setCategoryName(categoryName))
	if err != nil {
		return new(ActiveCategoryModel), err
	}
	com, err := types.NewAuditTrail(revision, createdAt, updatedAt)
	if err != nil {
		return new(ActiveCategoryModel), err
	}
	err = errors.Combine(err, ac.setAuditTrail(com))
	if err != nil {
		return new(ActiveCategoryModel), err
	}
	return ac, nil
}

func (ac *ActiveCategoryModel) GetCategoryId() types.IDENTIFICATION {
	return ac.categoryId
}

func (ac *ActiveCategoryModel) GetCategoryName() string {
	return ac.categoryName
}

func (ac *ActiveCategoryModel) GetAuditTrail() *types.AuditTrail {
	return ac.auditTrail
}

func (ac *ActiveCategoryModel) setCategoryId(categoryId int) error {
	i, err := types.NewIDENTIFICATION(categoryId)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0001, err.Error())
	}
	ac.categoryId = i
	return nil
}

func (ac *ActiveCategoryModel) setCategoryName(categoryName string) error {
	ac.categoryName = categoryName
	return nil
}

func (ac *ActiveCategoryModel) setAuditTrail(auditTrail *types.AuditTrail) error {
	ac.auditTrail = auditTrail
	return nil
}
