package request

type TeamRequestDTO struct {
	Name string `json:"name" binding:"required"`
}
