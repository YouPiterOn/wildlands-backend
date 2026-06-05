package transport

import (
	"context"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"youpiteron.dev/wildlands-backend/internal/application"
	"youpiteron.dev/wildlands-backend/internal/domain"
)

type WSHandler struct {
	matchService *application.MatchService
}

func NewWSHandler(matchService *application.MatchService) *WSHandler {
	return &WSHandler{matchService: matchService}
}

func (h *WSHandler) HandleConnection(ctx context.Context, conn *websocket.Conn) {
	for {
		command, err := h.readCommand(ctx, conn)
		if err != nil {
			h.writeError(ctx, conn, err)
			continue
		}

		events, err := h.matchService.HandleCommand(ctx, command)
		if err != nil {
			h.writeError(ctx, conn, err)
			continue
		}

		err = h.writeEvents(ctx, conn, events)
		if err != nil {
			h.writeError(ctx, conn, err)
			continue
		}
	}
}

func (h *WSHandler) readCommand(ctx context.Context, conn *websocket.Conn) (domain.Command, error) {
	var msg JsonCommand

	err := wsjson.Read(ctx, conn, &msg)
	if err != nil {
		return nil, err
	}

	return ParseCommand(msg)
}

func (h *WSHandler) writeEvents(ctx context.Context, conn *websocket.Conn, events []domain.Event) error {
	jsonEvents := make([]JsonEvent, len(events))
	for i, event := range events {
		jsonEvent, err := ToJsonEvent(event)
		if err != nil {
			return err
		}
		jsonEvents[i] = jsonEvent
	}
	return wsjson.Write(ctx, conn, jsonEvents)
}

func (h *WSHandler) writeError(ctx context.Context, conn *websocket.Conn, err error) error {
	jsonError := ToJsonError(err)
	return wsjson.Write(ctx, conn, jsonError)
}
