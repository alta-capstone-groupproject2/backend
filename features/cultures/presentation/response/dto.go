package response

import (
	"lami/app/features/cultures"
)

type Culture struct {
	ID      int    `json:"cultureID"`
	Image   string `json:"image"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Details string `json:"details"`
}

type CultureReport struct {
	ID        int       `json:"reportID"`
	Message   string    `json:"message"`
	CreatedAt string `json:"createdAt"`
}

func FromCore(core cultures.Core) Culture {
	return Culture{
		ID:      core.ID,
		Image:   core.Image,
		Name:    core.Name,
		City:    core.City,
		Details: core.Details,
	}
}

func FromCoreReport(core cultures.CoreReport) CultureReport {
	return CultureReport{
		ID:        core.ID,
		Message:   core.Message,
		CreatedAt: core.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func FromCoreListReport(data []cultures.CoreReport) []CultureReport {
	res := []CultureReport{}
	for key := range data {
		res = append(res, FromCoreReport(data[key]))
	}
	return res
}

func FromCoreList(data []cultures.Core) []Culture {
	res := []Culture{}
	for key := range data {
		res = append(res, FromCore(data[key]))
	}
	return res
}
