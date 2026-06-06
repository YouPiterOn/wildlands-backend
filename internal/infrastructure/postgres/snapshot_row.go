package postgres

import (
	"encoding/json"

	"youpiteron.dev/wildlands-backend/internal/domain"
)

type SnapshotRow struct {
	MatchID     string
	State       string
	CurrentTurn int
	BoardsJson  []byte
	Version     int
}

func (s SnapshotRow) ToDomainMatch() (*domain.Match, error) {
	matchID, err := domain.ParseMatchID(s.MatchID)
	if err != nil {
		return nil, err
	}
	state, err := domain.ParseMatchState(s.State)
	if err != nil {
		return nil, err
	}
	boards := []domain.Board{}
	err = json.Unmarshal(s.BoardsJson, &boards)
	if err != nil {
		return nil, err
	}
	return &domain.Match{
		ID:          matchID,
		State:       state,
		Boards:      boards,
		CurrentTurn: s.CurrentTurn,
		Version:     s.Version,
	}, nil
}

func ToSnapshotRow(match *domain.Match) (SnapshotRow, error) {
	boardsJson, err := json.Marshal(match.Boards)
	if err != nil {
		return SnapshotRow{}, err
	}
	return SnapshotRow{
		MatchID:     match.ID.String(),
		State:       match.State.String(),
		CurrentTurn: match.CurrentTurn,
		BoardsJson:  boardsJson,
		Version:     match.Version,
	}, nil
}
