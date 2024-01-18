package main

import (
	"fmt"
)

// using little endian rank file (LERF) mapping
// see https://www.chessprogramming.org/Square_Mapping_Considerations#LittleEndianRankFileMapping
type Sq uint8

func SqToCoord(sq Sq) string {
	return fmt.Sprintf("%s%d", string(rune(FileIdx(sq))+'A'), 1+RankIdx(sq))
}

func CoordToSq(coord string) Sq {
	x := int8(coord[0] - 'A')
	y := int8(coord[1] - '1')
	return Sq(x + y*8)
}

func RankIdx(sq Sq) int {
	return int(sq >> 3)
}

func FileIdx(sq Sq) int {
	return int(sq & 7)
}

func OnBoard(sq Sq) bool {
	return sq <= 63
}
