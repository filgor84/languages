package arithspecial

import (
	"fmt"
	"testing"
)

func parseStringAndCheckStack(data []byte) (string, string, error) {
	stackData := make([]int64, STACKSIZE)
	topData := 0
	stackSymbol := make([]uint16, STACKSIZE)
	topSymbol := 0
	err := parseString(data, stackData, stackSymbol, &topData, &topSymbol)
	if err != nil {
		return "", "", err
	}
	stackDataToS := printStackInt64(stackData, topData)
	stackSymbolToS := printStackSymbol(stackSymbol, topSymbol)
	return stackSymbolToS, stackDataToS, nil
}

func checkStringParsed(myString string, exp_sym string, exp_values string) error {
	symbols, values, err := parseStringAndCheckStack([]byte(myString))
	if err != nil {
		return err
	}
	if symbols != exp_sym {
		return fmt.Errorf("Expected:\n%s\nFound:\n%s", exp_sym, symbols)
	}

	if values != exp_values {
		return fmt.Errorf("Expected:\n%s\nFound:\n%s", exp_values, values)
	}
	return nil
}

func TestIncompleteString(t *testing.T) {
	myString := "1)*2+(5*(5+6)"
	exp_symbols := "E_F_S_T\nRPAR\nTIMES\nE_F_S_T\nPLUS\nLPAR\nE_S_T"
	exp_values := "1\n2\n55"
	err := checkStringParsed(myString, exp_symbols, exp_values)
	if err != nil {
		t.Error(err)
	}
}

func TestSampleIncString(t *testing.T) {
	myString := "2) + 1 * 2 * 3 + 11 * 222 * 3333 * (1 + 2) + 1 * 2 * 3 + 11 * 222 * 3333 * (1 + 2) + 1 * 2 * 3 + 11 * 222 * 3333 * (1 + 2)"
	exp_symbols := "E_F_S_T\nRPAR\nPLUS\nE_S"
	exp_values := "2\n73252692"
	err := checkStringParsed(myString, exp_symbols, exp_values)
	if err != nil {
		t.Error(err)
	}
}

func TestCompleteString(t *testing.T) {
	myString := "5+5+5+5+5*5"
	exp_symbols := "E_S"
	exp_values := "45"
	err := checkStringParsed(myString, exp_symbols, exp_values)
	if err != nil {
		t.Error(err)
	}
}
