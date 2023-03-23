package models

import (
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/types"
)

type HistoryBlogWithCategoryModel struct {
	Id         types.IDENTIFICATION
	ActiveId   string
	CategoryId types.IDENTIFICATION
	BlogId     types.IDENTIFICATION
}

func NewHistoryBlogWithCategoryModel(
	id int,
	activeId string,
	categoryId int,
	blogId int,
) (*HistoryBlogWithCategoryModel, error) {
	hbc := new(HistoryBlogWithCategoryModel)

	var err error
	err = errors.Combine(err, hbc.setId(id))
	err = errors.Combine(err, hbc.setActiveId(activeId))
	err = errors.Combine(err, hbc.setCategoryId(categoryId))
	err = errors.Combine(err, hbc.setBlogId(blogId))
	if err != nil {
		return new(HistoryBlogWithCategoryModel), err
	}
	return hbc, nil
}

func (abc *HistoryBlogWithCategoryModel) GetId() types.IDENTIFICATION {
	return abc.Id
}

func (abc *HistoryBlogWithCategoryModel) GetActiveId() string {
	return abc.ActiveId
}

func (abc *HistoryBlogWithCategoryModel) GetCategoryId() types.IDENTIFICATION {
	return abc.CategoryId
}

func (abc *HistoryBlogWithCategoryModel) GetBlogId() types.IDENTIFICATION {
	return abc.BlogId
}

func (abc *HistoryBlogWithCategoryModel) setId(id int) error {
	i, err := types.NewIDENTIFICATION(id)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0001, err.Error())
	}
	abc.Id = i
	return nil
}

func (abc *HistoryBlogWithCategoryModel) setActiveId(activeId string) error {
	if activeId == "" {
		return errors.Add(errors.NewCustomError(), errors.EN0007)
	}
	// 他にもバリデーションを考えよう
	abc.ActiveId = activeId
	return nil
}

func (abc *HistoryBlogWithCategoryModel) setCategoryId(categoryId int) error {
	i, err := types.NewIDENTIFICATION(categoryId)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0001, err.Error())
	}
	abc.CategoryId = i
	return nil
}

func (abc *HistoryBlogWithCategoryModel) setBlogId(blogId int) error {
	i, err := types.NewIDENTIFICATION(blogId)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0001, err.Error())
	}
	abc.BlogId = i
	return nil
}
