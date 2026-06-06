package domain

type MatchMetadata struct {
	MatchID          MatchID
	BoardsSeed       int64
	ExploreDeckSeed  int64
	ScoringCardsSeed int64
}

func NewMatchMetadata(matchID MatchID, boardsSeed int64, exploreDeckSeed int64, scoringCardsSeed int64) *MatchMetadata {
	return &MatchMetadata{
		MatchID:          matchID,
		BoardsSeed:       boardsSeed,
		ExploreDeckSeed:  exploreDeckSeed,
		ScoringCardsSeed: scoringCardsSeed,
	}
}
