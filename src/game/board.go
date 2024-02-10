package game

import (
	"fmt"
)

type IBoard interface {
	GetWidth() int
	GetHeight() int
	GetPiece(x, y int) Piece
	PlacePiece(x, y int, piece Piece)
	Clear()
	Print()
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

func (b *Board) GetPiece(x, y int) Piece {
	return b.board[y][x]
}

func (b *Board) PlacePiece(x, y int, piece Piece) {
	b.board[y][x] = piece
}

func (b *Board) Clear() {
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			b.board[y][x] = Empty
		}
	}
}

func (b *Board) Print() {
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			switch b.board[y][x] {
			case Empty:
				fmt.Print("- ")
			case X:
				fmt.Print("x ")
			case O:
				fmt.Print("o ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
