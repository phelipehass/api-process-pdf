package service

import (
	"api/config"
	"api/extract_data_from_pdf/repository"
	"fmt"
	"github.com/apex/log"
	"github.com/ledongthuc/pdf"
	"mime/multipart"
	"regexp"
	"strings"
)

type ExtractService struct {
	InitExtraction       string
	FinishExtraction     string
	Title                string
	Subtitle             string
	IndicationRepository *repository.PostgresRepository
}

func NewService(indicationRepo *repository.PostgresRepository) *ExtractService {
	return &ExtractService{
		InitExtraction:       config.GetInitExtraction(),
		FinishExtraction:     config.GetFinishExtraction(),
		Title:                config.GetTitle(),
		Subtitle:             config.GetSubtitle(),
		IndicationRepository: indicationRepo,
	}
}

func (s *ExtractService) ExtractDataFromPDF(fileBuffer *multipart.FileHeader) error {
	defer func() {
		if recovered := recover(); recovered != nil {
			fmt.Println("[ExtractDataFromPDF] - Aconteceu um erro inesperado")
		}
	}()

	file, err := fileBuffer.Open()
	if err != nil {
		return err
	}

	pdfReader, err := pdf.NewReader(file, fileBuffer.Size)
	if err != nil {
		return err
	}

	totalPages := pdfReader.NumPage()
	data := s.processFilePerPage(pdfReader, &totalPages)
	file.Close()

	indications := s.ProcessData(data)

	if len(indications) > 0 {
		for _, indication := range indications {
			errSave := s.IndicationRepository.SaveIndication(&indication)
			if errSave != nil {
				log.Errorf("[ExtractDataFromPDF] - Erro ao salvar indicação no banco - %s", errSave)
			}
		}
	}

	return nil
}

func (s *ExtractService) processFilePerPage(pdfReader *pdf.Reader, totalPages *int) *string {
	data := ""
	pageIndex := 1
	processTextUntilNextTitle := false

	for pageIndex = 1; pageIndex <= *totalPages; pageIndex++ {
		var err error = nil
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
			log.Error(err.Error())
			continue
		}

		isBreak, pageData := s.processTextByRow(&rows, &pageWithIndications, &pageIndex)
		if pageIndex != pageWithIndications {
			data += pageData
		} else {
			data += strings.Replace(pageData, " INDICAÇÕES |||", "", 1)
		}

		if isBreak {
			break
		}
	}

	return &data
}

func (s *ExtractService) processTextByRow(rows *pdf.Rows, pageWithIndications *int, pageIndex *int) (bool, string) {
	markRow := -1
	isBreak := false
	textByRow := ""

	for valueRow, row := range *rows {
		for _, word := range row.Content {
			if strings.Contains(word.S, s.FinishExtraction) {
				isBreak = true
				break
			}

			//tratativa
			re := regexp.MustCompile(`N[^a-zA-Z0-9]+`)
			loc := re.FindStringIndex(word.S)
			if loc != nil && loc[0] >= 0 && loc[0] < 10 {
				word.S = replaceFirstDegreeSymbol("F" + word.S)
			}

			if *pageWithIndications == *pageIndex && strings.Contains(word.S, s.InitExtraction) {
				markRow = valueRow
				textByRow += " " + word.S
			} else if *pageWithIndications == *pageIndex && markRow != -1 && valueRow > markRow {
				textByRow += " " + word.S
			} else if *pageIndex > *pageWithIndications && !strings.Contains(word.S, s.Title) && !strings.Contains(word.S, s.Subtitle) {
				textByRow += " " + word.S
			}
		}

		if isBreak {
			break
		}
	}

	return isBreak, textByRow
}

func replaceFirstDegreeSymbol(line string) string {
	re := regexp.MustCompile(`FN[^a-zA-Z0-9]+`)
	subst := "|||"

	return re.ReplaceAllString(line, subst)
}
