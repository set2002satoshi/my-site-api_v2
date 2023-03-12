package models

import (
	"time"

	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/types"
)

type HistoryCategoryModel struct {
	Id           types.IDENTIFICATION
	activeId     types.IDENTIFICATION
	categoryName string
	auditTrail   *types.AuditTrail
}

func NewHistoryCategoryModel(
	id int,
	activeId int,
	categoryName string,
	revision int,
	createdAt time.Time,
	updatedAt time.Time,
) (*HistoryCategoryModel, error) {
	hc := new(HistoryCategoryModel)

	var err error
	err = errors.Combine(err, hc.setId(id))
	err = errors.Combine(err, hc.setActiveId(activeId))
	err = errors.Combine(err, hc.setCategoryName(categoryName))
	if err != nil {
		return new(HistoryCategoryModel), err
	}
	com, err := types.NewAuditTrail(revision, createdAt, updatedAt)
	if err != nil {
		return new(HistoryCategoryModel), err
	}
	err = errors.Combine(err, hc.setAuditTrail(com))
	if err != nil {
		return new(HistoryCategoryModel), err
	}
	return hc, nil
}

func (hc *HistoryCategoryModel) GetId() types.IDENTIFICATION {
	return hc.Id
}

func (hc *HistoryCategoryModel) GetActiveId() types.IDENTIFICATION {
	return hc.activeId
}

func (hc *HistoryCategoryModel) GetCategoryName() string {
	return hc.categoryName
}

func (hc *HistoryCategoryModel) GetAuditTrail() *types.AuditTrail {
	return hc.auditTrail
}

func (hc *HistoryCategoryModel) setId(categoryId int) error {
	i, err := types.NewIDENTIFICATION(categoryId)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0001, err.Error())
	}
	hc.Id = i
	return nil
}

func (hc *HistoryCategoryModel) setActiveId(activeId int) error {
	i, err := types.NewOneOrMoreIDENTIFICATION(activeId)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0001, err.Error())
	}
	hc.activeId = i
	return nil
}

func (hc *HistoryCategoryModel) setCategoryName(categoryName string) error {
	hc.categoryName = categoryName
	return nil
}

func (hc *HistoryCategoryModel) setAuditTrail(auditTrail *types.AuditTrail) error {
	hc.auditTrail = auditTrail
	return nil
}
