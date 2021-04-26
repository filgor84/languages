package arithspecial

import "testing"

func TestMatchNumber(t *testing.T) {
	myString := "12345hvd"
	exp_end := 5
	exp_rule := 4
	end, rule := matchArith([]byte(myString))
	if exp_end != end {
		t.Errorf("Wrong end\nFound: %d\nExpected: %d", end, exp_end)
	}
	if exp_rule != rule {
		t.Errorf("Wrong rule\nFound: %d\nExpected:%d", rule, exp_rule)
	}
}

func TestLPar(t *testing.T) {
	myString := "(5hvd"
	exp_end := 1
	exp_rule := 0
	end, rule := matchArith([]byte(myString))
	if exp_end != end {
		t.Errorf("Wrong end\nFound: %d\nExpected: %d", end, exp_end)
	}
	if exp_rule != rule {
		t.Errorf("Wrong rule\nFound: %d\nExpected:%d", rule, exp_rule)
	}
}

func TestRPar(t *testing.T) {
	myString := ")5hvd"
	exp_end := 1
	exp_rule := 1
	end, rule := matchArith([]byte(myString))
	if exp_end != end {
		t.Errorf("Wrong end\nFound: %d\nExpected: %d", end, exp_end)
	}
	if exp_rule != rule {
		t.Errorf("Wrong rule\nFound: %d\nExpected:%d", rule, exp_rule)
	}
}

func TestTimes(t *testing.T) {
	myString := "*5hvd"
	exp_end := 1
	exp_rule := 2
	end, rule := matchArith([]byte(myString))
	if exp_end != end {
		t.Errorf("Wrong end\nFound: %d\nExpected: %d", end, exp_end)
	}
	if exp_rule != rule {
		t.Errorf("Wrong rule\nFound: %d\nExpected:%d", rule, exp_rule)
	}
}

func TestPlus(t *testing.T) {
	myString := "+5hvd"
	exp_end := 1
	exp_rule := 3
	end, rule := matchArith([]byte(myString))
	if exp_end != end {
		t.Errorf("Wrong end\nFound: %d\nExpected: %d", end, exp_end)
	}
	if exp_rule != rule {
		t.Errorf("Wrong rule\nFound: %d\nExpected:%d", rule, exp_rule)
	}
}
