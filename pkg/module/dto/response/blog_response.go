package response

import "github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"

type (
	FindAllActiveBlogResponse struct {
		Result ActiveBlogResults `json:"result"`

		Errors []errors.ErrorInfo
	}

	CreateActiveBlogResponse struct {
		Result ActiveBlogResult `json:"results"`

		Errors []errors.ErrorInfo
	}

	UpdateActiveBlogResponse struct {
		Result ActiveBlogResult `json:"results"`

		Errors []errors.ErrorInfo
	}
	DeleteByIdActiveBlogResponse struct {
		Errors []errors.ErrorInfo
	}
)

type (
	ActiveBlogResult struct {
		Blog ActiveBlogEntity `json:"blog"`
	}
	ActiveBlogResults struct {
		Blogs []ActiveBlogEntity `json:"blogs"`
	}

	// HistoryBlogResult struct {
	// 	Student *HistoryUserEntity `json:"student"`
	// }
	// HistoryBlogResults struct {
	// 	Students []*HistoryBlogEntity `json:"students"`
	// }

)

type (
	ActiveBlogEntity struct {
		BlogId     int                   `json:"blog_id"`
		UserId     int                   `json:"user_id"`
		Nickname   string                `json:"nickname"`
		Title      string                `json:"title"`
		Context    string                `json:"content"`
		// Categories []ActiveCategoryModel `json:"categories"`
		Option     Options               `json:"option"`
	}
)

type (
	BlogAndCategoryEntity struct {
		Id         int `json:"id"`
		BlogId     int `json:"blog_id"`
		CategoryId int `json:"category_id"`
	}
)
