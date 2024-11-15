package service

import (
	"mime/multipart"
	"trash_report/config"
	"trash_report/entities"
	"trash_report/helper"
	"trash_report/repo/interface"
)
type reportService struct {
	reportRepo repoInterface.ReportRepository
}

func NewReportService(reportRepo repoInterface.ReportRepository) *reportService {
	return &reportService{reportRepo}
}

func (s *reportService) CreateReport(report *entities.Report, file *multipart.FileHeader) error {
	bucket := config.GetStorageBucket()
	photoURL, err := helper.UploadFileToFirebase(file, bucket)
	if err != nil {
		return err
	}
	report.PhotoUrl = photoURL
	return s.reportRepo.CreateReport(report)
}

func (s *reportService) GetReportsByUser(userID uint) ([]entities.Report, error) {
	return s.reportRepo.GetReportsByUser(userID)
}

func (s *reportService) UpdateReportByUser(report *entities.Report) error {
	return s.reportRepo.UpdateReportByUser(report)
}

func (s *reportService) GetAllReports() ([]entities.Report, error) {
	return s.reportRepo.GetAllReports()
}

func (s *reportService) UpdateReportStatus(reportID uint, status string) error {
	return s.reportRepo.UpdateReportStatus(reportID, status)
}

func (s *reportService) DeleteReportByAdmin(reportID uint) error {
	return s.reportRepo.DeleteReportByAdmin(reportID)
}