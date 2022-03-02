package service

import (
	"api/extract_data_from_pdf/config"
	"fmt"
	"github.com/ledongthuc/pdf"
	"strings"
)

type ExtractService struct {
	InitExtraction   string
	FinishExtraction string
	Title            string
	Subtitle         string
}

func NewService() *ExtractService {
	return &ExtractService{
		InitExtraction:   config.GetInitExtraction(),
		FinishExtraction: config.GetFinishExtraction(),
		Title:            config.GetTitle(),
		Subtitle:         config.GetSubtitle(),
	}
}

func (s *ExtractService) ExtractDataFromPDF(pdfPath string) error {
	file, pdfReader, err := pdf.Open(pdfPath)
	defer func() {
		if recovered := recover(); recovered != nil {
			fmt.Println("[ExtractDataFromPDF] - Aconteceu um erro inesperado")
		}
	}()

	defer file.Close()
	if err != nil {
		return err
	}

	totalPages := pdfReader.NumPage()
	processTextUntilNextTitle := false

	for pageIndex := 1; pageIndex <= totalPages; pageIndex++ {
		err = nil
		var pageWithIndications int
		pageFile := pdfReader.Page(pageIndex)

		if pageFile.V.IsNull() {
			continue
		}

		text, err := pageFile.GetPlainText(nil)
		if err != nil {
			continue
		}

		if strings.Contains(text, s.InitExtraction) {
			pageWithIndications = pageIndex
			processTextUntilNextTitle = true
		} else if !processTextUntilNextTitle {
			continue
		}

		text = ""
		rows, err := pageFile.GetTextByRow()
		if err != nil {
			continue
		}

		isBreak := s.processTextByRow(rows, pageWithIndications, pageIndex)
		if isBreak {
			break
		}

	}

	return nil
}

func (s *ExtractService) processTextByRow(rows pdf.Rows, pageWithIndications int, pageIndex int) bool {
	//TODO melhorar processamento, pois só está logando dado extraído
	markRow := -1
	isBreak := false
	textByRow := ""

	for valueRow, row := range rows {
		for _, word := range row.Content {
			if strings.Contains(word.S, s.FinishExtraction) {
				isBreak = true
				break
			}

			if pageWithIndications == pageIndex && strings.Contains(word.S, s.InitExtraction) {
				markRow = valueRow
				textByRow += word.S
			} else if pageWithIndications == pageIndex && markRow != -1 && valueRow > markRow {
				textByRow += word.S
			} else if pageIndex > pageWithIndications && !strings.Contains(word.S, s.Title) && !strings.Contains(word.S, s.Subtitle) {
				textByRow += word.S
			}
		}

		if len(textByRow) > 0 {
			fmt.Println(textByRow)
			textByRow = ""
		}

		if isBreak {
			break
		}
	}

	return isBreak
}
