package game

type Coord struct {
	X, Y int
}

type IBoard interface {
	GetWidth() int
	GetHeight() int
	IsInBounds(x, y int) bool
	GetPiece(x, y int) Piece
	PlacePiece(x, y int, piece Piece)
	IsFull() bool
	Clear()
}

// TicTacToeBoard concrete implementation of the tic-tac-toe game board
type Board struct {
	board  [][]Piece
	Width  int
	Height int
}

func NewBoard(width, height int) *Board {
	board := make([][]Piece, height)
	for i := range board {
		board[i] = make([]Piece, width)
	}
	return &Board{
		board:  board,
		Width:  width,
		Height: height,
	}
}

func (b *Board) GetWidth() int {
	return b.Width
}

func (b *Board) GetHeight() int {
	return b.Height
}

func (b *Board) IsInBounds(x, y int) bool {
	return x >= 0 && y >= 0 && x < b.GetWidth() && y < b.GetHeight()
}

func (b *Board) GetPiece(x, y int) Piece {
	return b.board[y][x]
}

func (b *Board) PlacePiece(x, y int, piece Piece) {
	b.board[y][x] = piece
}

func (b *Board) IsFull() bool {
	for col := 0; col < b.GetWidth(); col++ {
		for row := 0; row < b.GetHeight(); row++ {
			if b.GetPiece(col, row) == nil {
				return false
			}
		}
	}
	return true
}

func (b *Board) Clear() {
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			b.board[y][x] = nil
		}
	}
}
