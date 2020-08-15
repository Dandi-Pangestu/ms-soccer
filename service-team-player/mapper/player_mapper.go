package mapper

import (
	"ms-soccer/service/service-team-player/dto/response"
	"ms-soccer/service/service-team-player/models"
)

func ToPlayerResponseDTO(entity *models.Player) response.PlayerResponseDTO {
	return response.PlayerResponseDTO{
		ID:     entity.ID,
		Name:   entity.Name,
		Number: entity.Number,
	}
}

func ToListPlayerResponseDTOs(entities *[]models.Player) []response.PlayerResponseDTO {
	dtos := make([]response.PlayerResponseDTO, len(*entities))
	for i, item := range *entities {
		dto := ToPlayerResponseDTO(&item)
		dtos[i] = dto
	}

	return dtos
}
