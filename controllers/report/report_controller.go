package report

import (
	"fmt"
	"strconv"
	"trash_report/controllers/base"
	"trash_report/controllers/report/request"
	"trash_report/helper"
	"trash_report/services/interface"

	"github.com/labstack/echo/v4"
)

type ReportController struct {
	reportService serviceInterface.ReportService
}

func NewReportController(reportService serviceInterface.ReportService) *ReportController {
	return &ReportController{reportService}
}

func (rc *ReportController) CreateReport(c echo.Context) error {
	var reportReq request.ReportRequest
	if err := c.Bind(&reportReq); err != nil {
		return base.ErrorResponse(c, err)
	}

	location, err := helper.GetAddressFromCoordinates(reportReq.Latitude, reportReq.Longitude)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	file, err := c.FormFile("photo_url")
	if err != nil {
		return base.ErrorResponse(c, fmt.Errorf("photo is required"))
	}

	report := reportReq.ToEntities()
	report.Location = location
	
	if err := rc.reportService.CreateReport(&report, file); err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessMultiResponse(c, report, "Report Created Successfully")
}

func (rc *ReportController) GetReportsByUser(c echo.Context) error {
	userIDStr := c.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	reports, err := rc.reportService.GetReportsByUser(uint(userID))
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessMultiResponse(c, reports, "Successfully Get Report Data")
}

func (rc *ReportController) UpdateReport(c echo.Context) error {
	reportIDStr := c.Param("report_id")
	reportID, err := strconv.Atoi(reportIDStr)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	var reportRequest request.ReportRequest
	if err := c.Bind(&reportRequest); err != nil {
		return base.ErrorResponse(c, err)
	}

	report := reportRequest.ToEntities()
	report.ID = reportID

	if err := rc.reportService.UpdateReportByUser(&report); err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, report)
}

func (rc *ReportController) GetAllReports(c echo.Context) error {
	reports, err := rc.reportService.GetAllReports()
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessMultiResponse(c, reports, "Successfully Get Report Data")
}

func (rc *ReportController) UpdateReportStatus(c echo.Context) error {
	reportID, err := strconv.Atoi(c.Param("report_id"))
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	var reportReq request.UpdateReportStatusRequest
	if err := c.Bind(&reportReq); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := rc.reportService.UpdateReportStatus(uint(reportID), reportReq.Status); err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessMultiResponse(c, "Report status updated successfully", "Update Success")
}

func (rc *ReportController) DeleteReportByAdmin(c echo.Context) error {
	reportIDStr := c.Param("report_id")
	reportID, err := strconv.Atoi(reportIDStr)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	err = rc.reportService.DeleteReportByAdmin(uint(reportID))
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, "Report deleted by admin successfully")
}