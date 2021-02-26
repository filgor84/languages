package arithmetic_easy

type Rune struct {
	Token uint16
	Value interface{}
}

/*
//Already defined in lexer.go
const (
	_ERROR       = -1
	_END_OF_FILE = 0
	_LEX_CORRECT = 1
	_SKIP        = 2
)
*/

type LexerOnline struct {
	Automa lexerDfaBetter
}

func (l *LexerOnline) nextChar(data byte, eof bool) (Rune, int) {
	isValid := l.Automa.nextState(int(data))
	if !isValid {
		//if (l.Automa.)

	}
	return _, 0
}
