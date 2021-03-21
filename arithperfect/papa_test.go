package arithperfect

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func papaTestFile(fileName string, exp_res int64) error {
	var res int64
	stackSymbol := make([]uint16, STACKSIZE)
	var topSymbol int
	stackData := make([]int64, STACKSIZE)
	var topData int
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	err = papaSequence(data, stackData, &topData, stackSymbol, &topSymbol, 0, true, true)
	if err != nil {
		return err
	} else {
		res = stackData[0]
		if res != exp_res {
			return fmt.Errorf("True res: %d\nObtained res: %d\n", exp_res, res)
		}
	}
	return nil
}

func papaString(data []byte, isStart bool, isEnd bool) error {
	stackSymbol := make([]uint16, STACKSIZE)
	var topSymbol int
	stackData := make([]int64, STACKSIZE)
	var topData int
	err := papaSequence(data, stackData, &topData, stackSymbol, &topSymbol, 0, isStart, isEnd)
	if err != nil {
		return err
	}
	return nil
}
func TestStringStart(t *testing.T) {
	data := []byte("4 * 5 + 3")
	err := papaString(data, true, false)
	if err != nil {
		t.Error(err)
	}
}

func TestStringEnd(t *testing.T) {
	data := []byte("+ 3 + 2 + 4 * 5")
	err := papaString(data, false, true)
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkPapa1Mb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := papaTestFile("data/1MB.txt", ONE_MB)
		if err != nil {
			b.Error(err)
		}
	}
}
