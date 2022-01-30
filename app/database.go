package app

import (
	"database/sql"
	"github.com/SennenHyorou/golang-microservice-building/config"
	"github.com/SennenHyorou/golang-microservice-building/helper"
	"time"
)

func NewDB(config *config.Config) *sql.DB {
	db, err := sql.Open("mysql", config.GetDBConnStr())
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
