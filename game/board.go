package game

import "fmt"

type Coord struct {
	X, Y int
}

type IBoard interface {
	GetWidth() int
	GetHeight() int
	IsInBounds(x, y int) bool
	GetPiece(x, y int) Piece
	PlacePiece(x, y int, piece Piece)
	MovePiece(fromX, fromY, toX, toY int)
	RemovePiece(x, y int)
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
	p := b.GetPiece(x, y)
	if p != nil {
		panic(fmt.Sprintf("Space at %v %v is already occupied", x, y))
	}
	b.board[y][x] = piece
}

func (b *Board) MovePiece(fromX, fromY, toX, toY int) {
	p := b.GetPiece(fromX, fromY)
	if p == nil {
		panic(fmt.Sprintf("No piece to move at %v %v", fromX, fromY))
	}
	b.RemovePiece(fromX, fromY)
	b.PlacePiece(toX, toY, p)
}

func (b *Board) RemovePiece(x, y int) {
	p := b.GetPiece(x, y)
	if p == nil {
		panic(fmt.Sprintf("No piece to remove at %v %v", x, y))
	}
	b.board[y][x] = nil
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
