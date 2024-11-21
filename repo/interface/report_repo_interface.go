package repoInterface

import "trash_report/entities"

type ReportRepository interface {
	CreateReport(report *entities.Report) error
	GetReportsByUser(userID uint) ([]entities.Report, error)
	UpdateReportByUser(report *entities.Report) error
	GetAllReports() ([]entities.Report, error)
	UpdateReportStatus(reportID uint, status string) error
	DeleteReportByAdmin(reportID uint) error
	AddReportAnalysis(reportID int) (*entities.ReportAnalysis, error)
}