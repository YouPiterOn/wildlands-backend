package serializable

import "youpiteron.dev/wildlands-backend/internal/domain"

type JsonCell struct {
	Terrain  string `json:"terrain"`
	HasRuins bool   `json:"has_ruins"`
}

func ToJsonCell(cell domain.Cell) JsonCell {
	return JsonCell{
		Terrain:  cell.Terrain.String(),
		HasRuins: cell.HasRuins,
	}
}

type JsonBoard struct {
	BoardNumber int          `json:"board_number"`
	PlayerID    string       `json:"player_id"`
	Score       int          `json:"score"`
	Coins       int          `json:"coins"`
	Cells       [][]JsonCell `json:"cells"`
}

func ToJsonBoard(board *domain.Board) JsonBoard {
	cells := make([][]JsonCell, len(board.Cells))
	for i := range board.Cells {
		cells[i] = make([]JsonCell, len(board.Cells[i]))
		for j := range board.Cells[i] {
			cells[i][j] = ToJsonCell(board.Cells[i][j])
		}
	}
	return JsonBoard{
		BoardNumber: board.BoardNumber,
		PlayerID:    board.PlayerID.String(),
		Score:       board.Score,
		Coins:       board.Coins,
		Cells:       make([][]JsonCell, len(board.Cells)),
	}
}
