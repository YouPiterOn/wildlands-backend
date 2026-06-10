package postgres

import "youpiteron.dev/wildlands-backend/internal/domain"

type MatchMetadataRow struct {
	MatchID          string
	BoardsSeed       int64
	ExploreDeckSeed  int64
	ScoringCardsSeed int64
}

func (r MatchMetadataRow) ToDomainMatchMetadata() (*domain.MatchMetadata, error) {
	matchID, err := domain.ParseMatchID(r.MatchID)
	if err != nil {
		return nil, err
	}
	return &domain.MatchMetadata{
		MatchID:          matchID,
		BoardsSeed:       r.BoardsSeed,
		ExploreDeckSeed:  r.ExploreDeckSeed,
		ScoringCardsSeed: r.ScoringCardsSeed,
	}, nil
}

func ToMatchMetadataRow(metadata *domain.MatchMetadata) (MatchMetadataRow, error) {
	return MatchMetadataRow{
		MatchID:          metadata.MatchID.String(),
		BoardsSeed:       metadata.BoardsSeed,
		ExploreDeckSeed:  metadata.ExploreDeckSeed,
		ScoringCardsSeed: metadata.ScoringCardsSeed,
	}, nil
}
