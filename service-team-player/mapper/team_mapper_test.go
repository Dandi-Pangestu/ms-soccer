package mapper

import (
	"ms-soccer/service/service-team-player/models"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestToTeamResponseDTO(t *testing.T) {
	team := models.Team{}
	team.ID = uuid.NewV4().String()
	team.Name = "Chelsea"

	resp := ToTeamResponseDTO(&team)

	assert.Equal(t, resp.ID, team.ID)
	assert.Equal(t, resp.Name, team.Name)
}

func TestToListTeamResponseDTOs(t *testing.T) {
	t.Run("test with empty teams", func(t *testing.T) {
		resp := ToListTeamResponseDTOs(&[]models.Team{})
		assert.Equal(t, len(resp), 0)
	})

	t.Run("test with all team", func(t *testing.T) {
		team1 := models.Team{}
		team1.ID = uuid.NewV4().String()
		team1.Name = "Chelsea"

		team2 := models.Team{}
		team2.ID = uuid.NewV4().String()
		team2.Name = "Barcelona"

		teams := []models.Team{team1, team2}
		resp := ToListTeamResponseDTOs(&teams)

		assert.Equal(t, len(resp), 2)
	})
}
