package domain

import "github.com/google/uuid"

type PlayerID uuid.UUID

func (id PlayerID) String() string {
	return uuid.UUID(id).String()
}

func ParsePlayerID(s string) (PlayerID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return PlayerID{}, err
	}
	return PlayerID(id), nil
}

type Player struct {
	ID   PlayerID
	Name string
}
