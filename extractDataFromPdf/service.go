package extractDataFromPdf

import (
	"api/models"
	"mime/multipart"
)

type Service interface {
	ExtractDataFromPDF(fileBuffer *multipart.FileHeader) error
	ProcessData(data *string) []models.Indication
}
