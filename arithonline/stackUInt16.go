package arithonline

import (
	"errors"
	"fmt"
)

//16StackUInt16 is a simple fixed size stack implementation for int64
type StackUInt16 struct {
	data [STACKSIZE]uint16
	Top  int
}

func (s StackUInt16) stackEmpty() bool {
	return s.Top == 0
}

func (s StackUInt16) stackFull() bool {
	return s.Top == STACKSIZE
}

func (s *StackUInt16) push(i uint16) error {
	if !s.stackFull() {
		s.data[s.Top] = i
		s.Top++
		return nil
	}
	return errors.New("StackUInt16.push(i int): You tried to push data on a full stack")
}

func (s StackUInt16) read(pos int) (uint16, error) {
	if pos < 0 || pos > s.Top {
		return 0, fmt.Errorf("StackUInt16.read(): invalid position %d", pos)
	}
	return s.data[pos], nil
}

func (s *StackUInt16) pop() (uint16, error) {
	if !s.stackEmpty() {
		s.Top--
		return s.data[s.Top], nil
	}
	return 0, errors.New("StackUInt16.pop():You tried to pop data on empty stack")
}
