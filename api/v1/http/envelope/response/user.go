package response

import (
	"bitbucket.org/asadventure/be-core-lib/response"
	"github.com/google/uuid"
)

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
	UserId     int       `json:"user_id"`
	UserUUID   uuid.UUID `json:"user_uuid"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	OnboardSet bool      `json:"onboard_set"`
}
