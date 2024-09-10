package model

type SignUpUserReq struct {
	Email     string `json:"email" validate:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Password  string `json:"password" validate:"required"`
}

type SignInUserReq struct {
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
}
