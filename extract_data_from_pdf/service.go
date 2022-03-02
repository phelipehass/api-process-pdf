package extract_data_from_pdf

type Service interface {
	ExtractDataFromPDF(pdfPath string) error
}
