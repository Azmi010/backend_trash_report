package auth

import (
	"trash_report/controllers/auth/request"
	"trash_report/controllers/auth/response"
	"trash_report/controllers/base"
	serviceInterface "trash_report/services/interface"

	"github.com/labstack/echo/v4"
)

func NewAuthController(as serviceInterface.AuthInterface) *AuthController {
	return &AuthController{
		authService: as,
	}
}

type AuthController struct {
	authService serviceInterface.AuthInterface
}

func (userController AuthController) LoginController(c echo.Context) error {
	userLogin := request.LoginRequest{}
	c.Bind(&userLogin)
	user, err := userController.authService.Login(userLogin.ToEntities())
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, response.FromLoginEntities(user))
}

func (authController AuthController) RegisterController(c echo.Context) error {
	userRegister := request.RegisterRequest{}
	if err := c.Bind(&userRegister); err != nil {
		return base.ErrorResponse(c, err)
	}

	user, err := authController.authService.Register(userRegister.ToEntities())
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, response.FromRegisterEntities(user))
}