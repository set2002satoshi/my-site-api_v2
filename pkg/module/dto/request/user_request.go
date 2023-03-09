package request

type (
	UserFindByIdRequest struct {
		Id int `json:"id"`
	}
	UserDeleteRequest struct {
		Id int `json:"id"`
	}
	UserCreateRequest struct {
		Email    string `form:"email" json:"email"`
		Name     string `form:"name" json:"name"`
		Password string `form:"pass" json:"pass"`
		Roll     string `form:"roll" json:"roll"`
	}

	UserUpdateRequest struct {
		Id       int    `json:"id"`
		Email    string `form:"email" json:"email"`
		Name     string `form:"name" json:"name"`
		Password string `form:"pass" json:"pass"`
		Roll     string `form:"roll" json:"roll"`
		Revision int    `json:"revision"`
	}
)

type (
	UserLoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"pass"`
	}
)
