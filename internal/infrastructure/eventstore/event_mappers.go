package eventstore

import (
	"encoding/json"
	"errors"

	"youpiteron.dev/wildlands-backend/internal/domain"
)

func (s *PostgresEventStore) StoredEventToDomainEvent(storedEvent StoredEvent) (domain.Event, error) {
	mapper, ok := toDomainEventMapperRegistry[storedEvent.Type]
	if !ok {
		return nil, errors.New("invalid event type")
	}
	return mapper(storedEvent)
}

func (s *PostgresEventStore) DomainEventToStoredEvent(domainEvent domain.Event) (StoredEvent, error) {
	mapper, ok := toStoredEventMapperRegistry[domainEvent.EventType()]
	if !ok {
		return StoredEvent{}, errors.New("invalid event type")
	}
	return mapper(domainEvent)
}

var toDomainEventMapperRegistry = map[domain.EventType]func(StoredEvent) (domain.Event, error){
	domain.EventTypeMatchCreated: func(storedEvent StoredEvent) (domain.Event, error) {
		var event MatchCreatedEventData
		err := json.Unmarshal(storedEvent.Data, &event)
		if err != nil {
			return nil, err
		}
		matchID, err := domain.ParseMatchID(storedEvent.MatchID)
		if err != nil {
			return nil, err
		}
		return &domain.EventMatchCreated{
			MatchID:    matchID,
			SeatsCount: event.SeatsCount,
		}, nil
	},
	domain.EventTypePlayerJoined: func(storedEvent StoredEvent) (domain.Event, error) {
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
			MatchID:    matchID,
			PlayerID:   playerID,
			SeatNumber: event.SeatNumber,
		}, nil
	},
	domain.EventTypeGameStarted: func(storedEvent StoredEvent) (domain.Event, error) {
		matchID, err := domain.ParseMatchID(storedEvent.MatchID)
		if err != nil {
			return nil, err
		}
		return &domain.EventGameStarted{
			MatchID: matchID,
		}, nil
	},
}

var toStoredEventMapperRegistry = map[domain.EventType]func(domain.Event) (StoredEvent, error){
	domain.EventTypeMatchCreated: func(event domain.Event) (StoredEvent, error) {
		domainEvent, ok := event.(*domain.EventMatchCreated)
		if !ok {
			return StoredEvent{}, errors.New("invalid domain event type")
		}
		data, err := json.Marshal(MatchCreatedEventData{
			SeatsCount: domainEvent.SeatsCount,
		})
		if err != nil {
			return StoredEvent{}, err
		}
		return StoredEvent{
			Type:    domainEvent.EventType(),
			MatchID: domainEvent.MatchID.String(),
			Data:    data,
		}, nil
	},
	domain.EventTypePlayerJoined: func(event domain.Event) (StoredEvent, error) {
		domainEvent, ok := event.(*domain.EventPlayerJoined)
		if !ok {
			return StoredEvent{}, errors.New("invalid domain event type")
		}
		data, err := json.Marshal(PlayerJoinedEventData{
			SeatNumber: domainEvent.SeatNumber,
			PlayerID:   domainEvent.PlayerID.String(),
		})
		if err != nil {
			return StoredEvent{}, err
		}
		return StoredEvent{
			Type:    domainEvent.EventType(),
			MatchID: domainEvent.MatchID.String(),
			Data:    data,
		}, nil
	},
	domain.EventTypeGameStarted: func(event domain.Event) (StoredEvent, error) {
		domainEvent, ok := event.(*domain.EventGameStarted)
		if !ok {
			return StoredEvent{}, errors.New("invalid domain event type")
		}
		return StoredEvent{
			Type:    domainEvent.EventType(),
			MatchID: domainEvent.MatchID.String(),
		}, nil
	},
}
