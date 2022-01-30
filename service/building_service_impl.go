package service

import (
	"context"
	"database/sql"
	"github.com/SennenHyorou/golang-microservice-building/exception"
	"github.com/SennenHyorou/golang-microservice-building/helper"
	"github.com/SennenHyorou/golang-microservice-building/model/domain"
	"github.com/SennenHyorou/golang-microservice-building/model/web"
	"github.com/SennenHyorou/golang-microservice-building/repository"
	"github.com/go-playground/validator/v10"
)

type BuildingServiceImpl struct {
	BuildingRepository repository.BuildingRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewBuildingService(buildingRepository repository.BuildingRepository, DB *sql.DB, validate *validator.Validate) BuildingService {
	return &BuildingServiceImpl{
		BuildingRepository: buildingRepository,
		DB:                 DB,
		Validate:           validate}
}

func (service *BuildingServiceImpl) Create(ctx context.Context, request web.BuildingCreateRequest) web.BuildingResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

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
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	building, err := service.BuildingRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

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
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.BuildingRepository.Delete(ctx, tx, building)
}

func (service *BuildingServiceImpl) FindById(ctx context.Context, buildingId int) web.BuildingResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	building, err := service.BuildingRepository.FindById(ctx, tx, buildingId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToBuildingResponse(building)
}

func (service BuildingServiceImpl) FindAll(ctx context.Context) []web.BuildingResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	buildings := service.BuildingRepository.FindAll(ctx, tx)

	return helper.ToBuildingResponses(buildings)
}
