package api

import "youpiteron.dev/wildlands-backend/internal/domain"

type EventStore interface {
	Apply(event domain.Event) error
	GetEvents(matchID string) ([]domain.Event, error)
}
