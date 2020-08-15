package request

type PlayerRequestDTO struct {
	Name   string `json:"name" binding:"required"`
	Number int    `json:"number" binding:"required"`
}
