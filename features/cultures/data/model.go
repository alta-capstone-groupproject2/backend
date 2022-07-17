package data

import (
	"lami/app/features/cultures"

	"gorm.io/gorm"
)

type Culture struct {
	gorm.Model
	Name    string `json:"name"`
	Details string `json:"details"`
	Image   string `json:"image"`
	City    string `json:"city"`
}

type Report struct {
	gorm.Model
	CultureID int    `json:"culture_id"`
	Message   string `json:"message"`
	Culture   Culture
}

func fromCore(core cultures.Core) Culture {
	return Culture{
		Name:    core.Name,
		Details: core.Details,
		Image:   core.Image,
		City:    core.City,
	}
}

func fromCoreReport(core cultures.CoreReport) Report {
	return Report{
		CultureID: core.CultureID,
		Message:   core.Message,
		
	}
}

func (data *Culture) toCore() cultures.Core {
	return cultures.Core{
		ID:      int(data.ID),
		Name:    data.Name,
		Image:   data.Image,
		City:    data.City,
		Details: data.Details,
	}
}

func ToCore(data Culture) cultures.Core {
	return data.toCore()
}

func ToCoreList(data []Culture) []cultures.Core {
	res := []cultures.Core{}
	for key := range data {
		res = append(res, data[key].toCore())
	}
	return res
}

func (data *Report) toCoreReport() cultures.CoreReport {
	return cultures.CoreReport{
		ID:      int(data.ID),
		Message: data.Message,
		CreatedAt: data.CreatedAt,
	}
}

func ToCoreReportList(data []Report) []cultures.CoreReport {
	res := []cultures.CoreReport{}
	for key := range data {
		res = append(res, data[key].toCoreReport())
	}
	return res
}



