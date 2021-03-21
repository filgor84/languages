package arithperfect

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func parseParaTestFile(threads int, fileName string, exp_res int64) error {
	var res int64
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	res, err = parseWhole(bytes, threads)
	if err != nil {
		return err
	} else {
		if res != exp_res {
			return fmt.Errorf("True res: %d\nObtained res: %d\n", exp_res, res)
		}
	}
	return nil
}

func BenchmarkPaParse1Mb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := parseTestFile("data/1MB.txt", ONE_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse1T10Mb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := parseParaTestFile(1, "data/10MB.txt", TEN_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse2T10Mb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := parseParaTestFile(2, "data/10MB.txt", TEN_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse4T10Mb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := parseParaTestFile(4, "data/10MB.txt", TEN_MB)
		if err != nil {
			b.Error(err)
		}
	}
}
