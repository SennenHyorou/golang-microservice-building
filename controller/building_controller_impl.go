package controller

import (
	"github.com/SennenHyorou/golang-microservice-building/helper"
	"github.com/SennenHyorou/golang-microservice-building/model/web"
	"github.com/SennenHyorou/golang-microservice-building/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type BuildingControllerImpl struct {
	BuildingService service.BuildingService
}

func NewBuildingController(buildingService service.BuildingService) BuildingController {
	return &BuildingControllerImpl{
		BuildingService: buildingService,
	}
}

func (controller BuildingControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	buildingCreateRequest := web.BuildingCreateRequest{}
	helper.ReadFromRequestBody(request, &buildingCreateRequest)

	buildingResponse := controller.BuildingService.Create(request.Context(), buildingCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   buildingResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller BuildingControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	buildingUpdateRequest := web.BuildingUpdateRequest{}
	helper.ReadFromRequestBody(request, &buildingUpdateRequest)

	buildingId := params.ByName("buildingId")
	id, err := strconv.Atoi(buildingId)
	helper.PanicIfError(err)

	buildingUpdateRequest.Id = id

	buildingResponse := controller.BuildingService.Update(request.Context(), buildingUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   buildingResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller BuildingControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	buildingId := params.ByName("buildingId")
	id, err := strconv.Atoi(buildingId)
	helper.PanicIfError(err)

	controller.BuildingService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller BuildingControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	buildingId := params.ByName("buildingId")
	id, err := strconv.Atoi(buildingId)
	helper.PanicIfError(err)

	buildingResponse := controller.BuildingService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   buildingResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller BuildingControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	buildingResponses := controller.BuildingService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   buildingResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
