package models

type Figure struct {
	symbol string
	color string
	name string
	possibleTurns [][]int
	canJump bool
}

// possibleTurns := [][]int{XBias, YBias, TurnType}
// turnTypes:
//   0 - move and bite
//   1 - just move
//   2 - just bite
func newFigure(symbol string, color string, name string, canJump bool, possibleTurns [][]int) *Figure {
	return &Figure{
		symbol: symbol,
		color: color,
		name: name,
		canJump: canJump,
		possibleTurns: possibleTurns,
	}
}

func (f *Figure) GetSymbol() string {
	return f.symbol
}

func (f *Figure) GetColor() string {
	return f.color
}

func (f *Figure) GetName() string {
	return f.name
}

func (f *Figure) PossibleTurns() [][]int {
	return f.possibleTurns
}