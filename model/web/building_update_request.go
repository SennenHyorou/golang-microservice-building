package web

type BuildingUpdateRequest struct {
	Id   int    `validate:"required"`
	Code string `validate:"required" json:"code"`
	Name string `validate:"required" json:"name"`
}
