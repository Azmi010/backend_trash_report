package repo

import (
	"fmt"
	"trash_report/entities"
	"trash_report/helper"
	"trash_report/repo/record"

	"gorm.io/gorm"
)

type ReportRepo struct {
	db *gorm.DB
	geminiHelper *helper.GeminiHelper
}

func NewReportRepo(db *gorm.DB, geminiHelper *helper.GeminiHelper) *ReportRepo {
	return &ReportRepo{
		db: db,
		geminiHelper: geminiHelper,
	}
}

func (r *ReportRepo) CreateReport(report *entities.Report) error {
	reports := record.FromEntity(*report)
	if err := r.db.Create(&reports).Error; err != nil {
		return err
	}

	report.ID = reports.ID
	return nil
}

func (r *ReportRepo) GetReportsByUser(userID uint) ([]entities.Report, error) {
	var reports []entities.Report
	if err := r.db.Where("user_id = ?", userID).Find(&reports).Error; err != nil {
		return nil, err
	}
	return reports, nil
}

func (r *ReportRepo) UpdateReportByUser(report *entities.Report) error {
	return r.db.Model(&entities.Report{}).
		Where("id = ? AND user_id = ?", report.ID, report.UserID).
		Updates(map[string]interface{}{
			"title":         report.Title,
			"description":   report.Description,
			"location":      report.Location,
			"points_earned": report.PointsEarned,
			"photo_url":     report.PhotoUrl,
		}).Error
}

func (r *ReportRepo) GetAllReports() ([]entities.Report, error) {
	var reportRecords []record.Report
	if err := r.db.Find(&reportRecords).Error; err != nil {
		return nil, err
	}

	var reports []entities.Report
	for _, reportRecord := range reportRecords {
		reports = append(reports, reportRecord.ToEntity())
	}

	return reports, nil
}

func (r *ReportRepo) UpdateReportStatus(reportID uint, status string) error {
	return r.db.Model(&record.Report{}).Where("id = ?", reportID).Update("status", status).Error
}

func (r *ReportRepo) DeleteReportByAdmin(reportID uint) error {
	return r.db.Where("id = ?", reportID).Delete(&entities.Report{}).Error
}

func (r *ReportRepo) AddReportAnalysis(reportID int) (*entities.ReportAnalysis, error) {
	var report record.Report
	if err := r.db.First(&report, "id = ?", reportID).Error; err != nil {
		return nil, err
	}

	apiResponse, err := r.geminiHelper.AnalyzeReport(report.Title, report.Description, report.PhotoUrl)
	if err != nil {
		return nil, err
	}

	reportAnalysis := &entities.ReportAnalysis{
		Status:   "Analyzed",
		Analysis: fmt.Sprintf("Title: %s", apiResponse.Data.Analysis),
		ReportID: reportID,
	}

	analysisRecord := record.ReportAnalysisFromEntity(*reportAnalysis)
	if err := r.db.Create(&analysisRecord).Error; err != nil {
		return nil, err
	}

	reportAnalysis.ID = analysisRecord.ID
	return reportAnalysis, nil
}