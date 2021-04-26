package arithspecial

import "testing"

func TestYYLexNumber(t *testing.T) {
	myString := "12345hvd"
	exp_end := 5
	exp_rule := 4
	end, rule, _ := yyLex([]byte(myString), 0)
	if exp_end != end {
		t.Errorf("Wrong end\nFound: %d\nExpected: %d", end, exp_end)
	}
	if exp_rule != rule {
		t.Errorf("Wrong rule\nFound: %d\nExpected:%d", rule, exp_rule)
	}
}

func TestYYLexNumberEnd(t *testing.T) {
	myString := "12345"
	exp_end := 5
	exp_rule := 4
	end, rule, _ := yyLex([]byte(myString), 0)
	if exp_end != end {
		t.Errorf("Wrong end\nFound: %d\nExpected: %d", end, exp_end)
	}
	if exp_rule != rule {
		t.Errorf("Wrong rule\nFound: %d\nExpected:%d", rule, exp_rule)
	}
}

func TestYYLexLPar(t *testing.T) {
	myString := "(2345"
	exp_end := 1
	exp_rule := 0
	end, rule, _ := yyLex([]byte(myString), 0)
	if exp_end != end {
		t.Errorf("Wrong end\nFound: %d\nExpected: %d", end, exp_end)
	}
	if exp_rule != rule {
		t.Errorf("Wrong rule\nFound: %d\nExpected:%d", rule, exp_rule)
	}
}

func TestYYLexNumberAfter(t *testing.T) {
	myString := "pvc12345hvd"
	exp_end := 8
	exp_rule := 4
	end, rule, _ := yyLex([]byte(myString), 3)
	if exp_end != end {
		t.Errorf("Wrong end\nFound: %d\nExpected: %d", end, exp_end)
	}
	if exp_rule != rule {
		t.Errorf("Wrong rule\nFound: %d\nExpected:%d", rule, exp_rule)
	}
}

func TestAltLexNumber(t *testing.T) {
	myString := "12345hvd"
	exp_end := 5
	exp_rule := 4
	end, rule, _ := altLexer([]byte(myString), 0)
	if exp_end != end {
		t.Errorf("Wrong end\nFound: %d\nExpected: %d", end, exp_end)
	}
	if exp_rule != rule {
		t.Errorf("Wrong rule\nFound: %d\nExpected:%d", rule, exp_rule)
	}
}

func TestAltLexNumberEnd(t *testing.T) {
	myString := "12345"
	exp_end := 5
	exp_rule := 4
	end, rule, _ := altLexer([]byte(myString), 0)
	if exp_end != end {
		t.Errorf("Wrong end\nFound: %d\nExpected: %d", end, exp_end)
	}
	if exp_rule != rule {
		t.Errorf("Wrong rule\nFound: %d\nExpected:%d", rule, exp_rule)
	}
}

func TestAltLexLPar(t *testing.T) {
	myString := "(2345"
	exp_end := 1
	exp_rule := 0
	end, rule, _ := altLexer([]byte(myString), 0)
	if exp_end != end {
		t.Errorf("Wrong end\nFound: %d\nExpected: %d", end, exp_end)
	}
	if exp_rule != rule {
		t.Errorf("Wrong rule\nFound: %d\nExpected:%d", rule, exp_rule)
	}
}

func TestAltLexNumberAfter(t *testing.T) {
	myString := "pvc12345hvd"
	exp_end := 8
	exp_rule := 4
	end, rule, _ := altLexer([]byte(myString), 3)
	if exp_end != end {
		t.Errorf("Wrong end\nFound: %d\nExpected: %d", end, exp_end)
	}
	if exp_rule != rule {
		t.Errorf("Wrong rule\nFound: %d\nExpected:%d", rule, exp_rule)
	}
}
