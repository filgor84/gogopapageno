package arithmetic_easy

import (
	"testing"
)

func TestDfaFromFirstToThird(t *testing.T) {
	dfa := newDfa([]lexerDfaState{

		newDfaState([]int{1, -1, -1}, false, []int{0}),
		newDfaState([]int{-1, 2, -1}, false, []int{0}),
		newDfaState([]int{-1, -1, -1}, false, []int{0}),
	})
	dfa.nextState(0)
	dfa.nextState(1)
	if dfa.CurState != 2 {
		t.Errorf("Expected final state: 2, found %d", dfa.CurState)
	}
}

func TestDfaWrongTransition(t *testing.T) {
	dfa := newDfa([]lexerDfaState{

		newDfaState([]int{1, -1, -1}, false, []int{0}),
		newDfaState([]int{-1, 2, -1}, false, []int{0}),
		newDfaState([]int{-1, -1, -1}, false, []int{0}),
	})
	dfa.nextState(0)
	if dfa.nextState(2) == true {
		t.Errorf("Expected invalid transition, found valid transition")
	}
}

func TestDfaWrongTransitionLastState(t *testing.T) {
	dfa := newDfa([]lexerDfaState{

		newDfaState([]int{1, -1, -1}, false, []int{0}),
		newDfaState([]int{-1, 2, -1}, false, []int{0}),
		newDfaState([]int{-1, -1, -1}, false, []int{0}),
	})
	dfa.nextState(0)
	dfa.nextState(2)
	if dfa.CurState != 1 {
		t.Errorf("Expected state 1, found state %d", dfa.CurState)
	}
}
