package models

import "database/sql"

type Indication struct {
	NumberIndication      string         `db:"number"`
	NamePersonResponsible string         `db:"name_responsible"`
	Entourage             string         `db:"entourage"`
	Street                sql.NullString `db:"street"`
	District              sql.NullString `db:"district"`
	Description           string         `db:"description"`
}
