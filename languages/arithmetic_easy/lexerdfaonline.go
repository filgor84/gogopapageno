package arithmetic_easy

type lexerDfaBetter struct {
	states   []lexerDfaState
	CurState int
}

func newDfa(states []lexerDfaState) lexerDfaBetter {
	return lexerDfaBetter{states: states, CurState: 0}
}

func (l lexerDfaBetter) getCurrentState() lexerDfaState {
	return l.states[l.CurState]
}

func (l lexerDfaBetter) finalState() bool {
	return l.getCurrentState().IsFinal
}

func (l *lexerDfaBetter) nextState(char int) bool {
	next := l.getCurrentState().Transitions[char]
	if next != -1 {
		l.CurState = next
	}
	return next != -1
}
