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
		processData.District = processStreet(object[3])
		//TODO criar process da rua para extrair
		processDatas = append(processDatas, processData)
	}

	fmt.Println(indications)
}

func processStreet(description string) string {
	district := ""
	searchMatcher := search.New(language.Portuguese, search.IgnoreCase)
	_, start := searchMatcher.IndexString(description, "bairro")
	if start != -1 {
		start += IGNORE_SPACE
		finish := len(description) - 1
		district = description[start:finish]
	}

	return district
}
