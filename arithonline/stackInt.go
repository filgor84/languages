package arithonline

import (
	"errors"
	"fmt"
)

//StackInt is a simple fixed size stack implementation for int64
type StackInt struct {
	data [STACKSIZE]int
	Top  int
}

func (s StackInt) stackEmpty() bool {
	return s.Top == 0
}

func (s StackInt) stackFull() bool {
	return s.Top == STACKSIZE
}

func (s *StackInt) push(i int) error {
	if !s.stackFull() {
		s.data[s.Top] = i
		s.Top++
		return nil
	}
	return errors.New("StackInt.push(i int): You tried to push data on a full stack")
}

func (s StackInt) read(pos int) (int, error) {
	if pos < 0 || pos > s.Top {
		return -1, fmt.Errorf("StackInt.read(): invalid position %d", pos)
	}
	return s.data[pos], nil
}

func (s *StackInt) pop() (int, error) {
	if !s.stackEmpty() {
		s.Top--
		return s.data[s.Top], nil
	}
	return 0, errors.New("StackInt.pop():You tried to pop data on empty stack")
}
