package web

type BuildingUpdateRequest struct {
	Id   int    `validate:"required"`
	Code string `validate:"required"`
	Name string `validate:"required"`
}
