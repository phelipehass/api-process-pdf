package service

import (
	"api/models"
	"database/sql"
	"github.com/apex/log"
	"golang.org/x/text/language"
	"golang.org/x/text/search"
	"strings"
)

var IGNORE_SPACE = 1

func (s *ExtractService) ProcessData(data *string) []models.Indication {
	indications := strings.Split(*data, "|||")
	var processDatas []models.Indication
	var processData models.Indication

	for _, indication := range indications {
		object := strings.Split(indication, " - ")

		if len(object) > 2 {
			processData.NumberIndication = object[0]
			log.Infof("Número da indicação: %s", object[0])
			processData.NamePersonResponsible = object[1]
			processData.Entourage = object[2]
			processData.Description = strings.TrimSpace(object[3])

			district := processDistrict(object[3])
			if district != "" {
				processData.District = sql.NullString{String: district, Valid: true}
			}

			log.Infof("Bairro: %s", processData.District)

			street := processStreet(object[3])
			if street != "" {
				processData.Street = sql.NullString{String: street, Valid: true}
			}

			log.Infof("Rua: %s", processData.Street)
			processDatas = append(processDatas, processData)
		}
	}

	//TODO após processamento, salvar no banco
	return processDatas
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

	return strings.TrimSpace(district)
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

	return strings.TrimSpace(street)
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
