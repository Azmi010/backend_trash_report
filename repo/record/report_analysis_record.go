package record

import "trash_report/entities"

type ReportAnalysis struct {
	ID       int `gorm:"primaryKey"`
	Status   string
	Analysis string
	ReportID int `gorm:"index"`
}

type GeminiResponse struct {
	Status string `json:"status"`
	Data   struct {
		Analysis string `json:"analysis"`
	} `json:"data"`
}

func (ra *ReportAnalysis) ToEntity() entities.ReportAnalysis {
	return entities.ReportAnalysis{
		ID:       ra.ID,
		Status:   ra.Status,
		Analysis: ra.Analysis,
		ReportID: ra.ReportID,
	}
}

func ReportAnalysisFromEntity(analysis entities.ReportAnalysis) ReportAnalysis {
	return ReportAnalysis{
		ID:       analysis.ID,
		Status:   analysis.Status,
		Analysis: analysis.Analysis,
		ReportID: analysis.ReportID,
	}
}
