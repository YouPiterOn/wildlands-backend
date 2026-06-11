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
	"youpiteron.dev/wildlands-backend/internal/utils"
)

type WS struct {
	matchService api.MatchService
	logger       api.Logger
}

func NewWS(matchService api.MatchService, logger api.Logger) *WS {
	return &WS{matchService: matchService, logger: logger.With(slog.String("tag", "WSController"))}
}

func (c *WS) WsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		c.logger.Error(utils.ErrWsAccept.Error(), slog.String("error", err.Error()))
		return
	}

	defer conn.Close(websocket.StatusNormalClosure, "OK")
	ctx := r.Context()
	for {
		command, err := c.readCommand(ctx, conn)
		if err != nil {
			c.writeError(ctx, conn, utils.ErrWsCommandRead)
			continue
		}

		events, err := c.matchService.HandleCommand(ctx, command)
		if err != nil {
			c.writeError(ctx, conn, utils.ErrCommandHandle)
			continue
		}

		err = c.writeEvents(ctx, conn, events)
		if err != nil {
			c.writeError(ctx, conn, utils.ErrWsEventWrite)
			continue
		}
	}
}

func (c *WS) readCommand(ctx context.Context, conn *websocket.Conn) (domain.Command, error) {
	var msg serializable.JsonCommand

	err := wsjson.Read(ctx, conn, &msg)
	if err != nil {
		c.logger.Error(utils.ErrWsCommandRead.Error(), slog.String("error", err.Error()))
		return nil, utils.ErrWsCommandRead
	}

	return serializable.ToDomainCommand(msg)
}

func (c *WS) writeEvents(ctx context.Context, conn *websocket.Conn, events []domain.Event) error {
	jsonEvents := make([]serializable.JsonEvent, len(events))
	for i, event := range events {
		jsonEvent, err := serializable.ToJsonEvent(event)
		if err != nil {
			c.logger.Error(utils.ErrWsEventSerialize.Error(), slog.String("error", err.Error()))
			return utils.ErrWsEventSerialize
		}
		jsonEvents[i] = jsonEvent
	}
	err := wsjson.Write(ctx, conn, jsonEvents)
	if err != nil {
		c.logger.Error(utils.ErrWsEventWrite.Error(), slog.String("error", err.Error()))
		return utils.ErrWsEventWrite
	}
	return nil
}

func (c *WS) writeError(ctx context.Context, conn *websocket.Conn, err error) error {
	jsonError := serializable.ToJsonError(err)
	err = wsjson.Write(ctx, conn, jsonError)
	if err != nil {
		c.logger.Error(utils.ErrWsErrorWrite.Error(), slog.String("error", err.Error()))
		return utils.ErrWsErrorWrite
	}
	return nil
}
