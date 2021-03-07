package arithonline

import (
	"io/ioutil"
	"testing"
)

const (
	SMALL  = 1 + 2*3*(4+5)
	ONE_MB = (1*2*3 + 11*222*3333*(1+2)) * 25966
	TEN_MB = (1*2*3 + 11*222*3333*(1+2)) * 257473
)

func parserTestFile(b *testing.B, fileName string, exp_res int64) {
	var res int64
	for i := 0; i < b.N; i++ {
		bytes, err := ioutil.ReadFile(fileName)
		if err != nil {
			b.Errorf("Error reading %s", fileName)
		}
		res, err = ParseStringNew(bytes)
		if err != nil {
			b.Errorf("Unexpected error: %v", err)
		} else {
			if res != exp_res {
				b.Errorf("True res: %d\nExpected res: %d\n", exp_res, res)
			}
		}
	}

}

func BenchmarkSmall(b *testing.B) {
	parserTestFile(b, "data/small.txt", SMALL)
}

func Benchmark1MB(b *testing.B) {
	parserTestFile(b, "data/1MB.txt", ONE_MB)
}

func Benchmark10MB(b *testing.B) {
	parserTestFile(b, "data/10MB.txt", TEN_MB)

}
