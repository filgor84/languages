package arithonline

import (
	"testing"
)

func dummyArith(rule Rule, mem *memory) int {
	switch rule.rule {
	case 0:
		{
			//*genSym = symbol{LPAR, 0, nil, nil, nil}
			(*mem).save("LPAR")
			return LEX_CORRECT
		}
	case 1:
		{
			//*genSym = symbol{RPAR, 0, nil, nil, nil}
			(*mem).save("RPAR")
			return LEX_CORRECT
		}
	case 2:
		{
			//*genSym = symbol{TIMES, 0, nil, nil, nil}
			(*mem).save("TIMES")
			return LEX_CORRECT
		}
	case 3:
		{
			(*mem).save("PLUS")
			//*genSym = symbol{PLUS, 0, nil, nil, nil}
			return LEX_CORRECT
		}
	case 4:
		{
			//num := lexerInt64Pools[thread].Get()
			err := error(nil)
			//var num int64
			//num, err = strconv.ParseInt(yytext, 10, 64)
			if err != nil {
				return LEX_ERROR
			}
			(*mem).save("Number: " + rule.yytext)
			//*genSym = symbol{NUMBER, 0, num, nil, nil}
			return LEX_CORRECT
		}
	case 5:
		{
			return LEX_SKIP
		}
	case 6:
		{
			return LEX_SKIP
		}
	case 7:
		{
			return LEX_ERROR
		}
	}
	return LEX_ERROR
}

func TestDummyF(t *testing.T) {
	low := ListOfWords{}
	var mem memory = &low
	dummyArith(Rule{3, "+", 0, 0}, &mem)
	if low.toString() != "PLUS" {
		t.Error("Expected PLUS, returned ", low.toString())

	}

}

func TestLexer(t *testing.T) {
	low := ListOfWords{}
	var mem memory = &low
	testString := "(3 + 2) * 5"
	lexer := Lexer{
		dfaLanguage,
		[]byte(testString),
		0,
		mem,
		dummyArith}
	err := lexer.lex()
	if err != nil {
		t.Error("Lexer not working")
	}
	if mem.toString() != "cnar" {
		t.Error("Expected Risultato, found", mem.toString())
	}

}
