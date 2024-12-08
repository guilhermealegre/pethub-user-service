package v1

import (
	"github.com/guilhermealegre/pethub-user-service/api/v1/http/envelope/request"
	"github.com/guilhermealegre/pethub-user-service/api/v1/http/envelope/response"
)

func (o *Onboard) FromAPIToDomain(req *request.OnboardRequest) {
	if o == nil || req == nil {
		return
	}

	o.FirstName = req.Body.FirstName
	o.LastName = req.Body.LastName
	o.Avatar = req.Body.Avatar
	o.PhoneNumberCode = req.Body.PhoneNumberCode
	o.PhoneNumber = req.Body.PhoneNumber
}

func (s *UserProfile) FromDomainToAPI() (userProfile *response.UserProfileResponse) {
	if s == nil {
		return
	}

	userProfile = &response.UserProfileResponse{
		FirstName: s.FirstName,
		LastName:  s.LastName,
		Avatar:    s.Avatar,
		Email:     s.Email,
	}

	return userProfile
}

func (s *UserProfile) FromAPIToDomain(req *request.UpdateUserProfileRequest) (userProfile *UserProfile) {
	if s == nil {
		return
	}

	userProfile = &UserProfile{
		FirstName: req.Body.FirstName,
		LastName:  req.Body.LastName,
	}

	return userProfile
}

func (u *UserMe) FromDomainToAPI() (userMe *response.UserMeResponse) {
	if u == nil {
		return
	}

	return &response.UserMeResponse{
		UserId:     u.UserId,
		UserUUID:   u.UserUUID,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		Email:      u.Email,
		OnboardSet: u.OnboardSet,
	}
}
