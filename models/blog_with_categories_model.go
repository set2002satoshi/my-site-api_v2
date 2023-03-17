package models

import (
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/types"
)

type ActiveBlogWithCategoryModel struct {
	Id         types.IDENTIFICATION
	CategoryId types.IDENTIFICATION
	BlogId     types.IDENTIFICATION
}

func NewActiveBlogWithCategoryModel(
	id int,
	categoryId int,
	blogId int,
) (*ActiveBlogWithCategoryModel, error) {
	abc := new(ActiveBlogWithCategoryModel)

	var err error
	err = errors.Combine(err, abc.setId(id))
	err = errors.Combine(err, abc.setCategoryId(categoryId))
	err = errors.Combine(err, abc.setBlogId(blogId))
	if err != nil {
		return new(ActiveBlogWithCategoryModel), err
	}
	return abc, nil
}

func (abc *ActiveBlogWithCategoryModel) GetId() types.IDENTIFICATION {
	return abc.Id
}

func (abc *ActiveBlogWithCategoryModel) GetCategoryId() types.IDENTIFICATION {
	return abc.CategoryId
}

func (abc *ActiveBlogWithCategoryModel) GetBlogId() types.IDENTIFICATION {
	return abc.BlogId
}

func (abc *ActiveBlogWithCategoryModel) setId(id int) error {
	i, err := types.NewIDENTIFICATION(id)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0001, err.Error())
	}
	abc.Id = i
	return nil
}

func (abc *ActiveBlogWithCategoryModel) setCategoryId(categoryId int) error {
	i, err := types.NewIDENTIFICATION(categoryId)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0001, err.Error())
	}
	abc.CategoryId = i
	return nil
}

func (abc *ActiveBlogWithCategoryModel) setBlogId(blogId int) error {
	i, err := types.NewIDENTIFICATION(blogId)
	if err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.EN0001, err.Error())
	}
	abc.BlogId = i
	return nil
}
