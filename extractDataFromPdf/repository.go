package extractDataFromPdf

import "api/models"

type Repository interface {
	SaveIndication(indication *models.Indication) error
}
