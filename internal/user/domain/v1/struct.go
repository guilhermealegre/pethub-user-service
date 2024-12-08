package v1

import "github.com/google/uuid"

type User struct {
	IdUser    int    `json:"id_user"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Tag       string `json:"tag"`
	Avatar    string `json:"avatar"`
}

type UserProfile struct {
	UserID     int       `json:"user_id"`
	UserUUID   uuid.UUID `json:"user_uuid"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	Avatar     string    `json:"avatar"`
	OnboardSet bool      `json:"onboard_set"`
	Animals    []Animal  `json:"animals"`
}

type Animal struct {
	AnimalID int    `json:"animal_id"`
	Name     string `json:"name"`
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
	UserId     int       `json:"user_id"`
	UserUUID   uuid.UUID `json:"user_uuid"`
	Email      string    `json:"email"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	OnboardSet bool      `json:"onboard_set"`
}
