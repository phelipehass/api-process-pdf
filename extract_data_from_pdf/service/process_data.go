package service

import (
	"fmt"
	"github.com/apex/log"
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
		log.Infof("número da indicação: %s", object[0])
		processData.NamePersonResponsible = object[1]
		processData.Entourage = object[2]
		processData.Description = object[3]
		processData.District = processDistrict(object[3])
		processData.Street = processStreet(object[3])
		processDatas = append(processDatas, processData)
	}

	//TODO após processamento, salvar no banco
	fmt.Println(indications)
}

func processDistrict(description string) string {
	district := ""
	searchMatcher := search.New(language.Portuguese, search.IgnoreCase)
	_, start := searchMatcher.IndexString(description, "bairro")

	if start != -1 {
		sizeDesc := len(description)
		descriptionTreated := description[start:sizeDesc]
		removeAllCharAfterPointer, _ := searchMatcher.IndexString(descriptionTreated, ".")
		if removeAllCharAfterPointer != -1 {
			district = description[0:removeAllCharAfterPointer]
		}
	}

	return district
}

func processStreet(description string) string {
	street := ""
	start := 0
	searchMatcher := search.New(language.Portuguese, search.IgnoreCase)
	patterns := [4]string{"rua", "servidão", "estrada", "avenida"}

	for _, pattern := range patterns {
		start, _ = searchMatcher.IndexString(description, pattern)

		if start != -1 {
			break
		}
	}

	if start != -1 {
		sizeDesc := len(description)
		descriptionTreated := description[start:sizeDesc]
		removeAllCharAfterComma, _ := searchMatcher.IndexString(descriptionTreated, ",")
		if removeAllCharAfterComma != -1 {
			street = descriptionTreated[0:removeAllCharAfterComma]
		}
	}

	return street
}
