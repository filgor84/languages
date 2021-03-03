package arithonline

import "errors"

const (
	STACKSIZE = 1024
)

//StackInt64 is a simple fixed size stack implementation for int64
type StackInt64 struct {
	data [1024]int64
	Top  int
}

func (s StackInt64) stackEmpty() bool {
	return s.Top == 0
}

func (s StackInt64) stackFull() bool {
	return s.Top == STACKSIZE
}

func (s *StackInt64) push(i int64) error {
	if !s.stackFull() {
		s.data[s.Top] = i
		s.Top++
		return nil
	}
	return errors.New("You tried to push data on a full stack")
}

func (s *StackInt64) pop() (int64, error) {
	if !s.stackEmpty() {
		s.Top--
		return s.data[s.Top], nil
	}
	return -1, errors.New("You tried to pop data on empty stack")
}
