package response

import "github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"

type (
	FindAllCategoryResponse struct {
		Results ActiveCategoryResults
		Errors  []errors.ErrorInfo
	}
	FindByIdCategoryResponse struct {
		Result ActiveCategoryResult
		Errors []errors.ErrorInfo
	}
	CreateCategoryResponse struct {
		Result ActiveCategoryResult
		Errors []errors.ErrorInfo
	}

	DeleteByIdCategoryResponse struct {
		Errors []errors.ErrorInfo
	}
)

type (
	ActiveCategoryResult struct {
		Category ActiveCategoryEntity `json:"category"`
	}
	ActiveCategoryResults struct {
		Categories []ActiveCategoryEntity `json:"categories"`
	}
)

type (
	ActiveCategoryEntity struct {
		Id           int    `json:"id"`
		CategoryName string `json:"category_name"`
	}
)
