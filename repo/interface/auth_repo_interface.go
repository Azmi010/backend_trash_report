package repoInterface

import "trash_report/entities"

type ReportRepository interface {
	Login(user entities.User) (entities.User, error)
	Register(user entities.User) (entities.User, error)
}