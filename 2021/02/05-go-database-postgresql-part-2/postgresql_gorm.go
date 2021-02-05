package main

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

type PostgreSQLgorm struct {
	db *gorm.DB
}

func NewPostgreSQLgorm() (*PostgreSQLgorm, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &PostgreSQLgorm{
		db: db,
	}, nil
}

func (p *PostgreSQLgorm) FindByNConst(nconst string) (Name, error) {
	var result gormNames

	tx := p.db.Where("nconst = ?", nconst).First(&result)
	if tx.Error != nil {
		return Name{}, tx.Error
	}

	return Name{
		NConst:    result.NConst,
		Name:      result.Name,
		BirthYear: result.BirthYear,
		DeathYear: result.DeathYear,
	}, nil
}

type gormNames struct {
	NConst    string `gorm:"primaryKey;column:nconst"`
	Name      string `gorm:"column:primary_name"`
	BirthYear string
	DeathYear string
}

func (gormNames) TableName() string {
	return "names"
}
