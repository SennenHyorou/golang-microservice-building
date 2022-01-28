package helper

import (
	"github.com/SennenHyorou/golang-microservice-building/model/domain"
	"github.com/SennenHyorou/golang-microservice-building/model/web"
)

func ToBuildingResponse(building domain.Building) web.BuildingResponse {
	return web.BuildingResponse{
		Id:   building.Id,
		Code: building.Code,
		Name: building.Name,
	}
}

func ToBuildingResponses(buildings []domain.Building) []web.BuildingResponse {
	var buildingResponses []web.BuildingResponse
	for _, building := range buildings {
		buildingResponses = append(buildingResponses, ToBuildingResponse(building))
	}
	return buildingResponses
}
