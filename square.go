package main

import (
	"fmt"
)

// using little endian rank file (LERF) mapping
// see https://www.chessprogramming.org/Square_Mapping_Considerations#LittleEndianRankFileMapping
// type Sq uint8

func SqToCoord(sq int) string {
	return fmt.Sprintf("%s%d", string(rune(FileIdx(sq))+'A'), 1+RankIdx(sq))
}

func CoordToSq(coord string) int {
	x := int(coord[0] - 'A')
	y := int(coord[1] - '1')
	return x + y*8
}

func RankIdx(sq int) int { return sq >> 3 }

func FileIdx(sq int) int { return sq & 7 }

func onBoard(sq int) bool { return 0 <= sq && sq <= 63 }
