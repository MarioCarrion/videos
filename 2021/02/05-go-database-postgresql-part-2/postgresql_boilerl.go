package main

import (
	"context"
	"database/sql"
	"os"

	"github.com/MarioCarrion/videos/2021/02/05-go-database-postgresql-part-2/models"
)

//go:generate sqlboiler --wipe --no-tests psql

type PostgreSQLboiler struct {
	db *sql.DB
}

func NewPostgreSQLboiler() (*PostgreSQLboiler, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return &PostgreSQLboiler{
		db: db,
	}, nil
}

func (p *PostgreSQLboiler) Close() {
	p.db.Close()
}

func (p *PostgreSQLboiler) FindByNConst(nconst string) (Name, error) {
	result, err := models.FindName(context.Background(), p.db, nconst)
	if err != nil {
		return Name{}, err
	}

	return Name{
		NConst:    result.Nconst,
		Name:      result.PrimaryName.String,
		BirthYear: result.BirthYear.String,
		DeathYear: result.DeathYear.String,
	}, nil
}
