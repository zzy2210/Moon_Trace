package model

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SubDomain struct {
	gorm.Model
	Domain    string
	SubDomain string
}

type Port struct {
	gorm.Model
	Host string
	Port string
}

type Urls struct {
	gorm.Model
	Host string
	Url  string
}
type PostgresDB struct {
}

func NewDB(host, port, user, pwd, dbname string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, pwd, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
