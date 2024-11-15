package v1

type User struct {
	IdUser    int    `json:"id_user"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Tag       string `json:"tag"`
	Avatar    string `json:"avatar"`
}

type UserProfile struct {
	IdUser      int    `json:"id_user"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Tag         string `json:"tag"`
	Avatar      string `json:"avatar"`
	OnboardSet  bool   `json:"onboard_set"`
	PasswordSet bool   `json:"password_set"`
}

type Onboard struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Avatar          string `json:"avatar"`
	PhoneNumberCode string `json:"phone_number_code"`
	PhoneNumber     string `json:"phone_number"`
}

type UpdateAuth struct {
	IdUser          int    `json:"id_user"`
	Email           string `json:"email"`
	PhoneNumberCode string `json:"phone_number_code"`
	PhoneNumber     string `json:"phone_number"`
}

type UserMe struct {
	IdUser      int  `json:"id_user"`
	OnboardSet  bool `json:"onboard_set"`
	PasswordSet bool `json:"password_set"`
}
