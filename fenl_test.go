package fenl

import (
	"testing"

	"gopkg.in/tylerb/is.v1"
)

func TestCompress(t *testing.T) {
	t.Skip()
	is := is.New(t)

	is.Equal(Compress("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"), "``")
	is.Equal(Compress("Rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"), "c_`")
	is.Equal(Compress("rnbqkbnr/pppppppp/8/8/8/4n3/PPPPPPPP/RNBQKBNR w KQkq - 0 1"), "Lk3`")

	is.Equal(Compress("1nbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"), "m_`")
	is.Equal(Compress("8/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"), "tX`")

	is.Equal(Compress("8/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w - - 0 1"), "tXQ")
	is.Equal(Compress("8/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b - - 0 1"), "tXA")
	is.Equal(Compress("8/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b K - 0 1"), "tXI")
	is.Equal(Compress("8/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b Q - 0 1"), "tXE")
	is.Equal(Compress("8/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b k - 0 1"), "tXC")
	is.Equal(Compress("8/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b q - 0 1"), "tXB")
	is.Equal(Compress("8/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b Kq - 0 1"), "tXJ")

	is.Equal(Compress("8/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b - a3 0 1"), "tXAa")
	is.Equal(Compress("8/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b - h3 0 1"), "tXAh")
	is.Equal(Compress("8/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq h3 0 1"), "tX`h")
	is.Equal(Compress("8/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq h3"), "tX`h")
}

func TestDecompress(t *testing.T) {
	t.Skip()
	is := is.New(t)
	fens := []string{
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 0",
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 0",
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b Kq - 0 0",
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b - - 0 0",
		"r1b1k2r/1pqnbppp/p2ppn2/6B1/3NPPP1/2N2Q2/PPP4P/2KR1B1R b kq - 0 0",
		"r1b1k2r/1pqnbppp/p2ppn2/6B1/3NPPP1/2N2Q2/PPP4P/2KR1B1R b kq a3 0 0",
		"r1b1k2r/1pqnbppp/p2ppn2/6B1/3NPPP1/2N2Q2/PPP4P/2KR1B1R b kq h3 0 0",
		"r1b1k2r/1pqnbppp/p2ppn2/6B1/3NPPP1/2N2Q2/PPP4P/2KR1B1R w kq a6 0 0",
		"r1b1k2r/1pqnbppp/p2ppn2/6B1/3NPPP1/2N2Q2/PPP4P/2KR1B1R w kq h6 0 0",
		"rnbq1rk1/p3ppbp/2p3p1/1p1nP3/3PN3/5N2/PPP1B1PP/R1BQ1RK1 b - - 0 0",
	}

	_, _ = fens, is
}

func TestEncodeFlags(t *testing.T) {
	is := is.New(t)

	is.Equal(encodeFlags("w", "KQkq", "-"), "`")
	is.Equal(encodeFlags("b", "Kq", "-"), "J")
	is.Equal(encodeFlags("b", "-", "h3"), "Ah")
}
