package httphandler

import (
	"ms-soccer/service/service-team-player/dto/request"
	"ms-soccer/service/service-team-player/mapper"
	"ms-soccer/service/service-team-player/models"
	"ms-soccer/service/service-team-player/repository/playerrepo"
	"ms-soccer/service/service-team-player/repository/teamrepo"
	"ms-soccer/service/shared/domains"
	"net/http"

	"github.com/gin-gonic/gin"
)

type teamHandler struct {
	TeamRepo   teamrepo.TeamRepository
	PlayerRepo playerrepo.PlayerRepository
}

func NewTeamHandler(teamRepo teamrepo.TeamRepository, playerRepo playerrepo.PlayerRepository) *teamHandler {
	return &teamHandler{
		TeamRepo:   teamRepo,
		PlayerRepo: playerRepo,
	}
}

func (h *teamHandler) GetList(c *gin.Context) {
	teams, err := h.TeamRepo.Fetch()
	if err != nil {
		domains.InternalServerError(c)
		return
	}

	domains.SuccessResponseWithData(c, http.StatusOK, mapper.ToListTeamResponseDTOs(teams))
}

func (h *teamHandler) Store(c *gin.Context) {
	var bodyReq request.TeamRequestDTO

	if err := c.ShouldBindJSON(&bodyReq); err != nil {
		domains.BadRequest(c, err.Error())
		return
	}

	var team models.Team
	team.Name = bodyReq.Name

	if _, err := h.TeamRepo.SaveOrUpdate(&team, nil); err != nil {
		domains.InternalServerError(c)
		return
	}

	domains.SuccessResponseWithoutData(c, http.StatusCreated, "Success")
}

func (h *teamHandler) StorePlayer(c *gin.Context) {
	var bodyReq request.PlayerRequestDTO
	id := c.Param("id")

	if err := c.ShouldBindJSON(&bodyReq); err != nil {
		domains.BadRequest(c, err.Error())
		return
	}

	p, _ := h.PlayerRepo.FindByTeamIDAndNumber(id, bodyReq.Number)
	if p != nil {
		domains.UnprocessableEntity(c, "The number already exists")
		return
	}

	var player models.Player
	player.Name = bodyReq.Name
	player.TeamID = id
	player.Number = bodyReq.Number

	if _, err := h.PlayerRepo.SaveOrUpdate(&player, nil); err != nil {
		domains.InternalServerError(c)
		return
	}

	domains.SuccessResponseWithoutData(c, http.StatusCreated, "Success")
}

func (h *teamHandler) GetListPlayer(c *gin.Context) {
	id := c.Param("id")

	players, err := h.PlayerRepo.FetchByTeamID(id)
	if err != nil {
		domains.InternalServerError(c)
		return
	}

	domains.SuccessResponseWithData(c, http.StatusOK, mapper.ToListPlayerResponseDTOs(players))
}
