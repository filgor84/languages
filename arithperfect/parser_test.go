package arithperfect

import (
	"fmt"
	"io/ioutil"
	"testing"
)

const (
	SMALL     = 1 + 2*3*(4+5)
	ONE_MB    = (1*2*3 + 11*222*3333*(1+2)) * 25966
	TEN_MB    = (1*2*3 + 11*222*3333*(1+2)) * 257473
	TWENTY_MB = 12573726911544
)

func parseTestFile(fileName string, exp_res int64) error {
	var res int64
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	res, err = parseString(bytes)
	if err != nil {
		return err
	} else {
		if res != exp_res {
			return fmt.Errorf("True res: %d\nObtained res: %d\n", exp_res, res)
		}
	}
	return nil
}

func parseTestFileNoPtr(fileName string, exp_res int64) error {
	var res int64
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	res, err = parseStringNoPtr(bytes)
	if err != nil {
		return err
	} else {
		if res != exp_res {
			return fmt.Errorf("True res: %d\nObtained res: %d\n", exp_res, res)
		}
	}
	return nil
}

func BenchmarkParse1Mb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := parseTestFile("data/1MB.txt", ONE_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse10Mb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := parseTestFile("data/10MB.txt", TEN_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse20Mb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := parseTestFile("data/20MB.txt", TWENTY_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse223Mb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := parseTestFile("data/223MB.txt", 277749900)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse223MbNoPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := parseTestFileNoPtr("data/223MB.txt", 277749900)
		if err != nil {
			b.Error(err)
		}
	}
}
