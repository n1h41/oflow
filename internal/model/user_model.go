package model

type SignUpUserReq struct {
	Email     string `json:"email" validate:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Password  string `json:"password" validate:"required"`
}

type SignInUserReq struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type FetchIdentityCredentialsReq struct {
	Token string `json:"token" validate:"required"`
}

type ConfirmUserReq struct {
	Email            string `json:"email"`
	ConfirmationCode string `json:"confirmation_code"`
}

type AddDeviceReq struct {
	DeviceMAC string `json:"device_mac" validate:"required"`
}

type ListUserDevicesReq struct {
	UserId string `json:"user_id" validate:"required"`
}

type AttachIOTPolicyReq struct {
	IdentityId string `json:"identity_id" validate:"required"`
}
