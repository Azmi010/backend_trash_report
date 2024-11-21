package main

import (
	"log"
	"os"
	"trash_report/config"
	"trash_report/controllers/auth"
	"trash_report/controllers/report"
	"trash_report/helper"
	"trash_report/middleware"
	"trash_report/repo/repo"
	"trash_report/routes"
	"trash_report/services/service"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()
	db, _ := config.ConnectDatabase()
	config.MigrateDB(db)
	config.InitializeFirebase()
	geminiHelper := helper.NewGeminiHelper(os.Getenv("GEMINI_API_KEY"))
	e := echo.New()

	authJwt := middleware.JwtAlta{}

	authRepo := repo.NewAuthRepo(db)
	authService := service.NewAuthService(authRepo, authJwt)
	authController := auth.NewAuthController(authService)
	reportRepo := repo.NewReportRepo(db, geminiHelper)
	reportService := service.NewReportService(reportRepo)
	reportController := report.NewReportController(reportService)

	routeController := routes.RouteController{
		AuthController:   authController,
		ReportController: reportController,
	}
	routeController.InitRoute(e)

	e.Start(":8000")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic("failed lod env")
	}
}