package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"youpiteron.dev/wildlands-backend/internal/api"
	"youpiteron.dev/wildlands-backend/internal/domain"
)

type EventStore struct {
	pool *pgxpool.Pool
}

var _ api.EventStore = (*EventStore)(nil)

func NewEventStore(pool *pgxpool.Pool) *EventStore {
	return &EventStore{pool: pool}
}

func (s *EventStore) Append(ctx context.Context, matchID domain.MatchID, expectedVersion int, events []domain.Event) error {
	version := expectedVersion + 1
	storedEvents := make([]EventRow, len(events))
	for i, event := range events {
		storedEvent, err := ToEventRow(event, version+i)
		if err != nil {
			return err
		}
		storedEvents[i] = storedEvent
	}
	err := s.saveEvents(ctx, storedEvents)
	if err != nil {
		return err
	}
	return nil
}

func (s *EventStore) LoadAll(ctx context.Context, matchID domain.MatchID) ([]domain.Event, int, error) {
	storedEvents, err := s.getEvents(ctx, matchID, 0)
	if err != nil {
		return nil, 0, err
	}
	events := make([]domain.Event, len(storedEvents))
	for i, storedEvent := range storedEvents {
		events[i], err = storedEvent.ToDomainEvent()
		if err != nil {
			return nil, 0, err
		}
	}
	return events, storedEvents[len(storedEvents)-1].Version, nil
}

func (s *EventStore) LoadSince(ctx context.Context, matchID domain.MatchID, version int) ([]domain.Event, int, error) {
	storedEvents, err := s.getEvents(ctx, matchID, version)
	if err != nil {
		return nil, 0, err
	}
	events := make([]domain.Event, len(storedEvents))
	for i, storedEvent := range storedEvents {
		events[i], err = storedEvent.ToDomainEvent()
		if err != nil {
			return nil, 0, err
		}
	}
	return events, storedEvents[len(storedEvents)-1].Version, nil
}

func (s *EventStore) saveEvents(ctx context.Context, events []EventRow) error {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	for _, event := range events {
		_, err = tx.Exec(
			ctx,
			`
      INSERT INTO events (match_id, version, type, data, metadata, created_at)
      VALUES ($1, $2, $3, $4, $5, $6)
      `,
			event.MatchID, event.Version, event.Type, event.Data, event.Metadata, event.CreatedAt,
		)
		if err != nil {
			return err
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *EventStore) getEvents(ctx context.Context, matchID domain.MatchID, version int) ([]EventRow, error) {
	matchIDString := matchID.String()
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	rows, err := conn.Query(
		ctx,
		`
    SELECT id, match_id, version, type, data, metadata, created_at FROM events
    WHERE match_id = $1 AND version > $2
    ORDER BY version ASC
    `,
		matchIDString, version,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []EventRow{}
	for rows.Next() {
		var event EventRow
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
