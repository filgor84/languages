package arithspecial

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"testing"
)

const (
	SMALL     = 1 + 2*3*(4+5)
	ONE_MB    = (1*2*3 + 11*222*3333*(1+2)) * 25966
	TEN_MB    = (1*2*3 + 11*222*3333*(1+2)) * 257473
	TWENTY_MB = 12573726911544
)

func paraParseFile(threads int, fileName string, exp_out int64) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	res, err := parseWhole(data, threads)
	if err != nil {
		return err
	}
	if res != exp_out {
		return fmt.Errorf("Expected: %d\nFound:%d", exp_out, res)
	}
	return nil
}

func BenchmarkParse10MB1T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(1, "data/10MB.txt", TEN_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse10MB2T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(2, "data/10MB.txt", TEN_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse10MB4T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(4, "data/10MB.txt", TEN_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse10MB8T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(8, "data/10MB.txt", TEN_MB)
		if err != nil {
			b.Error(err)
		}
	}
}
