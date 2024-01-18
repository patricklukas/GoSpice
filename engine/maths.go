package engine

// Basic integer functions
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

// Fast lsb and msb algorithms based on De Bruijn sequences
// adapted bitScanForward and bitScanReverse
// https://www.chessprogramming.org/BitScan
// Thanks @authors Kim Walisch, Mark Dickinson

const (
	DEBRUIJN = 0x03f79d71b4cb0a89
)

func lsb(b BB) int {
	return deBruijnLsbTable[((b&-b)*DEBRUIJN)>>58]
}

func msb(b BB) int {
	b |= b >> 1
	b |= b >> 2
	b |= b >> 4
	b |= b >> 8
	b |= b >> 16
	b |= b >> 32
	return deBruijnMsbTable[(b*DEBRUIJN)>>58]
}

var deBruijnLsbTable = [64]int{
	0, 1, 48, 2, 57, 49, 28, 3,
	61, 58, 50, 42, 38, 29, 17, 4,
	62, 55, 59, 36, 53, 51, 43, 22,
	45, 39, 33, 30, 24, 18, 12, 5,
	63, 47, 56, 27, 60, 41, 37, 16,
	54, 35, 52, 21, 44, 32, 23, 11,
	46, 26, 40, 15, 34, 20, 31, 10,
	25, 14, 19, 9, 13, 8, 7, 6,
}

var deBruijnMsbTable = [64]int{
	0, 47, 1, 56, 48, 27, 2, 60,
	57, 49, 41, 37, 28, 16, 3, 61,
	54, 58, 35, 52, 50, 42, 21, 44,
	38, 32, 29, 23, 17, 11, 4, 62,
	46, 55, 26, 59, 40, 36, 15, 53,
	34, 51, 20, 43, 31, 22, 10, 45,
	25, 39, 14, 33, 19, 30, 9, 24,
	13, 18, 8, 12, 7, 6, 5, 63,
}
