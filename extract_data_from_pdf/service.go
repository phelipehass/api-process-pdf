package extract_data_from_pdf

import "mime/multipart"

type Service interface {
	ExtractDataFromPDF(fileBuffer *multipart.FileHeader) error
}
