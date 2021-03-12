package arithperfect

import "testing"

func BenchmarkCountdownTo1K(b *testing.B) {
	s := StackInt64{}
	for m := 0; m < b.N; m++ {
		for n := 0; n < 1000; n++ {
			for i := 0; i < STACKSIZE; i++ {
				err := s.push(int64(i))
				if err != nil {
					b.Error(err)
				}
			}
			for i := 0; i < STACKSIZE; i++ {
				res, err := s.pop()
				if err != nil {
					b.Error(err)
				}
				if res+int64(i) != 1023 {
					b.Errorf("Something wrong: %d + %d != 1023", res, i)
				}
			}
		}
	}
}
func TestCountdownTo1000(t *testing.T) {
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

/*
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

}*/
