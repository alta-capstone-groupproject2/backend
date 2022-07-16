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
	UpdateCulture(dataReq Core, cultureID int, fileInfo *multipart.FileHeader, fileData multipart.File) error
	DeleteCulture(cultureID, idUser int) error
	// SelectCultureList() error
	SelectCulturebyCultureID(cultureID int) (Core, error)
	SelectMyCulture(idUser int) ([]Core, error)

	AddCultureReport(dataReq CoreReport) error
	SelectReport(cultureID int) ([]CoreReport, error)
}

type Data interface {
	AddDataCulture(dataReq Core) error
	UpdateDataCulture(dataReq Core, cultureID int) error
	DeleteDataCulture(cultureID, idUser int) error
	// SelectDataCultureList() error
	SelectDataCultureByCultureID(cultureID int) (Core, error)
	SelectDataMyCulture(idUser int) ([]Core, error)

	AddCultureDataReport(dataReq CoreReport) error
	SelectDataReport(cultureID int) ([]CoreReport, error)
}
