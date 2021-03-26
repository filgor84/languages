package arithspecial

import "errors"

func yyLexWithBuffer(dataBuffer *[]byte, start int) (int, int, bool, error) {
	state := 0
	for i := start; i < len(*dataBuffer); i++ {
		nextState := dfaStates[state][(*dataBuffer)[i]]
		if nextState == -1 {
			if isFinalState[state] {
				return i, ruleState[state], false, nil
			}
			return 0, -1, false, errors.New("Invalid Token")
		}
		state = nextState
	}
	if isFinalState[state] {
		return len(*dataBuffer), ruleState[state], true, nil
	}
	return start, -1, true, nil
}
