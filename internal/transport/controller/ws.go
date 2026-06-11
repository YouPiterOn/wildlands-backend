package controller

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"youpiteron.dev/wildlands-backend/internal/api"
	"youpiteron.dev/wildlands-backend/internal/domain"
	"youpiteron.dev/wildlands-backend/internal/transport/serializable"
)

type WS struct {
	matchService api.MatchService
	logger       api.Logger
}

func NewWS(matchService api.MatchService, logger api.Logger) *WS {
	return &WS{matchService: matchService, logger: logger.With(slog.String("tag", "WSController"))}
}

func (c *WS) WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close(websocket.StatusNormalClosure, "OK")
	ctx := r.Context()
	for {
		command, err := c.readCommand(ctx, conn)
		if err != nil {
			c.writeError(ctx, conn, err)
			continue
		}

		events, err := c.matchService.HandleCommand(ctx, command)
		if err != nil {
			c.writeError(ctx, conn, err)
			continue
		}

		err = c.writeEvents(ctx, conn, events)
		if err != nil {
			c.writeError(ctx, conn, err)
			continue
		}
	}
}

func (c *WS) readCommand(ctx context.Context, conn *websocket.Conn) (domain.Command, error) {
	var msg serializable.JsonCommand

	err := wsjson.Read(ctx, conn, &msg)
	if err != nil {
		return nil, err
	}

	return serializable.ToDomainCommand(msg)
}

func (c *WS) writeEvents(ctx context.Context, conn *websocket.Conn, events []domain.Event) error {
	jsonEvents := make([]serializable.JsonEvent, len(events))
	for i, event := range events {
		jsonEvent, err := serializable.ToJsonEvent(event)
		if err != nil {
			return err
		}
		jsonEvents[i] = jsonEvent
	}
	return wsjson.Write(ctx, conn, jsonEvents)
}

func (c *WS) writeError(ctx context.Context, conn *websocket.Conn, err error) error {
	jsonError := serializable.ToJsonError(err)
	return wsjson.Write(ctx, conn, jsonError)
}
