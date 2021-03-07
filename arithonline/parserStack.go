package arithonline

import (
	"errors"
	"fmt"
)

type parserStack struct {
	stackSymbol      StackSymbol
	stackTerminalPos StackInt
	//stackTerminalPos record the position of the last terminal symbols pushed on StackSymbol
	stackYieldPPos StackInt
	//stackYieldTerminalPPos record the position on StackTerminalPos
	//which refers to the position of the last terminal symbol on Stac
}

func (p parserStack) isEmpty() bool {
	return p.stackSymbol.Top > 0
}

func (p parserStack) isFull() bool {
	return p.stackSymbol.Top < STACKSIZE
}

func (p parserStack) hasTerminal() bool {
	return p.stackTerminalPos.Top > 0
}

func (p parserStack) hasYieldPrec() bool {
	return p.stackYieldPPos.Top > 0
}

func (p *parserStack) pushSymbol(s Symbol) error {

	err := p.stackSymbol.push(s)
	if err != nil {
		return fmt.Errorf("parserStack.pushSymbol(Symbol): stackSymbol is full")
	}
	if isTerminal(s.symbolId) {
		err = p.stackTerminalPos.push(p.stackSymbol.Top - 1)
		if err != nil {
			return fmt.Errorf("parserStack.pushSymbol(Symbol): stackTerminalPos is full")
		}
		if s.precedence == _YIELDS_PREC {
			err = p.stackYieldPPos.push(p.stackTerminalPos.Top - 1)
			if err != nil {
				return fmt.Errorf("parserStack.pushSymbol(Symbol): stackYieldPPos is full")
			}
		}
	}
	return err
}

func (p *parserStack) getLastTerminalSymbol() (Symbol, error) {
	if p.hasTerminal() {
		pos := p.stackTerminalPos.data[p.stackTerminalPos.Top-1]
		return p.stackSymbol.read(pos)
	}

	return Symbol{}, errors.New("No last terminal")

}

func (p *parserStack) popCandidateRule() ([]uint16, error) {
	var res []Symbol
	tPos, err := p.stackYieldPPos.pop()
	if err != nil {
		return nil, err
	}
	sPos, err := p.stackTerminalPos.read(tPos)
	if err != nil {
		return nil, err
	}
	p.stackTerminalPos.Top = tPos
	if sPos == 0 || isTerminal(p.stackSymbol.data[sPos-1].symbolId) {
		res = p.stackSymbol.data[sPos:p.stackSymbol.Top]
		p.stackSymbol.Top = sPos
	} else {
		res = p.stackSymbol.data[sPos-1 : p.stackSymbol.Top]
		p.stackSymbol.Top = sPos - 1
	}

	var candRule = make([]uint16, 6)
	for i := 0; i < len(res); i++ {
		candRule[i] = res[i].symbolId
	}

	return candRule[0:len(res)], nil
}

/*Huge memory allocation
var candRule []uint16
for _, v := range res {
	candRule = append(candRule, v.symbolId)
}*/
