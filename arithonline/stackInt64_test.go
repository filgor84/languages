package arithonline

import "testing"

func TestCountdownTo10(t *testing.T) {
	s := StackInt64{}
	for i := 0; i < STACKSIZE; i++ {
		err := s.push(int64(i))
		if err != nil {
			t.Error(err)
		}
	}
	for i := 0; i < STACKSIZE; i++ {
		res, err := s.pop()
		if err != nil {
			t.Error(err)
		}
		if res+int64(i) != 1023 {
			t.Errorf("Something wrong: %d + %d != 1023", res, i)
		}
	}

}

func TestPopOnEmptyStack(t *testing.T) {
	s := StackInt64{}
	_, err := s.pop()
	if err == nil {
		t.Error("Popped on Empty Stack")
	}
}
func TestPushOnFullStack(t *testing.T) {
	s := StackInt64{}
	for i := 0; i < STACKSIZE; i++ {
		err := s.push(int64(i))
		if err != nil {
			t.Error(err)
		}
	}
	errLast := s.push(0)
	if errLast == nil {
		t.Error("Pushed on Full Stack")
	}

}
