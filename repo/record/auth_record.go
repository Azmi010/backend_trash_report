package record

import "trash_report/entities"

type User struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Email    string
	Password string
	Reports  []Report `gorm:"foreignKey:UserID"`
}

func FromEntities(user entities.User) User {
	return User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func (user User) ToEntities() entities.User {
	return entities.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}