package web

type BuildingCreateRequest struct {
	Code string `validate:"required" json:"code"`
	Name string `validate:"required,max=255" json:"name"`
}
