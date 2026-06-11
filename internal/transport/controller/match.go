package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"youpiteron.dev/wildlands-backend/internal/api"
	"youpiteron.dev/wildlands-backend/internal/domain"
	"youpiteron.dev/wildlands-backend/internal/transport/serializable"
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
	var createMatchRequest serializable.CreateMatchRequest
	err := json.NewDecoder(r.Body).Decode(&createMatchRequest)
	if err != nil {
		c.logger.Error(utils.ErrRequestDecode.Error(), slog.String("error", err.Error()))
		http.Error(w, utils.ErrRequestDecode.Error(), http.StatusBadRequest)
		return
	}

	playerID, err := domain.ParsePlayerID(createMatchRequest.PlayerID)
	if err != nil {
		publicErr := utils.CustomErrUUIDParse("player_id", createMatchRequest.PlayerID)
		c.logger.Error(publicErr.Error(), slog.String("player_id", createMatchRequest.PlayerID), slog.String("error", err.Error()))
		http.Error(w, publicErr.Error(), http.StatusBadRequest)
		return
	}

	match, err := c.matchService.CreateMatch(r.Context(), playerID)
	if err != nil {
		http.Error(w, utils.ErrMatchCreate.Error(), http.StatusInternalServerError)
		return
	}

	response := serializable.CreateMatchResponse{
		Match: serializable.ToJsonMatch(match),
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		c.logger.Error(utils.ErrResponseEncode.Error(), slog.String("error", err.Error()))
		http.Error(w, utils.ErrResponseEncode.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *Match) JoinMatch(w http.ResponseWriter, r *http.Request) {
	var joinMatchRequest serializable.JoinMatchRequest
	err := json.NewDecoder(r.Body).Decode(&joinMatchRequest)
	if err != nil {
		c.logger.Error(utils.ErrRequestDecode.Error(), slog.String("error", err.Error()))
		http.Error(w, utils.ErrRequestDecode.Error(), http.StatusBadRequest)
		return
	}

	playerID, err := domain.ParsePlayerID(joinMatchRequest.PlayerID)
	if err != nil {
		publicErr := utils.CustomErrUUIDParse("player_id", joinMatchRequest.PlayerID)
		c.logger.Error(publicErr.Error(), slog.String("player_id", joinMatchRequest.PlayerID), slog.String("error", err.Error()))
		http.Error(w, publicErr.Error(), http.StatusBadRequest)
		return
	}

	matchID, err := domain.ParseMatchID(joinMatchRequest.MatchID)
	if err != nil {
		publicErr := utils.CustomErrUUIDParse("match_id", joinMatchRequest.MatchID)
		c.logger.Error(publicErr.Error(), slog.String("match_id", joinMatchRequest.MatchID), slog.String("error", err.Error()))
		http.Error(w, publicErr.Error(), http.StatusBadRequest)
		return
	}

	match, err := c.matchService.JoinMatch(r.Context(), matchID, playerID)
	if err != nil {
		http.Error(w, utils.ErrMatchJoin.Error(), http.StatusInternalServerError)
		return
	}

	response := serializable.JoinMatchResponse{
		Match: serializable.ToJsonMatch(match),
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		c.logger.Error(utils.ErrResponseEncode.Error(), slog.String("error", err.Error()))
		http.Error(w, utils.ErrResponseEncode.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
