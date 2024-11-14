package response

import "trash_report/entities"

type RegisterResponse struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Email string `json:"email"`
}

func FromRegisterEntities(user entities.User) RegisterResponse {
	return RegisterResponse{
		ID:    user.ID,
		Nama:  user.Name,
		Email: user.Email,
	}
}