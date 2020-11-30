package models

import (
	"fmt"
	"log"
	"strconv"
)

var xCoordinatesMap = map[string]int{
	"a": 0,
	"b": 1,
	"c": 2,
	"d": 3,
	"e": 4,
	"f": 5,
	"g": 6,
	"h": 7,
}

type Battlefield struct {
	isWhiteTurn bool
	field [][]*Figure
}

func NewBattlefield() *Battlefield {
	return &Battlefield{
		isWhiteTurn: true,
		field: [][]*Figure{
			/* 1 */ { WhiteRock, WhiteKnight, WhiteBishop, WhiteQueen, WhiteKing, WhiteBishop, WhiteKnight, WhiteRock},
			/* 2 */ { WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn},
			/* 3 */ { nil, nil, nil, nil, nil, nil, nil, nil},
			/* 4 */ { nil, nil, nil, nil, nil, nil, nil, nil},
			/* 5 */ { nil, nil, nil, nil, nil, nil, nil, nil},
			/* 6 */ { nil, nil, nil, nil, nil, nil, nil, nil},
			/* 7 */ { BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn},
			/* 8 */ { BlackRock, BlackKing, BlackBishop, BlackQueen, BlackKing, BlackBishop, BlackKnight, BlackRock},
			        // a    b    c    d    e    f    g    h
		},
	}
}

func (b *Battlefield) Move(fromCoordinates string, toCoordinates string) error {
	fromX, fromY := b.coordinatesToIndex(fromCoordinates)
	toX, toY := b.coordinatesToIndex(toCoordinates)

	figure := b.field[fromY][fromX]
	if figure == nil {
		return fmt.Errorf("no figure at: %s (%d, %d)", fromCoordinates, fromX, fromY)
	}

	if b.isWhiteTurn && figure.GetColor() != "White" {
		return fmt.Errorf("%s its not you turn", figure.GetColor())
	}

	canMove, declineReason := b.canFigureMakeTurn(figure, fromX, fromY, toX, toY)
	if !canMove {
		return fmt.Errorf("%s", declineReason)
	}

	b.field[fromY][fromX] = nil
	b.field[toY][toX] = figure
	b.isWhiteTurn = !b.isWhiteTurn
	return nil
}

func (b *Battlefield) canFigureMakeTurn(figure *Figure, fromX, fromY, toX, toY int) (bool, string) {
	deltaX := toX - fromX
	deltaY := toY - fromY
	log.Printf("DELTA: {%d, %d}\n", deltaX, deltaY)
	isFigureCanMakeTurn, turnType := b.isTurnPossible(figure, deltaX, deltaY)
	if !isFigureCanMakeTurn {
		return false, "figure cant move like this"
	}

	toCell := b.field[toY][toX]
	if toCell == nil && turnType != 0 && turnType != 1 {
		return false, "move turn not allowed"
	} else if toCell != nil && turnType != 0 && turnType != 2 {
		return false, "bite turn not allowed"
	}

	way := b.getWay(fromX, fromY, toX, toY)
	for turnIndex, turn := range way {
		isLastTurn := turnIndex < len(way)
		if b.field[turn[1]][turn[0]] != nil && !isLastTurn {
			return false, "way not clean"
		}
	}

	return true, ""
}

func (b *Battlefield) isTurnPossible(figure *Figure, deltaX, deltaY int) (bool, int) {
	for _, move := range figure.PossibleTurns() {
		if move[0] == deltaX && move[1] == deltaY {
			return true, move[2]
		}
	}
	return false, 0
}



func (b *Battlefield) getWay(fromX, fromY, toX, toY int) [][]int {
	log.Printf("Calc way (%d, %d) -> (%d, %d)\n", fromX, fromY, toX, toY)
	var turns [][]int
	curX := fromX
	curY := fromY
	for curX != toX || curY != toY {
		if curX != toX {
			if curX > toX {
				curX = curX - 1
			} else if curX < toX {
				curX = curX + 1
			}
		}
		if curY != toY {
			if curY > toY {
				curY = curY - 1
			} else if curY < toY {
				curY = curY + 1
			}
			turns = append(turns, []int{curX, curY})
		}
	}
	return turns
}

func (b *Battlefield) coordinatesToIndex(coordinates string) (int, int) {
	r := []rune(coordinates)

	x := xCoordinatesMap[string(r[0])]
	y, _ := strconv.ParseInt(string(r[1]), 10, 8)

	return x, int(y - 1)
}

func (b *Battlefield) Print() {
	fmt.Println("    a   b   c   d   e   f   g   h")
	fmt.Println("  |---|---|---|---|---|---|---|---|")
	for rowIndex, row := range b.field {
		for i := 0; i < 3; i++ {
			if i != 1 {
				fmt.Print("  ")
			} else {
				fmt.Printf("%d ", rowIndex+1)
			}
			for columnIndex, figure := range row {
				isEven := false
				if (rowIndex + columnIndex) % 2 == 0 {
					isEven = true
				}
				if i != 1 {
					if isEven {
						fmt.Print("|***")
					} else {
						fmt.Print("|   ")
					}
				} else {
					if isEven {
						symbol := "*"
						if figure != nil {
							symbol = figure.GetSymbol()
						}
						fmt.Printf("|*%s*", symbol)
					} else {
						symbol := " "
						if figure != nil {
							symbol = figure.GetSymbol()
						}
						fmt.Printf("| %s ", symbol)
					}
				}
			}
			fmt.Println("|")
		}
	}

}