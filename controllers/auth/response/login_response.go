package response

import "trash_report/entities"

type AuthResponse struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FromLoginEntities(user entities.User) AuthResponse {
	return AuthResponse{
		ID:    user.ID,
		Nama:  user.Name,
		Email: user.Email,
		Token: user.Token,
	}
}