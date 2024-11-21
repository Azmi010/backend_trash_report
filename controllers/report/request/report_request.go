package request

import "trash_report/entities"

type ReportRequest struct {
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	// Latitude     string  `json:"latitude" form:"latitude"`
	// Longitude    string  `json:"longitude" form:"longitude"`
	Location     string `json:"location" form:"location"`
	Status       string `json:"status" form:"status"`
	PointsEarned int    `json:"points_earned" form:"points_earned"`
	PhotoUrl     string `json:"photo_url" form:"photo_url"`
	UserID       int    `json:"user_id" form:"user_id"`
}

func (reportRequest ReportRequest) ToEntities() entities.Report {
	return entities.Report{
		Title:        reportRequest.Title,
		Description:  reportRequest.Description,
		Status:       reportRequest.Status,
		PointsEarned: reportRequest.PointsEarned,
		PhotoUrl:     reportRequest.PhotoUrl,
		UserID:       reportRequest.UserID,
	}
}

type UpdateReportStatusRequest struct {
	Status string `json:"status"`
}
