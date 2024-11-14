package request

import "trash_report/entities"

type RegisterRequest struct {
	Name     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (registerRequest RegisterRequest) ToEntities() entities.User {
	return entities.User{
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
	}
}