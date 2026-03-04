package eventstore

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"youpiteron.dev/wildlands-backend/internal/domain"
)

type PostgresEventStore struct {
	pool *pgxpool.Pool
}

func NewPostgresEventStore(pool *pgxpool.Pool) *PostgresEventStore {
	return &PostgresEventStore{pool: pool}
}

func (s *PostgresEventStore) Append(ctx context.Context, events ...domain.Event) error {
	for _, event := range events {
		storedEvent, err := s.DomainEventToStoredEvent(event)
		if err != nil {
			return err
		}
		err = s.saveEvent(ctx, storedEvent)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *PostgresEventStore) Load(ctx context.Context, matchID string) ([]domain.Event, error) {
	storedEvents, err := s.getEvents(ctx, matchID)
	if err != nil {
		return nil, err
	}
	events := make([]domain.Event, len(storedEvents))
	for i, storedEvent := range storedEvents {
		events[i], err = s.StoredEventToDomainEvent(storedEvent)
		if err != nil {
			return nil, err
		}
	}
	return events, nil
}

func (s *PostgresEventStore) saveEvent(ctx context.Context, event StoredEvent) error {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(
		ctx,
		`
    INSERT INTO events (match_id, version, type, data, metadata, created_at)
    VALUES ($1, $2, $3, $4, $5, $6)
    `,
		event.MatchID, event.Version, event.Type, event.Data, event.Metadata, event.CreatedAt,
	)
	return err
}

func (s *PostgresEventStore) getEvents(ctx context.Context, matchID string) ([]StoredEvent, error) {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	rows, err := conn.Query(
		ctx,
		`
    SELECT id, match_id, version, type, data, metadata, created_at FROM events
    WHERE match_id = $1
    ORDER BY version ASC
    `,
		matchID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []StoredEvent{}
	for rows.Next() {
		var event StoredEvent
		err := rows.Scan(&event.ID, &event.MatchID, &event.Version, &event.Type, &event.Data, &event.Metadata, &event.CreatedAt)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}
