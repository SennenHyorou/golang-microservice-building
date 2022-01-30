package main

import (
	"github.com/SennenHyorou/golang-microservice-building/app"
	"github.com/SennenHyorou/golang-microservice-building/config"
	"github.com/SennenHyorou/golang-microservice-building/controller"
	"github.com/SennenHyorou/golang-microservice-building/exception"
	"github.com/SennenHyorou/golang-microservice-building/helper"
	"github.com/SennenHyorou/golang-microservice-building/middleware"
	"github.com/SennenHyorou/golang-microservice-building/repository"
	"github.com/SennenHyorou/golang-microservice-building/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	conf := config.Get() // Load all config from Env first

	db := app.NewDB(conf)
	validate := validator.New()
	buildingRepository := repository.NewBuildingRepository()
	buildingService := service.NewBuildingService(buildingRepository, db, validate)
	buildingController := controller.NewBuildingController(buildingService)

	router := httprouter.New()
	router.GET("/api/buildings", buildingController.FindAll)
	router.GET("/api/buildings/:buildingId", buildingController.FindById)
	router.POST("/api/buildings", buildingController.Create)
	router.PUT("/api/buildings/:buildingId", buildingController.Update)
	router.DELETE("/api/buildings/:buildingId", buildingController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost" + conf.GetAPIPort(),
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
