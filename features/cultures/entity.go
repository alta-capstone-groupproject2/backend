package cultures

import (
	"mime/multipart"
	"time"
)

type Core struct {
	ID        int
	Name      string
	Image     string
	City      string
	Details   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CoreReport struct {
	ID        int
	CultureID int
	Message   string
	CreatedAt time.Time
}

type Business interface {
	AddCulture(dataReq Core, fileInfo *multipart.FileHeader, fileData multipart.File) error
	SelectCulture(limit, page int) ([]Core, int64,error)
	SelectCulturebyCultureID(cultureID int) (Core, error)
	UpdateCulture(dataReq Core, cultureID int, fileInfo *multipart.FileHeader, fileData multipart.File) error
	DeleteCulture(cultureID int) error
	

	AddCultureReport(dataReq CoreReport) error
	SelectReport(cultureID int) ([]CoreReport, error)
}

type Data interface {
	AddDataCulture(dataReq Core) error
	SelectDataCulture(limit, offset int) ([]Core, int64, error)
	SelectDataCultureByCultureID(cultureID int) (Core, error)
	UpdateDataCulture(dataReq map[string]interface{}, cultureID int) error
	DeleteDataCulture(cultureID int) error

	AddCultureDataReport(dataReq CoreReport) error
	SelectDataReport(cultureID int) ([]CoreReport, error)
}
