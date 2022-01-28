package app

import (
	"database/sql"
	"github.com/SennenHyorou/golang-microservice-building/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:admin123@tcp(localhost:3306)/microservice_building?tls=skip-verify&autocommit=true")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
