package model

import (
	"Moon_Trace/eng/conf"
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

func NewGormDB(conf *conf.Conf) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", conf.PGConf.Host, conf.PGConf.User, conf.PGConf.PassWord, conf.PGConf.DBName, conf.PGConf.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
