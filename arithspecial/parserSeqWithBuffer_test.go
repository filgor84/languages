package arithspecial

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func parseFileWithBuffer(fileName string, exp_out int64) error {
	data, err := ioutil.ReadFile(DATA_DIR + fileName)
	if err != nil {
		return err
	}
	stackData := make([]int64, STACKSIZE)
	topData := 0
	stackSymbol := make([]uint16, STACKSIZE)
	topSymbol := 0

	err = parseStringWithBuffer(data, stackData, stackSymbol, &topData, &topSymbol)
	if err != nil {
		return err
	}
	res := stackData[0]
	if res != exp_out {
		return fmt.Errorf("Expected: %d\nFound:%d", exp_out, res)
	}
	return nil
}

func parseStringWithBufferAndCheckStack(data []byte) (string, string, error) {
	stackData := make([]int64, STACKSIZE)
	topData := 0
	stackSymbol := make([]uint16, STACKSIZE)
	topSymbol := 0
	err := parseStringWithBuffer(data, stackData, stackSymbol, &topData, &topSymbol)
	if err != nil {
		return "", "", err
	}
	stackDataToS := printStackInt64(stackData, topData)
	stackSymbolToS := printStackSymbol(stackSymbol, topSymbol)
	return stackSymbolToS, stackDataToS, nil
}

func checkStringWithBufferParsed(myString string, exp_sym string, exp_values string) error {
	symbols, values, err := parseStringWithBufferAndCheckStack([]byte(myString))
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

func TestIncompleteStringWithBuffer(t *testing.T) {
	myString := "1)*2+(5*(5+6)"
	exp_symbols := "E_F_S_T\nRPAR\nTIMES\nE_F_S_T\nPLUS\nLPAR\nE_S_T"
	exp_values := "1\n2\n55"
	err := checkStringWithBufferParsed(myString, exp_symbols, exp_values)
	if err != nil {
		t.Error(err)
	}
}

func TestSampleIncStringWithBuffer(t *testing.T) {
	myString := "2) + 1 * 2 * 3 + 11 * 222 * 3333 * (1 + 2) + 1 * 2 * 3 + 11 * 222 * 3333 * (1 + 2) + 1 * 2 * 3 + 11 * 222 * 3333 * (1 + 2)"
	exp_symbols := "E_F_S_T\nRPAR\nPLUS\nE_S"
	exp_values := "2\n73252692"
	err := checkStringWithBufferParsed(myString, exp_symbols, exp_values)
	if err != nil {
		t.Error(err)
	}
}

func TestCompleteStringWithBuffer(t *testing.T) {
	myString := "5+5+5+5+5*5"
	exp_symbols := "E_S"
	exp_values := "45"
	err := checkStringWithBufferParsed(myString, exp_symbols, exp_values)
	if err != nil {
		t.Error(err)
	}
}

/*
func BenchmarkParseWithBuffer10MB(b *testing.B) {
	fileName := "data/10MB.txt"
	for i := 0; i < b.N; i++ {
		err := parseFileWithBuffer(fileName, TEN_MB)
		if err != nil {
			b.Error(err)
		}
	}

}
*/
