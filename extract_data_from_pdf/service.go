package extract_data_from_pdf

import (
	"api/extract_data_from_pdf/service"
	"mime/multipart"
)

type Service interface {
	ExtractDataFromPDF(fileBuffer *multipart.FileHeader) error
	ProcessData(data *string) []service.ProcessData
}
