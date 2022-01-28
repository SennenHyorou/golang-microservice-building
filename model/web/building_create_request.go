package web

type BuildingCreateRequest struct {
	Code string `validate:"required"`
	Name string `validate:"required,max=255"`
}
