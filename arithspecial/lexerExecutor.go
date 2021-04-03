package arithspecial

import (
	"errors"
	"reflect"
	"strconv"
	"unsafe"
)

const (
	_LEX_ERROR   = iota
	_LEX_CORRECT = iota
	_LEX_SKIP    = iota
)

/*
lexerFunction is the semantic function of the lexer.
*/
func lexerExecutor(rule int, start int, end int, data []byte, stack []int64, top *int) (int, uint16, error) {
	var symbolID uint16
	var err error
	controlState := _LEX_ERROR
	yytext := BytesToString((data)[start:end])
	//yytext := string((*data)[start:end])

	switch rule {
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
			var yyValue int64
			symbolID = NUMBER
			yyValue, err = strconv.ParseInt(yytext, 10, 64)
			if err == nil {
				stack[*top] = yyValue
				*top++
			}
			controlState = _LEX_CORRECT
		}
	case 5:
		{
			controlState = _LEX_SKIP
		}
	case 6:
		{
			controlState = _LEX_SKIP
		}
	case 7:
		{
			controlState = _LEX_ERROR
		}
	}
	if controlState == _LEX_ERROR {
		err = errors.New("lexerExecutor: found an invalid token")
	}
	if controlState == _LEX_CORRECT {
		return controlState, symbolID, err
	}

	return controlState, symbolID, err
}

func BytesToString(bytes []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}))
}
