package arithperfect

import "errors"

const (
	STACKSIZE = 1024
)

//StackInt64 is a simple fixed size stack implementation for int64
type StackInt64 struct {
	data [STACKSIZE]int64
	top  int
}

func (s StackInt64) stackEmpty() bool {
	return s.top == 0
}

func (s StackInt64) stackFull() bool {
	return s.top == STACKSIZE
}

func (s *StackInt64) push(i int64) error {
	//if !s.stackFull() {
	s.data[s.top] = i
	s.top++
	return nil

	//return errors.New("StackInt64.push(int):You tried to push data on a full stack")
}

func (s *StackInt64) pop() (int64, error) {
	if !s.stackEmpty() {
		s.top--
		return s.data[s.top], nil
	}
	return -1, errors.New("StackInt64.pop():You tried to pop data on empty stack")
}
