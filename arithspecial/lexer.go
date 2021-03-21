package arithspecial

import "errors"

func yyLex(data *[]byte, start int) (int, int, error) {
	state := 0
	for i := start; i < len(*data); i++ {
		nextState := dfaStates[state][(*data)[i]]
		if nextState == -1 {
			if isFinalState[state] {
				return i, ruleState[state], nil
			}
			return 0, 0, errors.New("Invalid Token")
		}
		state = nextState
	}
	if isFinalState[state] {
		return len(*data), ruleState[state], nil
	}
	return 0, 0, errors.New("Invalid Token at EOF")
}
