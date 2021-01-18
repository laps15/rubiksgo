package cube

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
	colors       [2]byte
	orientations [2]byte
}

type CornerPiece struct {
	colors       [3]byte
	orientations [3]byte
}

type Cube struct {
	edge_pieces   [12]EdgePiece
	corner_pieces [8]CornerPiece
}

func newSolvedCube() Cube {
	edge_pieces := [12]EdgePiece{
		EdgePiece{[2]byte{yellow, orange}, [2]byte{upper, back}},
		EdgePiece{[2]byte{yellow, blue}, [2]byte{upper, left}},
		EdgePiece{[2]byte{yellow, red}, [2]byte{upper, front}},
		EdgePiece{[2]byte{yellow, green}, [2]byte{upper, right}},
		EdgePiece{[2]byte{blue, red}, [2]byte{left, front}},
		EdgePiece{[2]byte{red, green}, [2]byte{front, right}},
		EdgePiece{[2]byte{green, orange}, [2]byte{right, back}},
		EdgePiece{[2]byte{orange, blue}, [2]byte{back, left}},
		EdgePiece{[2]byte{white, orange}, [2]byte{down, back}},
		EdgePiece{[2]byte{white, blue}, [2]byte{down, left}},
		EdgePiece{[2]byte{white, red}, [2]byte{down, front}},
		EdgePiece{[2]byte{white, green}, [2]byte{down, right}},
	}
	corner_pieces := [8]CornerPiece{
		CornerPiece{[3]byte{yellow, blue, orange}, [3]byte{upper, left, back}},
		CornerPiece{[3]byte{yellow, blue, red}, [3]byte{upper, left, front}},
		CornerPiece{[3]byte{yellow, green, red}, [3]byte{upper, right, front}},
		CornerPiece{[3]byte{yellow, green, orange}, [3]byte{upper, right, back}},
		CornerPiece{[3]byte{white, blue, orange}, [3]byte{down, left, back}},
		CornerPiece{[3]byte{white, blue, red}, [3]byte{down, left, front}},
		CornerPiece{[3]byte{white, green, red}, [3]byte{down, right, front}},
		CornerPiece{[3]byte{white, green, orange}, [3]byte{down, right, back}},
	}

	return Cube{edge_pieces, corner_pieces}
}

func toString() string {
	result := "   xxx\n   xxx\n   xxx\nxxxxxxxxxxxx\nxxxxxxxxxxxx\nxxxxxxxxxxxx\n   xxx\n   xxx\n   xxx\n"

	return result
}
