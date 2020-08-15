package mapper

import (
	"ms-soccer/service/service-team-player/models"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestToPlayerResponseDTO(t *testing.T) {
	player := models.Player{}
	player.ID = uuid.NewV4().String()
	player.TeamID = uuid.NewV4().String()
	player.Name = "John"
	player.Number = 10

	resp := ToPlayerResponseDTO(&player)

	assert.Equal(t, resp.ID, player.ID)
	assert.Equal(t, resp.Name, player.Name)
	assert.Equal(t, resp.Number, player.Number)

}

func TestToListPlayerResponseDTOs(t *testing.T) {
	t.Run("test with empty players", func(t *testing.T) {
		resp := ToListPlayerResponseDTOs(&[]models.Player{})
		assert.Equal(t, len(resp), 0)
	})

	t.Run("test with full team", func(t *testing.T) {
		player1 := models.Player{}
		player1.ID = uuid.NewV4().String()
		player1.TeamID = uuid.NewV4().String()
		player1.Name = "John"
		player1.Number = 10

		player2 := models.Player{}
		player2.ID = uuid.NewV4().String()
		player2.TeamID = uuid.NewV4().String()
		player2.Name = "Berbatov"
		player2.Number = 11

		players := []models.Player{player1, player2}
		resp := ToListPlayerResponseDTOs(&players)

		assert.Equal(t, len(resp), 2)
	})
}
