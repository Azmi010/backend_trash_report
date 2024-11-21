package routes

import (
	"os"
	"trash_report/controllers/auth"
	"trash_report/controllers/report"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthController *auth.AuthController
	ReportController *report.ReportController
}

func (rc RouteController) InitRoute(e *echo.Echo) {
	e.POST("/login", rc.AuthController.LoginController)
	e.POST("/register", rc.AuthController.RegisterController)
	eJWT := e.Group("")
	eJWT.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))
	eUser := eJWT.Group("/users")
	eUser.POST("/reports", rc.ReportController.CreateReport)
	eUser.GET("/reports/:user_id", rc.ReportController.GetReportsByUser)
	eUser.PUT("/reports/:report_id", rc.ReportController.UpdateReport)
	eAdmin := eJWT.Group("/admin")
	eAdmin.GET("/reports", rc.ReportController.GetAllReports)
	eAdmin.PUT("/reports/:report_id", rc.ReportController.UpdateReportStatus)
	eAdmin.DELETE("/reports/:report_id", rc.ReportController.DeleteReportByAdmin)
	eAdmin.POST("/reports/:report_id/analyze", rc.ReportController.AddReportAnalysis)
}