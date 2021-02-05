package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgreSQLpgx struct {
	pool *pgxpool.Pool
}

func NewPostgreSQLpgx() (*PostgreSQLpgx, error) {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return &PostgreSQLpgx{
		pool: pool,
	}, nil
}

func (p *PostgreSQLpgx) Close() {
	p.pool.Close()
}

func (p *PostgreSQLpgx) FindByNConst(nconst string) (Name, error) {
	query := `SELECT nconst, primary_name, birth_year, death_year FROM "names" WHERE nconst = $1`

	var res Name

	if err := p.pool.QueryRow(context.Background(), query, nconst).
		Scan(&res.NConst, &res.Name, &res.BirthYear, &res.DeathYear); err != nil {
		return Name{}, err
	}

	return res, nil
}
