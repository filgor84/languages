package arithonline

import (
	"fmt"
	"strings"
	"testing"
)

func yyLexOneToken(t *testing.T, data string) (string, error) {
	lexerReader := LexerReader{
		dfaLanguage,
		[]byte(data),
		0}
	rule, yyData, err := lexerReader.yyLex()
	res := fmt.Sprintf("Rule: %s, yytext: %s, Position: %d", rule.token, yyData.yytext, lexerReader.Pos)
	return res, err
}

func yyLexWholeString(t *testing.T, data string) (string, error) {
	lexerReader := LexerReader{
		dfaLanguage,
		[]byte(data),
		0}
	var res []string
	for !lexerReader.eof() {
		//t.Logf("Position %d", lexerReader.Pos)
		rule, yyData, err := lexerReader.yyLex()
		if err != nil {
			return strings.Join(res, "\n"), err
		}
		res = append(res, fmt.Sprintf("Rule: %s, yytext: %s", rule.token, yyData.yytext))
	}
	return strings.Join(res, "\n"), nil
}

type TestValue struct {
	input    string
	expected string
}

func TestYYLex(t *testing.T) {
	var testValues []TestValue
	testValues = append(testValues, TestValue{"3+2", "Rule: NUMBER, yytext: 3, Position: 1"})
	testValues = append(testValues, TestValue{"324", "Rule: NUMBER, yytext: 324, Position: 3"})
	testValues = append(testValues, TestValue{"011", "Rule: NUMBER, yytext: 011, Position: 3"})
	testValues = append(testValues, TestValue{"+11", "Rule: PLUS, yytext: +, Position: 1"})

	for _, v := range testValues {
		res, err := yyLexOneToken(t, v.input)
		if err != nil {
			t.Error("Lexer not working")
			break
		}
		if res != v.expected {
			t.Errorf("Error!\nInput:\n%s\nExpected result:\n%s\nFound:\n%s\n", v.input, v.expected, res)
		}
	}

}

func TestYYLexWholeString(t *testing.T) {
	var testValues []TestValue
	testValues = append(testValues, TestValue{"324", "Rule: NUMBER, yytext: 324"})
	testValues = append(testValues, TestValue{"011", "Rule: NUMBER, yytext: 011"})
	testValues = append(testValues, TestValue{"3+2",
		`Rule: NUMBER, yytext: 3
Rule: PLUS, yytext: +
Rule: NUMBER, yytext: 2`})
	testValues = append(testValues, TestValue{"(32 + 25) * 15",
		`Rule: LPAR, yytext: (
Rule: NUMBER, yytext: 32
Rule: SKIP, yytext:  
Rule: PLUS, yytext: +
Rule: SKIP, yytext:  
Rule: NUMBER, yytext: 25
Rule: RPAR, yytext: )
Rule: SKIP, yytext:  
Rule: TIMES, yytext: *
Rule: SKIP, yytext:  
Rule: NUMBER, yytext: 15`})

	for _, v := range testValues {
		res, err := yyLexWholeString(t, v.input)
		if err != nil {
			t.Errorf("Lexer not working\nError: %s\nPartial res:\n%s", err, res)
			break
		}
		if res != v.expected {
			t.Errorf("Error!\nInput:\n%s\nExpected result:\n%s\nFound:\n%s\n", v.input, v.expected, res)
		}
	}

}

/*
func TestLexer(t *testing.T) {
	low := ListOfWords{}
	var mem memory = &low
	testString := "(3 + 2) * 5"
	lexer := LexerReader{
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
*/