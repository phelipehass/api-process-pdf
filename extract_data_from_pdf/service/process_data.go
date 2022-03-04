package service

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/search"
	"strings"
)

var IGNORE_SPACE = 1

type ProcessData struct {
	NumberIndication      string
	NamePersonResponsible string
	Entourage             string
	Street                string
	District              string
	Description           string
}

func (s *ExtractService) ProcessData(data string) {
	indications := strings.Split(data, "|")
	var processDatas []ProcessData
	var processData ProcessData

	for _, indication := range indications {
		object := strings.Split(indication, " - ")
		processData.NumberIndication = object[0]
		processData.NamePersonResponsible = object[1]
		processData.Entourage = object[2]
		processData.Description = object[3]
		processData.District = processStreet(object[3])
		//TODO criar process da rua para extrair
		processDatas = append(processDatas, processData)
	}

	//TODO após processamento, salvar no banco
	fmt.Println(indications)
}

func processStreet(description string) string {
	district := ""
	searchMatcher := search.New(language.Portuguese, search.IgnoreCase)
	_, start := searchMatcher.IndexString(description, "bairro")

	if start != -1 {
		removeAllCharAfterPointer, _ := searchMatcher.IndexString(description, ".")
		if removeAllCharAfterPointer == -1 {
			removeAllCharAfterPointer = len(description)
		}

		start += IGNORE_SPACE
		district = description[start:removeAllCharAfterPointer]
	}

	return district
}
