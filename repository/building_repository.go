package repository

import (
	"context"
	"database/sql"
	"golang-microservice-building/model/domain"
)

type BuildingRepository interface {
	Store(ctx context.Context, tx *sql.Tx, building domain.Building) domain.Building
	Update(ctx context.Context, tx *sql.Tx, building domain.Building) domain.Building
	Delete(ctx context.Context, tx *sql.Tx, building domain.Building)
	FindById(ctx context.Context, tx *sql.Tx, buildingId int) (domain.Building, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Building
}
