package arithonline

import (
	"strings"
	"testing"
)

func TestPStackPushNonTerminal(t *testing.T) {
	pStack := parserStack{}
	pStack.pushSymbol(Symbol{E_S, _NO_PREC})
	if pStack.hasTerminal() {
		t.Error("Pushed a nonterminal, but terminal are available")
	}
}

func TestPStackPushTerminal(t *testing.T) {
	pStack := parserStack{}
	pStack.pushSymbol(Symbol{LPAR, _NO_PREC})
	if !pStack.hasTerminal() {
		t.Error("Pushed a terminal, but  no terminal are available")
	}
	if pStack.hasYieldPrec() {
		t.Error("Pushed a terminal with NO_PREC, but  YIELD_PREC terminal are available")
	}
}

func TestPStackPushYieldPrecTerminal(t *testing.T) {
	pStack := parserStack{}
	pStack.pushSymbol(Symbol{LPAR, _YIELDS_PREC})
	if !pStack.hasTerminal() {
		t.Error("Pushed a terminal, but  no terminal are available")
	}
	if !pStack.hasYieldPrec() {
		t.Error("Pushed a terminal with YIELD_PREC, but  YIELD_PREC terminal are nots available")
	}

}

func TestPStackGetLastTerminal(t *testing.T) {
	pStack := parserStack{}
	pStack.pushSymbol(Symbol{RPAR, _YIELDS_PREC})
	pStack.pushSymbol(Symbol{E_S_T, _NO_PREC})
	pStack.pushSymbol(Symbol{LPAR, _YIELDS_PREC})
	pStack.pushSymbol(Symbol{E_S_T, _NO_PREC})
	s, err := pStack.getLastTerminalSymbol()
	if err != nil {
		t.Error(err)
	}
	if s.symbolId != LPAR {
		t.Errorf("Expected LPAR, found %s", tokenToString(s.symbolId))
	}
}

func TestPStackPopCandidateRule(t *testing.T) {
	pStack := parserStack{}
	pStack.pushSymbol(Symbol{RPAR, _YIELDS_PREC})
	pStack.pushSymbol(Symbol{PLUS, _NO_PREC})
	pStack.pushSymbol(Symbol{LPAR, _YIELDS_PREC})
	pStack.pushSymbol(Symbol{E_S_T, _NO_PREC})
	pStack.pushSymbol(Symbol{RPAR, _YIELDS_PREC})
	pStack.pushSymbol(Symbol{E_S_T, _NO_PREC})
	pStack.pushSymbol(Symbol{LPAR, _YIELDS_PREC})
	pStack.pushSymbol(Symbol{E_S_T, _NO_PREC})
	candRule, err := pStack.popCandidateRule()
	if err != nil {
		t.Error(err)
	}
	var res []string
	for _, v := range candRule {
		res = append(res, tokenToString(v))
	}

	if strings.Join(res, " ") != "E_S_T LPAR E_S_T" {
		t.Errorf("testPopCandRule(): Expected E_S_T LPAR E_S_T, found %s", strings.Join(res, " "))
	}

	candRule, err = pStack.popCandidateRule()

	if err != nil {
		t.Error(err)
	}
	var res2 []string
	for _, v := range candRule {
		res2 = append(res2, tokenToString(v))
	}

	if strings.Join(res2, " ") != "E_S_T RPAR" {
		t.Errorf("testPopCandRule(): Expected E_S_T RPAR, found %s", strings.Join(res, " "))
	}

	candRule, err = pStack.popCandidateRule()

	if err != nil {
		t.Error(err)
	}
	var res3 []string
	for _, v := range candRule {
		res3 = append(res3, tokenToString(v))
	}

	if strings.Join(res3, " ") != "LPAR" {
		t.Errorf("testPopCandRule(): Expected LPAR, found %s", strings.Join(res, " "))
	}

}
