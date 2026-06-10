package postgres

import (
	"encoding/json"
	"errors"
	"time"

	"youpiteron.dev/wildlands-backend/internal/domain"
)

type EventRow struct {
	ID        string
	MatchID   string
	Version   int
	Type      domain.EventType
	Data      []byte
	Metadata  []byte
	CreatedAt time.Duration
}

func (s EventRow) ToDomainEvent() (domain.Event, error) {
	mapper, ok := toDomainEventMapperRegistry[s.Type]
	if !ok {
		return nil, errors.New("invalid event type")
	}
	return mapper(s)
}

func ToEventRow(event domain.Event, version int) (EventRow, error) {
	mapper, ok := toStoredEventMapperRegistry[event.EventType()]
	if !ok {
		return EventRow{}, errors.New("invalid event type")
	}
	return mapper(event, version)
}

var toDomainEventMapperRegistry = map[domain.EventType]func(EventRow) (domain.Event, error){
	domain.EventTypePlayerJoined: func(storedEvent EventRow) (domain.Event, error) {
		var event PlayerJoinedEventData
		err := json.Unmarshal(storedEvent.Data, &event)
		if err != nil {
			return nil, err
		}
		playerID, err := domain.ParsePlayerID(event.PlayerID)
		if err != nil {
			return nil, err
		}
		matchID, err := domain.ParseMatchID(storedEvent.MatchID)
		if err != nil {
			return nil, err
		}
		return &domain.EventPlayerJoined{
			MatchID:  matchID,
			PlayerID: playerID,
		}, nil
	},
	domain.EventTypeGameStarted: func(storedEvent EventRow) (domain.Event, error) {
		matchID, err := domain.ParseMatchID(storedEvent.MatchID)
		if err != nil {
			return nil, err
		}
		return &domain.EventGameStarted{
			MatchID: matchID,
		}, nil
	},
}

var toStoredEventMapperRegistry = map[domain.EventType]func(domain.Event, int) (EventRow, error){
	domain.EventTypePlayerJoined: func(event domain.Event, version int) (EventRow, error) {
		domainEvent, ok := event.(*domain.EventPlayerJoined)
		if !ok {
			return EventRow{}, errors.New("invalid domain event type")
		}
		data, err := json.Marshal(PlayerJoinedEventData{
			PlayerID: domainEvent.PlayerID.String(),
		})
		if err != nil {
			return EventRow{}, err
		}
		return EventRow{
			Type:    domainEvent.EventType(),
			MatchID: domainEvent.MatchID.String(),
			Version: version,
			Data:    data,
		}, nil
	},
	domain.EventTypeGameStarted: func(event domain.Event, version int) (EventRow, error) {
		domainEvent, ok := event.(*domain.EventGameStarted)
		if !ok {
			return EventRow{}, errors.New("invalid domain event type")
		}
		return EventRow{
			Type:    domainEvent.EventType(),
			MatchID: domainEvent.MatchID.String(),
			Version: version,
		}, nil
	},
}
