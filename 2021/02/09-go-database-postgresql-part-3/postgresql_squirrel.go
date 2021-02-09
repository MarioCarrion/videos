package main

import (
	"context"
	"database/sql"
	"os"

	sq "github.com/Masterminds/squirrel"
)

type PostgreSQLsquirrel struct {
	db *sql.DB
}

func NewPostgreSQLsquirrel() (*PostgreSQLsquirrel, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return &PostgreSQLsquirrel{
		db: db,
	}, nil
}

func (p *PostgreSQLsquirrel) Close() {
	p.db.Close()
}

func (p *PostgreSQLsquirrel) FindByNConst(nconst string) (Name, error) {
	var res Name

	query := sq.
		Select("nconst", "primary_name", "birth_year", "death_year").
		From("names").
		Where("nconst = ?", nconst).
		PlaceholderFormat(sq.Dollar).
		RunWith(p.db)

	if err := query.
		ScanContext(context.Background(), &res.NConst, &res.Name, &res.BirthYear, &res.DeathYear); err != nil {
		return Name{}, err
	}

	return res, nil
}
