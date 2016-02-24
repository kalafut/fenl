WHITE = 0
BLACK = 1

START_POS = "rnbqkbnrppppppppEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEPPPPPPPPRNBQKBNR"

PIECE_MAP = {
    "K": 0, "Q": 1, "R": 2, "B": 3, "N": 4, "P": 5,
    "k": 6, "q": 7, "r": 8, "b": 9, "n": 10, "p": 11
}

def encode_flags(side_to_move, castling, enpassant):
    "Encode all flags into one byte"
    val = 0

    if side_to_move == "w":
        val |= (1 << 4)

    if "K" in castling:
        val |= (1 << 3)

    if "Q" in castling:
        val |= (1 << 2)

    if "k" in castling:
        val |= (1 << 1)

    if "q" in castling:
        val |= (1 << 0)

    ret_str = chr(val + 65)

	# We don't need the row information since knowing
	# the side to move allows us to figure that out
    if enpassant != "-":
        ret_str += enpassant[0]

    return ret_str
