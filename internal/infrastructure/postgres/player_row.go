package postgres

import "youpiteron.dev/wildlands-backend/internal/domain"

type PlayerRow struct {
	ID   string
	Name string
}

func (s PlayerRow) ToDomainPlayer() (*domain.Player, error) {
	playerID, err := domain.ParsePlayerID(s.ID)
	if err != nil {
		return nil, err
	}
	return &domain.Player{ID: playerID, Name: s.Name}, nil
}

func ToPlayerRow(player *domain.Player) (PlayerRow, error) {
	return PlayerRow{ID: player.ID.String(), Name: player.Name}, nil
}
