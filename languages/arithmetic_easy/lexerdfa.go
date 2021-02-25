package arithmetic_easy

type lexerDfaState struct {
	Transitions     [256]int
	IsFinal         bool
	AssociatedRules []int
}

func newDfaState(transitions []int, final bool, rules []int) lexerDfaState {
	var expTransitions [256]int
	for i := range expTransitions {
		if i < len(transitions) {
			expTransitions[i] = transitions[i]
		} else {
			expTransitions[i] = -1
		}
	}
	return lexerDfaState{Transitions: expTransitions, IsFinal: final, AssociatedRules: rules}
}

type lexerDfa []lexerDfaState
