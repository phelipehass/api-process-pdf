package service

import (
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

func (s *ExtractService) ProcessData(data *string) {
	indications := strings.Split(*data, "|||")
	var processDatas []ProcessData
	var processData ProcessData

	for _, indication := range indications {
		object := strings.Split(indication, " - ")
		processData.NumberIndication = object[0]
		log.Infof("Número da indicação: %s", object[0])
		processData.NamePersonResponsible = object[1]
		processData.Entourage = object[2]
		processData.Description = object[3]
		processData.District = processDistrict(object[3])
		log.Infof("Bairro: %s", processData.District)
		processData.Street = processStreet(object[3])
		log.Infof("Rua: %s", processData.Street)
		processDatas = append(processDatas, processData)
	}

	//TODO após processamento, salvar no banco
	//fmt.Println(indications)
}

func processDistrict(description string) string {
	district := ""
	start := 0
	searchMatcher := search.New(language.Portuguese, search.IgnoreCase)
	patterns := [2]string{"bairro", "distrito"}

	for _, pattern := range patterns {
		_, start = searchMatcher.IndexString(description, pattern)

		if start != -1 {
			break
		}
	}

	if start != -1 {
		sizeDesc := len(description)
		descriptionTreated := description[start:sizeDesc]
		district = removeAllCharacterRemaining(descriptionTreated, searchMatcher)
	}

	return district
}

func processStreet(description string) string {
	street := ""
	start := 0
	searchMatcher := search.New(language.Portuguese, search.IgnoreCase)
	patterns := [6]string{"rua", "servidão", "estrada", "avenida", "ruas", "av"}

	for _, pattern := range patterns {
		start, _ = searchMatcher.IndexString(description, pattern)

		if start != -1 {
			break
		}
	}

	if start != -1 {
		sizeDesc := len(description)
		descriptionTreated := description[start:sizeDesc]
		removeAllWordAfterDistrict, _ := searchMatcher.IndexString(descriptionTreated, "no bairro")

		if removeAllWordAfterDistrict != -1 {
			descriptionTreated = descriptionTreated[0:removeAllWordAfterDistrict]
		}
		street = removeAllCharacterRemaining(descriptionTreated, searchMatcher)
	}

	return street
}

func removeAllCharacterRemaining(descriptionTreated string, searchMatcher *search.Matcher) string {
	removeAllCharAfterPointer, _ := searchMatcher.IndexString(descriptionTreated, ".")

	if removeAllCharAfterPointer != -1 {
		descriptionTreated = descriptionTreated[0:removeAllCharAfterPointer]
	}

	removeAllCharAfterComma, _ := searchMatcher.IndexString(descriptionTreated, ",")
	if removeAllCharAfterComma != -1 {
		descriptionTreated = descriptionTreated[0:removeAllCharAfterComma]
	}

	return descriptionTreated
}
