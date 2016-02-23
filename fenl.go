package fenl

import "strings"

const (
	WHITE = iota
	BLACK
)

func Compress(fen string) string {
	return ""
}

func encodeFlags(sideToMove, castling, enpassant string) string {
	var val byte

	if sideToMove == "w" {
		val |= (1 << 4)
	}
	if strings.Contains(castling, "K") {
		val |= (1 << 3)
	}
	if strings.Contains(castling, "Q") {
		val |= (1 << 2)
	}
	if strings.Contains(castling, "k") {
		val |= (1 << 1)
	}
	if strings.Contains(castling, "q") {
		val |= (1 << 0)
	}

	retStr := string(val + 65)
	// We don't need the row information since knowing
	// the side to move allows us to figure that out
	if enpassant != "-" {
		retStr += string(enpassant[0])
	}

	return retStr
}
