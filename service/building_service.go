package service

import (
	"context"
	"golang-microservice-building/model/web"
)

type BuildingService interface {
	Create(ctx context.Context, request web.BuildingCreateRequest) web.BuildingResponse
	Update(ctx context.Context, request web.BuildingUpdateRequest) web.BuildingResponse
	Delete(ctx context.Context, buildingId int)
	FindById(ctx context.Context, buildingId int) web.BuildingResponse
	FindAll(ctx context.Context) []web.BuildingResponse
}
