package repo

import (
	"trash_report/entities"
	"trash_report/repo/record"

	"gorm.io/gorm"
)

type ReportRepo struct {
	db *gorm.DB
}

func NewReportRepo(db *gorm.DB) *ReportRepo {
	return &ReportRepo{db: db}
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