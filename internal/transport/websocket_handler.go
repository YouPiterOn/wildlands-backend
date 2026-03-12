package transport

import (
	"context"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"youpiteron.dev/wildlands-backend/internal/application"
)

type WSHandler struct {
	matchService *application.MatchService
}

func (h *WSHandler) handleConnection(ctx context.Context, conn *websocket.Conn) {
	for {
		var msg JsonCommand

		err := wsjson.Read(ctx, conn, &msg)
		if err != nil {
			return
		}

		command, err := ParseCommand(msg)
		if err != nil {
			continue
		}

		events, err := h.matchService.HandleCommand(ctx, command)
		if err != nil {
			continue
		}

		jsonEvents := make([]JsonEvent, len(events))
		for i, event := range events {
			jsonEvent, err := ToJsonEvent(event)
			if err != nil {
				continue
			}
			jsonEvents[i] = jsonEvent
		}
		wsjson.Write(ctx, conn, jsonEvents)
	}
}
