package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/SennenHyorou/golang-microservice-building/helper"
	"github.com/SennenHyorou/golang-microservice-building/model/domain"
)

type BuildingRepositoryImpl struct {
}

func (b *BuildingRepositoryImpl) Store(ctx context.Context, tx *sql.Tx, building domain.Building) domain.Building {
	SQL := "insert into classroom_buildings(code,name) values(?)"
	result, err := tx.ExecContext(ctx, SQL, building.Code, building.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	building.Id = int(id)

	return building
}

func (b *BuildingRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, building domain.Building) domain.Building {
	SQL := "update classroom_buildings set code = ?, name = ? where id=?"
	_, err := tx.ExecContext(ctx, SQL, building.Code, building.Name, building.Id)
	helper.PanicIfError(err)

	return building
}

func (b *BuildingRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, building domain.Building) {
	SQL := "delete from classroom_buildings where id=?"
	_, err := tx.ExecContext(ctx, SQL, building.Code, building.Name, building.Id)
	helper.PanicIfError(err)
}

func (b *BuildingRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, buildingId int) (domain.Building, error) {
	SQL := "select id, code, name from classroom_buildings where id=?"
	rows, err := tx.QueryContext(ctx, SQL, buildingId)
	helper.PanicIfError(err)

	building := domain.Building{}
	if rows.Next() {
		err := rows.Scan(&building.Id, &building.Code, &building.Name)
		helper.PanicIfError(err)
		return building, nil
	} else {
		return building, errors.New("building is not found")
	}
}

func (b BuildingRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Building {
	SQL := "select id, code, name from classroom_buildings"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var buildings []domain.Building
	for rows.Next() {
		building := domain.Building{}
		err := rows.Scan(&building.Id, &building.Code, &building.Name)
		helper.PanicIfError(err)
		buildings = append(buildings, building)
	}
	return buildings
}
