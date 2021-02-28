package arithonline

//possible outcomes:
//_TOKEN_INCOMPLETE char accepted, no symbol to create
//_TOKEN_COMPLETE char rejected, symbol ready to create
//LEX_FAILURE char rejected, no symbol ready

const (
	TOKEN_ERROR      = -1
	TOKEN_INCOMPLETE = 0
	TOKEN_COMPLETE   = 1
)

type lexerDfa struct {
	states   []lexerDfaState
	CurState int
}

func newDfa(states []lexerDfaState) lexerDfa {
	return lexerDfa{states: states, CurState: 0}
}

func (l lexerDfa) getCurrentState() lexerDfaState {
	return l.states[l.CurState]
}

func (l lexerDfa) isFinal() bool {
	return l.getCurrentState().IsFinal
}

func (l lexerDfa) getRuleNumber() int {
	return l.getCurrentState().AssociatedRules[0]
}

func (l *lexerDfa) nextState(char int) bool {
	next := l.getCurrentState().Transitions[char]
	if next != -1 {
		l.CurState = next
	}
	return next != -1
}

func (l *lexerDfa) nextChar(data int) (int, int) {
	isValidTransaction := l.nextState(data)
	if isValidTransaction {
		return -1, TOKEN_INCOMPLETE
	}
	if l.isFinal() {
		ruleNumber := l.getRuleNumber()
		l.CurState = 0
		return ruleNumber, TOKEN_COMPLETE
	}
	return -1, TOKEN_ERROR
}
