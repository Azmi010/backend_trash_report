package request

import "trash_report/entities"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (loginRequest LoginRequest) ToEntities() entities.User {
	return entities.User{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	}
}