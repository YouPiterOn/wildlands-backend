package eventstore

import (
	"context"

	"github.com/jackc/pgx/v5"
	"youpiteron.dev/wildlands-backend/internal/domain"
)

type PostgresEventStore struct {
	conn *pgx.Conn
}

func NewPostgresEventStore(conn *pgx.Conn) *PostgresEventStore {
	return &PostgresEventStore{conn: conn}
}

func (s *PostgresEventStore) Apply(event domain.Event) error {
	return nil
}

func (s *PostgresEventStore) GetEvents(matchID string) ([]domain.Event, error) {
	return nil, nil
}

func (storedEvent StoredEvent) toDomainEvent() domain.Event {
	return nil
}

func (s *PostgresEventStore) saveEvent(event StoredEvent) error {
	_, err := s.conn.Exec(
		context.Background(),
		`
    INSERT INTO events (match_id, version, type, data, metadata, created_at)
    VALUES ($1, $2, $3, $4, $5, $6)
    `,
		event.MatchID, event.Version, event.Type, event.Data, event.Metadata, event.CreatedAt,
	)
	return err
}

func (s *PostgresEventStore) getEvents(matchID string) ([]StoredEvent, error) {
	rows, err := s.conn.Query(
		context.Background(),
		`
    SELECT id, match_id, version, type, data, metadata, created_at FROM events WHERE match_id = $1
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
