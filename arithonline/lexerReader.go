package arithonline

import (
	"errors"
)

//LexerRule represents the number and the
type LexerRule struct {
	rule  int
	token string
}

type yyData struct {
	yytext string
	start  int
	end    int
}

// LexerReader is an abstraction to get
type LexerReader struct {
	Automa lexerDfa
	data   []byte
	Pos    int
	//mem    memory
	//action func(Rule, yyData, *memory) int
}

func (l LexerReader) currentByte() byte {
	return l.data[l.Pos]
}

func (l LexerReader) eof() bool {
	return l.Pos == len(l.data)
}

func (l *LexerReader) yyLex() (LexerRule, yyData, error) {
	controlState := TOKEN_INCOMPLETE
	start := l.Pos
	end := start
	for controlState == TOKEN_INCOMPLETE {
		if !l.eof() {
			isValidTransition := l.Automa.nextState(int(l.data[l.Pos]))
			{
				//fmt.Printf("Pos: %d, Start: %d\n", l.Pos, start)
				if isValidTransition {
					//Case 0: EOF not reached, char generates a valid transition
					controlState = TOKEN_INCOMPLETE
					l.Pos++
				} else if l.Pos != start && l.Automa.isFinal() {
					//Case 1: EOF not reached, char doesn't generate valid transition
					//but we just finished to lex a valid token
					controlState = TOKEN_COMPLETE
					end = l.Pos
				} else {
					//Case 3: EOF not reached, every other case invalid
					controlState = TOKEN_ERROR
				}
			}
		} else if l.Automa.isFinal() {
			//Case 4: reached EOF, last char generates a valid transition to a final state
			controlState = TOKEN_COMPLETE
			end = len(l.data)
		} else {
			//Case 5: reached EOF, every other case invalid
			controlState = TOKEN_ERROR
		}
	}
	if controlState == TOKEN_ERROR {
		return LexerRule{}, yyData{}, errors.New("Error found during lexing")
	}
	rule := LexerRule{l.Automa.getRuleNumber(), l.Automa.getCurrentState().TokenString}
	yydata := yyData{string(l.data[start:end]), start, end}
	l.Automa.CurState = 0
	return rule, yydata, nil
}

/*

func (l *LexerReader) lex() error {
	var err error
	rule := Rule{}
	yydata := yyData{}
	for {
		rule, yydata, err = l.yyLex()
		if err != nil {
			fmt.Println("Token not matched")
			break
		}
		actionCode := l.action(rule, yydata, &l.mem)
		if actionCode == LEX_ERROR {
			err = fmt.Errorf("Something wrong during lexer rule application")
			break
		}
		if l.eof() {
			break
		}
	}
	return err
}
*/
