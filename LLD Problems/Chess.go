package main

// PieceType Color ValidMoves => Interface => Piece
// ChessPiece{ TypePiece Color} = { King={Castling()} Queen Pawn={Enpassant(), Promotion(), DoubleMove()} Bishop Rook Knight }
// Board {[8][8]Piece} 	
// Chess 

type Color int

const (
	White Color = iota
	Black
)

// PieceType represents the type of a chess piece
type PieceType int

const (
	King PieceType = iota
	Queen
	Rook
	Bishop
	Knight
	Pawn
)
