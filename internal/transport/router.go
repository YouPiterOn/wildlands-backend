package transport

import (
	"net/http"

	"youpiteron.dev/wildlands-backend/internal/transport/controller"
)

func NewRouter(
	matchController *controller.Match,
	playerController *controller.Player,
	wsController *controller.WS,
) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", wsController.WsHandler)
	mux.HandleFunc("/match/create", matchController.CreateMatch)
	mux.HandleFunc("/match/join", matchController.JoinMatch)
	mux.HandleFunc("/player/create", playerController.CreatePlayer)

	return mux
}
