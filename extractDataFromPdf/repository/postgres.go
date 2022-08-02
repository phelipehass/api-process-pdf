package repository

import (
	"api/models"
	"database/sql"
	"gopkg.in/gorp.v2"
)

type PostgresRepository struct {
	dbmap *gorp.DbMap
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(models.Indication{}, "indication")

	dbmap.TraceOn("[gorp]", &models.DBLog{})

	return &PostgresRepository{
		dbmap: dbmap,
	}
}

func (p *PostgresRepository) SaveIndication(indication *models.Indication) error {
	err := p.dbmap.Insert(indication)
	if err != nil {
		return err
	}

	return nil
}
