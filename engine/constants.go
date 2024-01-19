package engine

const (
	PAWN   Piece = iota
	KNIGHT       // 1
	BISHOP       // 2
	ROOK         // 3
	QUEEN        // 4
	KING         // 5
	EMPTY        // 6
)

const (
	EMPTY_CHECKMASK BB = 0xFFFFFFFFFFFFFFFF
)

const (
	WHITE uint8 = iota
	BLACK       // 1
)

const (
	QUEENSIDE BB = 0x0f0f0f0f0f0f0f0f
	KINGSIDE  BB = 0xf0f0f0f0f0f0f0f0
)

func flipColor(color uint8) uint8 {
	return color ^ 1
}

// Thanks @ https://tearth.dev/bitboard-viewer/ :)
var rowMask = [8]BB{
	0x00000000000000ff,
	0x000000000000ff00,
	0x0000000000ff0000,
	0x00000000ff000000,
	0x000000ff00000000,
	0x0000ff0000000000,
	0x00ff000000000000,
	0xff00000000000000,
}

var colMask = [8]BB{
	0x0101010101010101,
	0x0202020202020202,
	0x0404040404040404,
	0x0808080808080808,
	0x1010101010101010,
	0x2020202020202020,
	0x4040404040404040,
	0x8080808080808080,
}

// Castling
var castlingMasks = [2][2]BB{{0x1c, 0x70}, {0x1c00000000000000, 0x7000000000000000}}
var kingToRookMasks = [2][2]BB{{0xe, 0x60}, {0xe00000000000000, 0x6000000000000000}}

// Clockwise from North
const (
	NO int = iota
	NE     // 1
	EA     // 2
	SE     // 3
	SO     // 4
	SW     // 5
	WE     // 6
	NW     // 7
)

const (
	A1 int = iota
	B1
	C1
	D1
	E1
	F1
	G1
	H1
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A8
	B8
	C8
	D8
	E8
	F8
	G8
	H8
	SQ_INVALID
)

const (
	A int = iota
	B     // 1
	C     // 2
	D     // 3
	E     // 4
	F     // 5
	G     // 6
	H     // 7
)
