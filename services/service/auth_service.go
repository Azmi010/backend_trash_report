package service

import (
	"errors"
	"trash_report/constant"
	"trash_report/entities"
	"trash_report/middleware"
	serviceInterface "trash_report/services/interface"

	"golang.org/x/crypto/bcrypt"
)

func NewAuthService(ar serviceInterface.AuthInterface, jt middleware.JwtInterface) *AuthService {
	return &AuthService{
		authRepoInterface: ar,
		jwtInterface:      jt,
	}
}

type AuthService struct {
	authRepoInterface serviceInterface.AuthInterface
	jwtInterface      middleware.JwtInterface
}

func (authService AuthService) Login(user entities.User) (entities.User, error) {
	if user.Email == "" {
		return entities.User{}, constant.EMAIL_IS_EMPTY
	} else if user.Password == "" {
		return entities.User{}, constant.PASSWORD_IS_EMPTY
	}

	storedUser, err := authService.authRepoInterface.Login(user)
	if err != nil {
		return entities.User{}, err
	}

	if !CheckPasswordHash(user.Password, storedUser.Password) {
		return entities.User{}, errors.New("invalid credentials")
	}

	token, err := authService.jwtInterface.GenerateJWT(storedUser.ID, storedUser.Name)
	if err != nil {
		return entities.User{}, err
	}
	storedUser.Token = token

	return storedUser, nil
}

func (authService AuthService) Register(user entities.User) (entities.User, error) {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return entities.User{}, errors.New("failed to hash password")
	}
	user.Password = hashedPassword

	createdUser, err := authService.authRepoInterface.Register(user)
	if err != nil {
		return entities.User{}, errors.New("failed to create user")
	}

	return createdUser, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}