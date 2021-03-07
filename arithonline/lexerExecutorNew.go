package arithonline

import (
	"strconv"
)

const (
	_LEX_CORRECT = 0
	_LEX_ERROR   = 1
	_SKIP        = 2
)

/*
lexerFunction is the semantic function of the lexer.
*/
func lexerExecutorNew(rule LexerRule, yydata yyData) (int, bool, int64, uint16, error) {
	var symbolID uint16
	var err error
	var yyValue int64
	hasValue := false
	controlState := LEX_ERROR

	switch rule.rule {
	case 0:
		{
			symbolID = LPAR
			controlState = _LEX_CORRECT
		}
	case 1:
		{
			symbolID = RPAR
			controlState = _LEX_CORRECT
		}
	case 2:
		{
			symbolID = TIMES
			controlState = _LEX_CORRECT
		}
	case 3:
		{
			symbolID = PLUS
			controlState = _LEX_CORRECT
		}
	case 4:
		{
			symbolID = NUMBER
			yyValue, err = strconv.ParseInt(yydata.yytext, 10, 64)
			hasValue = true
			controlState = _LEX_CORRECT
		}
	case 5:
		{
			controlState = _SKIP
		}
	case 6:
		{
			controlState = _SKIP
		}
	case 7:
		{
			controlState = _LEX_ERROR
		}
	}
	if err != nil {
		controlState = LEX_ERROR
	}
	if controlState == _LEX_CORRECT {
		return controlState, hasValue, yyValue, symbolID, err
	}

	return controlState, hasValue, yyValue, symbolID, err
}
