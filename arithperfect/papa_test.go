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
	err = papaString(data, &stackData, &topData, &stackSymbol, &topSymbol, 0, true, true)
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

func BenchmarkPapa1Mb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := papaTestFile("data/1MB.txt", ONE_MB)
		if err != nil {
			b.Error(err)
		}
	}
}
