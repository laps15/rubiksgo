package main

import (
	"fmt"
)

const (
	red    byte = 'R'
	green  byte = 'G'
	orange byte = 'O'
	blue   byte = 'B'
	yellow byte = 'Y'
	white  byte = 'W'
)

const (
	front byte = 'F'
	right byte = 'R'
	back  byte = 'B'
	left  byte = 'L'
	upper byte = 'U'
	down  byte = 'D'
)

type EdgePiece struct {
	colors [2]byte
}

type CornerPiece struct {
	colors [3]byte
}

type Cube struct {
	edge_pieces   map[[2]byte]EdgePiece
	corner_pieces map[[3]byte]CornerPiece
}

func newSolvedCube() Cube {
	edge_pieces := make(map[[2]byte]EdgePiece)
	edge_pieces[[2]byte{upper, back}] = EdgePiece{[2]byte{yellow, orange}}
	edge_pieces[[2]byte{upper, left}] = EdgePiece{[2]byte{yellow, blue}}
	edge_pieces[[2]byte{upper, front}] = EdgePiece{[2]byte{yellow, red}}
	edge_pieces[[2]byte{upper, right}] = EdgePiece{[2]byte{yellow, green}}
	edge_pieces[[2]byte{left, front}] = EdgePiece{[2]byte{blue, red}}
	edge_pieces[[2]byte{front, right}] = EdgePiece{[2]byte{red, green}}
	edge_pieces[[2]byte{right, back}] = EdgePiece{[2]byte{green, orange}}
	edge_pieces[[2]byte{back, left}] = EdgePiece{[2]byte{orange, blue}}
	edge_pieces[[2]byte{down, back}] = EdgePiece{[2]byte{white, orange}}
	edge_pieces[[2]byte{down, left}] = EdgePiece{[2]byte{white, blue}}
	edge_pieces[[2]byte{down, front}] = EdgePiece{[2]byte{white, red}}
	edge_pieces[[2]byte{down, right}] = EdgePiece{[2]byte{white, green}}

	corner_pieces := make(map[[3]byte]CornerPiece)
	corner_pieces[[3]byte{upper, left, back}] = CornerPiece{[3]byte{yellow, blue, orange}}
	corner_pieces[[3]byte{upper, left, front}] = CornerPiece{[3]byte{yellow, blue, red}}
	corner_pieces[[3]byte{upper, right, front}] = CornerPiece{[3]byte{yellow, green, red}}
	corner_pieces[[3]byte{upper, right, back}] = CornerPiece{[3]byte{yellow, green, orange}}
	corner_pieces[[3]byte{down, left, back}] = CornerPiece{[3]byte{white, blue, orange}}
	corner_pieces[[3]byte{down, left, front}] = CornerPiece{[3]byte{white, blue, red}}
	corner_pieces[[3]byte{down, right, front}] = CornerPiece{[3]byte{white, green, red}}
	corner_pieces[[3]byte{down, right, back}] = CornerPiece{[3]byte{white, green, orange}}

	return Cube{edge_pieces, corner_pieces}
}

func toString(cube Cube) string {
	matrix := [6][3][3]byte{}

	matrix[0][0][0] = cube.corner_pieces[[3]byte{upper, left, front}].colors[2]
	matrix[0][0][1] = cube.edge_pieces[[2]byte{upper, front}].colors[1]
	matrix[0][0][2] = cube.corner_pieces[[3]byte{upper, right, front}].colors[2]

	matrix[0][1][0] = cube.edge_pieces[[2]byte{left, front}].colors[1]
	matrix[0][1][1] = red
	matrix[0][1][2] = cube.edge_pieces[[2]byte{front, right}].colors[0]

	matrix[0][2][0] = cube.corner_pieces[[3]byte{down, left, front}].colors[2]
	matrix[0][2][1] = cube.edge_pieces[[2]byte{down, front}].colors[1]
	matrix[0][2][2] = cube.corner_pieces[[3]byte{down, right, front}].colors[2]

	return "   " + string([]byte{matrix[0][0][0], matrix[0][0][1], matrix[0][0][2]}) + "\n" +
		"   " + string([]byte{matrix[0][1][0], matrix[0][1][1], matrix[0][1][2]}) + "\n" +
		"   " + string([]byte{matrix[0][2][0], matrix[0][2][1], matrix[0][2][2]}) + "\n"
}

func main() {
	cube := newSolvedCube()
	fmt.Printf("Current cube:\n" + toString(cube))
}
