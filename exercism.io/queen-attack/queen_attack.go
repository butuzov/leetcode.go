package queenattack

import (
	"errors"
	"math"
)

// BoardPosition is Chess board possition
type BoardPosition struct {
	Row    int // valid 1-8
	Column int // valid 1-8
}

// CanQueenAttack determine attack possibility based on Queens position
// on a chess board.
func CanQueenAttack(pos ...string) (bool, error) {

	var queens [2]BoardPosition

	for i := range queens {
		var ok bool
		queens[i], ok = newBoardPosition(pos[i])

		if !ok {
			return false, errors.New("Impossible Queen Possition")
		}
	}

	// Same position
	if queens[0].Row == queens[1].Row && queens[0].Column == queens[1].Column {
		return false, errors.New("Two Queens - one cell")
	}

	// Horizontal position or Vertical one
	if queens[0].Row == queens[1].Row || queens[0].Column == queens[1].Column {
		return true, nil
	}

	// Diagonal position
	if math.Abs(float64(queens[0].Row-queens[1].Row)) == math.Abs(float64(queens[0].Column-queens[1].Column)) {
		return true, nil
	}

	return false, nil
}

// newBoardPosition - generate new Board Position
func newBoardPosition(bp string) (BoardPosition, bool) {

	if len(bp) != 2 {
		return *new(BoardPosition), false
	}

	var row, column int

	// ascii magic - letter to index (a=0... etc)
	row = int(bp[0] - 97)
	if row > 7 {
		return *new(BoardPosition), false
	}

	column = int(bp[1] - 49)
	if column > 7 {
		return *new(BoardPosition), false
	}

	// Just to deal with HumanUndestandable indexes
	return BoardPosition{row + 1, column + 1}, true
}
