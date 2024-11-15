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
		Tag:       s.Tag,
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

func (u *UserMe) FromDomainToAPI() (userMe *UserMe) {
	if u == nil {
		return
	}

	return &UserMe{
		IdUser:      u.IdUser,
		OnboardSet:  u.OnboardSet,
		PasswordSet: u.PasswordSet,
	}
}
