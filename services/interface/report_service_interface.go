package serviceInterface

import (
	"mime/multipart"
	"trash_report/entities"
)

type ReportService interface {
	CreateReport(report *entities.Report, file *multipart.FileHeader) error
	GetReportsByUser(userID uint) ([]entities.Report, error)
	UpdateReportByUser(report *entities.Report) error
	GetAllReports() ([]entities.Report, error)
	UpdateReportStatus(reportID uint, status string) error
	DeleteReportByAdmin(reportID uint) error
}