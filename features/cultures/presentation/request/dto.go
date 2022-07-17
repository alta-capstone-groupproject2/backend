package request

import (
	"lami/app/features/cultures"
)

type Culture struct {
	Name    string `json:"name" form:"name"`
	Details string `json:"details" form:"details"`
	City    string `json:"city" form:"city"`
	Image   string
}

type Report struct {
	CultureID int
	Message   string `json:"message" form:"message"`
}

//	For insert culture
func ToCore(cultureReq Culture) cultures.Core {
	return cultures.Core{
		Name:    cultureReq.Name,
		Image:   cultureReq.Image,
		City:    cultureReq.City,
		Details: cultureReq.Details,
	}
}

// For update culture
func ToCoreUpdate(cultureReq Culture) cultures.Core {
	return cultures.Core{
		ID:      0,
		Name:    cultureReq.Name,
		Image:   cultureReq.Image,
		City:    cultureReq.City,
		Details: cultureReq.Details,
	}
}

//	For insert report
func ToCoreReport(reportReq Report) cultures.CoreReport {
	return cultures.CoreReport{
		Message:   reportReq.Message,
		CultureID: reportReq.CultureID,
	}
}
