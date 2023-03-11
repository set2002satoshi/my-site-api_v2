package request


type (
	BlogFindByIdRequest struct {
		Id int `json:"id"`
	}
	BlogDeleteByIdRequest struct {
		Id int `json:"id"`
	}
	BlogCreateRequest struct {
		Title string `json:"title"`
		Context string `json:"context"`
	}
	BlogUpdateRequest struct {
		BlogId int `json:"blogId"`
		Title string `json:"title"`
		Context string `json:"context"`
		Revision int `json:"revision"`
	}
)