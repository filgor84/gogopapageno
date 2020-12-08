package xpath

const _NUM_NONTERMINALS = 3
const _NUM_TERMINALS = 5

const (
	ELEM         = iota
	NEW_AXIOM    = iota
	_EMPTY       = iota
	CLOSETAG     = 0x8000 + iota - _NUM_NONTERMINALS
	OPENCLOSETAG = 0x8000 + iota - _NUM_NONTERMINALS
	OPENTAG      = 0x8000 + iota - _NUM_NONTERMINALS
	TEXT         = 0x8000 + iota - _NUM_NONTERMINALS
	_TERM        = 0x8000 + iota - _NUM_NONTERMINALS
)

func tokenValue(token uint16) uint16 {
	return 0x7FFF & token
}

func isTerminal(token uint16) bool {
	return token >= 0x800
}

func tokenToString(token uint16) string {
	switch token {
	case ELEM:
		return "ELEM"
	case NEW_AXIOM:
		return "NEW_AXIOM"
	case _EMPTY:
		return "_EMPTY"
	case CLOSETAG:
		return "CLOSETAG"
	case OPENCLOSETAG:
		return "OPENCLOSETAG"
	case OPENTAG:
		return "OPENTAG"
	case TEXT:
		return "TEXT"
	case _TERM:
		return "_TERM"
	}
	return "UNKNOWN_TOKEN"
}
