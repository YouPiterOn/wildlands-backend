package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"youpiteron.dev/wildlands-backend/internal/api"
	"youpiteron.dev/wildlands-backend/internal/transport/serializable"
	"youpiteron.dev/wildlands-backend/internal/utils"
)

type Player struct {
	playerService api.PlayerService
	logger        api.Logger
}

func NewPlayer(playerService api.PlayerService, logger api.Logger) *Player {
	return &Player{playerService: playerService, logger: logger.With(slog.String("tag", "PlayerController"))}
}

func (p *Player) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var createPlayer serializable.CreatePlayerRequest
	err := json.NewDecoder(r.Body).Decode(&createPlayer)
	if err != nil {
		p.logger.Error(utils.ErrRequestDecode.Error(), slog.String("error", err.Error()))
		http.Error(w, utils.ErrRequestDecode.Error(), http.StatusBadRequest)
		return
	}

	player, err := p.playerService.Create(r.Context(), createPlayer.Name)
	if err != nil {
		http.Error(w, utils.ErrPlayerCreate.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := serializable.CreatePlayerResponse{
		Name:     player.Name,
		PlayerID: player.ID.String(),
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		p.logger.Error(utils.ErrResponseEncode.Error(), slog.String("player_id", player.ID.String()), slog.String("error", err.Error()))
		http.Error(w, utils.ErrResponseEncode.Error(), http.StatusInternalServerError)
		return
	}
}
