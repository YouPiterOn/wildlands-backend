package application

import "youpiteron.dev/wildlands-backend/internal/api"

type MatchService struct {
	eventStore api.EventStore
}

func NewMatchService(eventStore api.EventStore) *MatchService {
	return &MatchService{eventStore: eventStore}
}
