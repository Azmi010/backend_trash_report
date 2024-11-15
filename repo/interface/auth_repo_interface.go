package repoInterface

import "trash_report/entities"

type AuthRepository interface {
	Login(user entities.User) (entities.User, error)
	Register(user entities.User) (entities.User, error)
}