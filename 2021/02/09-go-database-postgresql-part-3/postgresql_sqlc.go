package main

import (
	"context"
	"database/sql"
	"os"
)

//go:generate sqlc generate

type PostgreSQLC struct {
	db *sql.DB
}

func NewPostgreSQLC() (*PostgreSQLC, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return &PostgreSQLC{
		db: db,
	}, nil
}

func (p *PostgreSQLC) Close() {
	p.db.Close()
}

func (p *PostgreSQLC) FindByNConst(nconst string) (Name, error) {
	row, err := New(p.db).SelectName(context.Background(), sql.NullString{String: nconst})

	if err != nil {
		return Name{}, err
	}

	return Name{
		NConst:    row.Nconst.String,
		Name:      row.PrimaryName.String,
		BirthYear: row.BirthYear.String,
		DeathYear: row.DeathYear.String,
	}, nil
}
