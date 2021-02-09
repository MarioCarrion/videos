package main

import (
	"context"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // pgx also supported
)

type PostgreSQLsqlx struct {
	db *sqlx.DB
}

func NewPostgreSQLsqlx() (*PostgreSQLsqlx, error) {
	db, err := sqlx.ConnectContext(context.Background(), "postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return &PostgreSQLsqlx{
		db: db,
	}, nil
}

func (p *PostgreSQLsqlx) Close() {
	p.db.Close()
}

func (p *PostgreSQLsqlx) FindByNConst(nconst string) (Name, error) {
	query := `SELECT nconst, primary_name, birth_year, death_year FROM "names" WHERE nconst = $1`

	var result struct {
		NConst    string `db:"nconst"`
		Name      string `db:"primary_name"`
		BirthYear string `db:"birth_year"`
		DeathYear string `db:"death_year"`
	}

	if err := p.db.QueryRowx(query, nconst).StructScan(&result); err != nil {
		return Name{}, err
	}

	return Name{
		NConst:    result.NConst,
		Name:      result.Name,
		BirthYear: result.BirthYear,
		DeathYear: result.DeathYear,
	}, nil
}
