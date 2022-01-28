package service

import (
	"context"
	"database/sql"
	"golang-microservice-building/helper"
	"golang-microservice-building/model/domain"
	"golang-microservice-building/model/web"
	"golang-microservice-building/repository"
)

type BuildingServiceImpl struct {
	BuildingRepository repository.BuildingRepository
	DB                 *sql.DB
}

func (service *BuildingServiceImpl) Create(ctx context.Context, request web.BuildingCreateRequest) web.BuildingResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	building := domain.Building{
		Code: request.Code,
		Name: request.Name,
	}

	building = service.BuildingRepository.Store(ctx, tx, building)

	return helper.ToBuildingResponse(building)
}

func (service *BuildingServiceImpl) Update(ctx context.Context, request web.BuildingUpdateRequest) web.BuildingResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	building, err := service.BuildingRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	building.Code = request.Code
	building.Name = request.Name

	building = service.BuildingRepository.Update(ctx, tx, building)

	return helper.ToBuildingResponse(building)
}

func (service *BuildingServiceImpl) Delete(ctx context.Context, buildingId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	building, err := service.BuildingRepository.FindById(ctx, tx, buildingId)
	helper.PanicIfError(err)

	service.BuildingRepository.Delete(ctx, tx, building)
}

func (service *BuildingServiceImpl) FindById(ctx context.Context, buildingId int) web.BuildingResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	building, err := service.BuildingRepository.FindById(ctx, tx, buildingId)
	helper.PanicIfError(err)
	return helper.ToBuildingResponse(building)
}

func (service BuildingServiceImpl) FindAll(ctx context.Context) []web.BuildingResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	buildings := service.BuildingRepository.FindAll(ctx, tx)

	return helper.ToBuildingResponses(buildings)
}
