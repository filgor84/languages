package arithonline

import (
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
	pStack.pushSymbol(Symbol{LPAR, _YIELDS_PREC})
	pStack.pushSymbol(Symbol{LPAR, _YIELDS_PREC})
	pStack.pushSymbol(Symbol{LPAR, _YIELDS_PREC})

	//pStack := parserStack{}
}

func TestPStackPopCandidateRule(t *testing.T) {
	//pStack := parserStack{}

}
