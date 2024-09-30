package dto

type PaginationDto struct {
	Page    int `form:"page" query:"page"    example:"1"  default:"1"`
	PerPage int `form:"perPage" query:"perPage" example:"5"  default:"10"`
}
