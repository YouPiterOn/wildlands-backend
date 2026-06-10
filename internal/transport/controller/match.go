package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"youpiteron.dev/wildlands-backend/internal/api"
	"youpiteron.dev/wildlands-backend/internal/domain"
	"youpiteron.dev/wildlands-backend/internal/transport/dto"
	"youpiteron.dev/wildlands-backend/internal/utils"
)

type Match struct {
	matchService api.MatchService
	logger       api.Logger
}

func NewMatch(matchService api.MatchService, logger api.Logger) *Match {
	return &Match{matchService: matchService, logger: logger.With(slog.String("tag", "MatchController"))}
}

func (c *Match) CreateMatch(w http.ResponseWriter, r *http.Request) {
	var createMatchRequest dto.CreateMatchRequest
	err := json.NewDecoder(r.Body).Decode(&createMatchRequest)
	if err != nil {
		c.logger.Error(utils.ErrRequestDecode.Error(), slog.String("error", err.Error()))
		http.Error(w, utils.ErrRequestDecode.Error(), http.StatusBadRequest)
		return
	}

	playerID, err := domain.ParsePlayerID(createMatchRequest.PlayerID)
	if err != nil {
		c.logger.Error(utils.ErrUUIDParse.Error(), slog.String("player_id", createMatchRequest.PlayerID), slog.String("error", err.Error()))
		http.Error(w, utils.ErrUUIDParse.Error(), http.StatusBadRequest)
		return
	}

	matchID, err := c.matchService.CreateMatch(r.Context(), playerID)
	if err != nil {
		http.Error(w, utils.ErrMatchCreate.Error(), http.StatusInternalServerError)
		return
	}

	response := dto.CreateMatchResponse{
		MatchID: matchID.String(),
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		c.logger.Error(utils.ErrResponseEncode.Error(), slog.String("match_id", matchID.String()), slog.String("error", err.Error()))
		http.Error(w, utils.ErrResponseEncode.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
