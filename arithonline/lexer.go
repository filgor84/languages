package arithonline

import (
	"errors"
	"fmt"
)

//Rule represents the number and the
type Rule struct {
	rule  int
	token string
}

type yyData struct {
	yytext string
	start  int
	end    int
}

// Lexer is an abstraction to get
type Lexer struct {
	Automa lexerDfa
	data   []byte
	pos    int
	mem    memory
	action func(Rule, yyData, *memory) int
}

func (l Lexer) currentByte() byte {
	return l.data[l.pos]
}

func (l Lexer) eof() bool {
	return l.pos == len(l.data)-1
}

func (l *Lexer) yyLex() (Rule, yyData, error) {
	controlState := TOKEN_INCOMPLETE
	start := l.pos
	end := start
	for controlState == TOKEN_INCOMPLETE {
		isValidTransaction := l.Automa.nextState(int(l.currentByte()))
		if isValidTransaction && !l.eof() {
			controlState = TOKEN_INCOMPLETE
		} else if isValidTransaction && l.eof() && l.Automa.isFinal() {
			controlState = TOKEN_COMPLETE
			end = len(l.data)
		} else if isValidTransaction && l.eof() && !l.Automa.isFinal() {
			controlState = TOKEN_ERROR
		} else if !isValidTransaction && !l.eof() {
			controlState = TOKEN_COMPLETE
			end = l.pos - 1
			l.pos--
		} else if !isValidTransaction && l.eof() {
			controlState = TOKEN_ERROR
		}
		l.pos++
	}
	if controlState == TOKEN_ERROR {
		return Rule{}, yyData{}, errors.New("Error found during lexing")
	}
	rule := Rule{l.Automa.getRuleNumber(), "p"}
	yydata := yyData{string(l.data[start:end]), start, end}
	l.Automa.CurState = 0
	return rule, yydata, nil
}

func (l *Lexer) lex() error {
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
