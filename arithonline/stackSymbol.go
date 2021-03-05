package arithonline

import (
	"errors"
	"fmt"
)

//StackSymbol is a simple fixed size stack implementation for int64
type StackSymbol struct {
	data [1024]Symbol
	Top  int
}

func (s StackSymbol) stackEmpty() bool {
	return s.Top == 0
}

func (s StackSymbol) stackFull() bool {
	return s.Top == STACKSIZE
}

func (s *StackSymbol) push(i Symbol) error {
	if !s.stackFull() {
		s.data[s.Top] = i
		s.Top++
		return nil
	}
	return errors.New("StackSymbol.push(): you tried to push data on a full stack")
}

func (s StackSymbol) read(pos int) (Symbol, error) {
	if pos < 0 || pos > s.Top {
		return Symbol{}, fmt.Errorf("StackSymbol.read(): invalid position %d", pos)
	}
	return s.data[pos], nil
}

func (s StackSymbol) readTop() (Symbol, error) {
	if !s.stackEmpty() {
		return s.data[s.Top-1], nil
	}
	return Symbol{}, errors.New("StackSymbol.readTop(): StackSymbol is empty")
}
func (s *StackSymbol) pop() (Symbol, error) {
	if !s.stackEmpty() {
		s.Top--
		return s.data[s.Top], nil
	}
	return Symbol{}, errors.New("StackSymbol.pop(): you tried to pop a symbol from an empty stack")
}
