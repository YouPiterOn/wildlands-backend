package transport

import (
	"context"

	"github.com/coder/websocket"
	"youpiteron.dev/wildlands-backend/internal/application"
)

type WSHandler struct {
	matchService *application.MatchService
}

func (h *WSHandler) handleConnection(conn *websocket.Conn) {
	for {
		var msg JsonCommand

		err := conn.Read(conn, context.Background(), &msg)
		if err != nil {
			return
		}

		cmd, err := ParseCommand(msg)
		if err != nil {
			continue
		}

		events, err := h.matchService.Handle(context.Background(), cmd)
		if err != nil {
			continue
		}

		wsjson.Write(conn, context.Background(), events)
	}
}
