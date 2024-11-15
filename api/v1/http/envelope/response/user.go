package response

import "bitbucket.org/asadventure/be-core-lib/response"

// swagger:model swaggerUserProfileResponse
type swaggerUserProfileResponse struct { //nolint:all
	response.Response
	Data UserProfileResponse `json:"data"`
}

//swagger:model UserProfileResponse
type UserProfileResponse struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Tag         string `json:"tag"`
	Avatar      string `json:"avatar"`
}

// swagger:model swaggerUserProfileResponse
type swaggerUserMeResponse struct { //nolint:all
	response.Response
	Data UserMeResponse `json:"data"`
}

//swagger:model UserProfileResponse
type UserMeResponse struct {
	IdUser  int  `json:"id_user"`
	Onboard bool `json:"onboard"`
}
