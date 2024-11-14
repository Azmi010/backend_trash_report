package base

import (
	"net/http"
	"trash_report/helper"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessMultiResponse(c echo.Context, data interface{}, message string) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"message": message,
		"data":   data,
	})
}

func SuccessResponse(c echo.Context, data any) error {
	return c.JSON(http.StatusOK, BaseResponse{
		Status:  true,
		Message: "success",
		Data:    data,
	})
}

func ErrorResponse(c echo.Context, err error) error {
	return c.JSON(helper.GetResponseCodeFromErr(err), BaseResponse{
		Status:  false,
		Message: err.Error(),
	})
}
