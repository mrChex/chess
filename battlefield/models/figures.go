package models

var kingMoves = [][]int{
	{1,0,0},
	{1,1,0},
	{0,1,0},
	{-1,1,0},
	{-1,0,0},
	{-1,1,0},
	{0,-1,0},
	{1,-1,0},
}
var WhiteKing = newFigure("♔", "White", "King", false, kingMoves)
var BlackKing = newFigure("♚", "Black", "King", false, kingMoves)

var queenMoves = [][]int{
	// right
	{1,0,0}, {2,0,0}, {3,0,0}, {4,0,0}, {5,0,0}, {6,0,0}, {7,0,0},
	// left
	{-1,0,0}, {-2,0,0}, {-3,0,0}, {-4,0,0}, {-5,0,0}, {-6,0,0}, {-7,0,0},
	// up
	{0,1,0}, {0,2,0}, {0,3,0}, {0,4,0}, {0,5,0}, {0,6,0}, {0,7,0},
	// down
	{0,-1,0}, {0,-2,0}, {0,-3,0}, {0,-4,0}, {0,-5,0}, {0,-6,0}, {0,-7,0},
	// right up
	{1,1,0}, {2,2,0}, {3,3,0}, {4,4,0}, {5,5,0}, {6,6,0}, {7,7,0},
	// left up
	{-1,1,0}, {-2,2,0}, {-3,3,0}, {-4,4,0}, {-5,5,0}, {-6,6,0}, {-7,7,0},
	// right down
	{1,-1,0}, {2,-2,0}, {3,-3,0}, {4,-4,0}, {5,-5,0}, {6,-6,0}, {7,-7,0},
	// left down
	{-1,-1,0}, {-2,-2,0}, {-3,-3,0}, {-4,-4,0}, {-5,-5,0}, {-6,-6,0}, {-7,-7,0},
}
var WhiteQueen = newFigure("♕", "White", "Queen", false, queenMoves)
var BlackQueen = newFigure("♛", "Black", "Queen", false, queenMoves)

var rockMoves = [][]int{
	// right
	{1,0,0}, {2,0,0}, {3,0,0}, {4,0,0}, {5,0,0}, {6,0,0}, {7,0,0},
	// left
	{-1,0,0}, {-2,0,0}, {-3,0,0}, {-4,0,0}, {-5,0,0}, {-6,0,0}, {-7,0,0},
	// up
	{0,1,0}, {0,2,0}, {0,3,0}, {0,4,0}, {0,5,0}, {0,6,0}, {0,7,0},
	// down
	{0,-1,0}, {0,-2,0}, {0,-3,0}, {0,-4,0}, {0,-5,0}, {0,-6,0}, {0,-7,0},
}
var WhiteRock = newFigure("♖", "White", "Rock", false, rockMoves)
var BlackRock = newFigure("♜", "Black", "Rock", false, rockMoves)

var bishopMoves = [][]int{
	// right up
	{1,1,0}, {2,2,0}, {3,3,0}, {4,4,0}, {5,5,0}, {6,6,0}, {7,7,0},
	// left up
	{-1,1,0}, {-2,2,0}, {-3,3,0}, {-4,4,0}, {-5,5,0}, {-6,6,0}, {-7,7,0},
	// right down
	{1,-1,0}, {2,-2,0}, {3,-3,0}, {4,-4,0}, {5,-5,0}, {6,-6,0}, {7,-7,0},
	// left down
	{-1,-1,0}, {-2,-2,0}, {-3,-3,0}, {-4,-4,0}, {-5,-5,0}, {-6,-6,0}, {-7,-7,0},
}
var WhiteBishop = newFigure("♗", "White", "Bishop", false, bishopMoves)
var BlackBishop = newFigure("♝", "Black", "Bishop", false, bishopMoves)

var knightMoves = [][]int{
	{-1,2,0},
	{1,2,0},
	{2,1,0},
	{-2,1,0},
	{-2,-1,0},
	{-1,-2,0},
	{2,-1,0},
	{1,-2,0},
}
var WhiteKnight = newFigure("♘", "White", "Knight", true, knightMoves)
var BlackKnight = newFigure("♞", "Black", "Knight", true, knightMoves)

var WhitePawn = newFigure("♙", "White", "Pawn", false, [][]int{
	{0,1,1},
	{0,2,1},
	{-1,1,2},
	{1,1,2},
})
var BlackPawn = newFigure("♟", "Black", "Pawn", false, [][]int{
	{0,-1,1},
	{0,-2,1},
	{-1,-1,2},
	{1,-1,2},
})

