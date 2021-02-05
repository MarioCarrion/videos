package main

import (
	"context"
	"database/sql"
	"os"

	_ "github.com/lib/pq" // pgx also supported
)

type PostgreSQLsql struct {
	pool *sql.DB
}

func NewPostgreSQLsql() (*PostgreSQLsql, error) {
	pool, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(); err != nil {
		return nil, err
	}

	return &PostgreSQLsql{
		pool: pool,
	}, nil
}

func (p *PostgreSQLsql) Close() {
	p.pool.Close()
}

func (p *PostgreSQLsql) FindByNConst(nconst string) (Name, error) {
	query := `SELECT nconst, primary_name, birth_year, death_year FROM "names" WHERE nconst = $1`

	var res Name

	if err := p.pool.QueryRowContext(context.Background(), query, nconst).
		Scan(&res.NConst, &res.Name, &res.BirthYear, &res.DeathYear); err != nil {
		return Name{}, err
	}

	return res, nil
}
