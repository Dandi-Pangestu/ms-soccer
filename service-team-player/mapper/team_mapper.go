package mapper

import (
	"ms-soccer/service/service-team-player/dto/response"
	"ms-soccer/service/service-team-player/models"
)

func ToTeamResponseDTO(entity *models.Team) response.TeamResponseDTO {
	return response.TeamResponseDTO{
		ID:   entity.ID,
		Name: entity.Name,
	}
}

func ToListTeamResponseDTOs(entities *[]models.Team) []response.TeamResponseDTO {
	dtos := make([]response.TeamResponseDTO, len(*entities))
	for i, item := range *entities {
		dto := ToTeamResponseDTO(&item)
		dtos[i] = dto
	}

	return dtos
}
