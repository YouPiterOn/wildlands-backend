package transport

import (
	"errors"

	"youpiteron.dev/wildlands-backend/internal/domain"
)

type JsonEvent struct {
	Type     string `json:"type"`
	MatchID  string `json:"match_id,omitempty"`
	PlayerID string `json:"player_id,omitempty"`
}

func ToJsonEvent(event domain.Event) (JsonEvent, error) {
	switch event.EventType() {
	case domain.EventTypeMatchCreated:
		domainEvent, ok := event.(*domain.EventMatchCreated)
		if !ok {
			return JsonEvent{}, errors.New("invalid event type")
		}
		return JsonEvent{
			Type:    domain.EventTypeMatchCreated.String(),
			MatchID: domainEvent.MatchID.String(),
		}, nil
	case domain.EventTypePlayerJoined:
		domainEvent, ok := event.(*domain.EventPlayerJoined)
		if !ok {
			return JsonEvent{}, errors.New("invalid event type")
		}
		return JsonEvent{
			Type:     domain.EventTypePlayerJoined.String(),
			MatchID:  domainEvent.MatchID.String(),
			PlayerID: domainEvent.PlayerID.String(),
		}, nil
	case domain.EventTypeGameStarted:
		domainEvent, ok := event.(*domain.EventGameStarted)
		if !ok {
			return JsonEvent{}, errors.New("invalid event type")
		}
		return JsonEvent{
			Type:    domain.EventTypeGameStarted.String(),
			MatchID: domainEvent.MatchID.String(),
		}, nil
	}
	return JsonEvent{}, errors.New("invalid event type")
}
