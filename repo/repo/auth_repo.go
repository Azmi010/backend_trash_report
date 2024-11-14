package repo

import (
	"trash_report/entities"
	"trash_report/repo/record"

	"gorm.io/gorm"
)

func NewAuthRepo(db *gorm.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

type AuthRepo struct {
	db *gorm.DB
}

func (authRepo AuthRepo) Login(user entities.User) (entities.User, error) {
	var userDb entities.User
	result := authRepo.db.First(&userDb, "email = ?", user.Email)
	if result.Error != nil {
		return entities.User{}, result.Error
	}
	return userDb, nil
}

func (authRepo AuthRepo) Register(user entities.User) (entities.User, error) {
	userDb := record.FromEntities(user)
	result := authRepo.db.Create(&userDb)
	if result.Error != nil {
		return entities.User{}, result.Error
	}
	return userDb.ToEntities(), nil
}