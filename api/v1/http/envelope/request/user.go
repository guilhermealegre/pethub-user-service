package request

//swagger:parameters UpdateUserProfileRequest
type UpdateUserProfileRequest struct {
	IdUser int `json:"id_user" validate:"required,gte=1"`
	// Body
	// in: body
	// required: true
	Body struct {
		// FirstName
		// required: true
		FirstName string `json:"first_name" validate:"required,gte=1"`
		// LastName
		// required: true
		LastName string `json:"last_name" validate:"required,gte=1"`
	}
}

//swagger:parameters OnboardRequest
type OnboardRequest struct {
	// Body
	// in: body
	// required: true
	Body struct {
		// FirstName
		// required: true
		FirstName string `json:"first_name" validate:"required,gte=1"`
		// LastName
		// required: true
		LastName string `json:"last_name" validate:"required,gte=1"`
		// Avatar
		// required: false
		Avatar string `json:"avatar"`
		// PhoneNumberCode
		// required: false
		PhoneNumberCode string `json:"phone_number_code"`
		// PhoneNumber
		// required: true
		PhoneNumber string `json:"phone_number"`
	}
}
