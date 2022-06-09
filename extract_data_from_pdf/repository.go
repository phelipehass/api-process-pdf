package extract_data_from_pdf

import "api/models"

type Repository interface {
	SaveIndication(indication *models.Indication) error
}
