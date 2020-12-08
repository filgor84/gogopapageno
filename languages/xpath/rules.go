package xpath

/*
rule represents a grammar rule of the language. Its lhs is a single token while its rhs is a slice of tokens.
*/
type rule struct {
	lhs uint16
	rhs []uint16
}

/*
The maximum length of the rhs of a rule of the language
*/
const _MAX_RHS_LEN = 4

/*
The rules of the language. They are sorted by their rhs
*/
var _RULES = []rule{
	rule{NEW_AXIOM, []uint16{ELEM}},
	rule{ELEM, []uint16{ELEM, OPENCLOSETAG}},
	rule{ELEM, []uint16{ELEM, OPENTAG, ELEM, CLOSETAG}},
	rule{ELEM, []uint16{ELEM, OPENTAG, CLOSETAG}},
	rule{ELEM, []uint16{ELEM, TEXT}},
	rule{ELEM, []uint16{OPENCLOSETAG}},
	rule{ELEM, []uint16{OPENTAG, ELEM, CLOSETAG}},
	rule{ELEM, []uint16{OPENTAG, CLOSETAG}},
	rule{ELEM, []uint16{TEXT}},
}

var compressedTrie = []uint16{2, 0, 4, 0, 11, 32769, 44, 32770, 47, 32771, 65, 1, 0, 3, 32769, 20, 32770, 23, 32771, 41, 0, 1, 0, 2, 0, 2, 0, 30, 32768, 38, 2, 0, 1, 32768, 35, 0, 2, 0, 0, 3, 0, 0, 4, 0, 0, 5, 0, 2, 0, 2, 0, 54, 32768, 62, 2, 0, 1, 32768, 59, 0, 6, 0, 0, 7, 0, 0, 8, 0}

/*
findMatch tries to find a match for the rhs using the compressed trie above.
On success it returns the corresponding lhs and the rule number.
On failure it returns an error.
*/
func findMatch(rhs []uint16) (uint16, uint16) {
	pos := uint16(0)

	for _, key := range rhs {
		//Skip the value and rule num for each node (except the last)
		pos += 2
		numIndices := compressedTrie[pos]
		if numIndices == 0 {
			return _EMPTY, 0
		}
		pos++
		low := uint16(0)
		high := uint16(numIndices - 1)
		startPos := pos
		foundNext := false

		for low <= high {
			indexpos := low + (high-low)/2
			pos = startPos + indexpos*2
			curKey := compressedTrie[pos]

			if key < curKey {
				high = indexpos - 1
			} else if key > curKey {
				low = indexpos + 1
			} else {
				pos = compressedTrie[pos+1]
				foundNext = true
				break
			}
		}
		if !foundNext {
			return _EMPTY, 0
		}
	}

	return compressedTrie[pos], compressedTrie[pos+1]
}
