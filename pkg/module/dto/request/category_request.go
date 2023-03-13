package request

type (
	CategoryFindByIdRequest struct {
		Id int `json:"id"`
	}
	CategoryDeleteByIdRequest struct {
		Id int `json:"id"`
	}
	CategoryCreateRequest struct {
		CategoryName string `json:"category_name"`
	}
)
