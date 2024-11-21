package record

import (
	"trash_report/entities"
)

type Report struct {
	ID           int `gorm:"primaryKey"`
	Title        string
	Description  string
	Location     string
	Status       string
	PointsEarned int
	PhotoUrl     string
	UserID       int `json:"user_id" gorm:"index"`
}

func (r *Report) ToEntity() entities.Report {
	return entities.Report{
		ID:           r.ID,
		Title:        r.Title,
		Description:  r.Description,
		Location:     r.Location,
		Status:       r.Status,
		PointsEarned: r.PointsEarned,
		PhotoUrl:     r.PhotoUrl,
		UserID:       r.UserID,
	}
}

func FromEntity(report entities.Report) Report {
	return Report{
		ID:           report.ID,
		Title:        report.Title,
		Description:  report.Description,
		Location:     report.Location,
		Status:       report.Status,
		PointsEarned: report.PointsEarned,
		PhotoUrl:     report.PhotoUrl,
		UserID:       report.UserID,
	}
}
