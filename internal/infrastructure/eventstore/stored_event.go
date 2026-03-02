package eventstore

import (
	"time"

	"youpiteron.dev/wildlands-backend/internal/domain"
)

type StoredEvent struct {
	ID        string
	MatchID   string
	Version   int
	Type      domain.EventType
	Data      []byte
	Metadata  []byte
	CreatedAt time.Duration
}
