package arithspecial

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"testing"
)

const (
	SMALL          = 1 + 2*3*(4+5)
	ONE_MB         = (1*2*3 + 11*222*3333*(1+2)) * 25966
	TEN_MB         = (1*2*3 + 11*222*3333*(1+2)) * 257473
	TWENTY_MB      = 2 * TEN_MB
	ONE_HUNDRED_MB = 10 * TEN_MB
	DATA_DIR       = "/dev/shm/"
)

func paraParseFile(threads int, fileName string, exp_out int64) error {
	data, err := ioutil.ReadFile(DATA_DIR + fileName)
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

func BenchmarkParse10MB16T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(16, "data/10MB.txt", TEN_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse20MB1T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(1, "data/20MB.txt", TWENTY_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse20MB2T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(2, "data/20MB.txt", TWENTY_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse20MB4T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(4, "data/20MB.txt", TWENTY_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse20MB8T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(8, "data/20MB.txt", TWENTY_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse20MB16T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(16, "data/20MB.txt", TWENTY_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse20MB32T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(32, "data/20MB.txt", TWENTY_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse20MB64T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(64, "data/20MB.txt", TWENTY_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse20MB128T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(128, "data/20MB.txt", TWENTY_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse20MB256T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(256, "data/20MB.txt", TWENTY_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse20MB512T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(512, "data/20MB.txt", TWENTY_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse20MB1024T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(1024, "data/20MB.txt", TWENTY_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse20MB2048T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(2048, "data/20MB.txt", TWENTY_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse20MB4096T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(4096, "data/20MB.txt", TWENTY_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse20MB8192T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(8192, "data/20MB.txt", TWENTY_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse100MB1T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(1, "data/100MB.txt", ONE_HUNDRED_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse100MB8T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(8, "data/100MB.txt", ONE_HUNDRED_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse100MB128T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(128, "data/100MB.txt", ONE_HUNDRED_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse100MB256T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(256, "data/100MB.txt", ONE_HUNDRED_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse100MB512T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(512, "data/100MB.txt", ONE_HUNDRED_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse100MB1024T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(1024, "data/100MB.txt", ONE_HUNDRED_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse100MB2048T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(2048, "data/100MB.txt", ONE_HUNDRED_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse100MB4096T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(4096, "data/100MB.txt", ONE_HUNDRED_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse100MB8192T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(8192, "data/100MB.txt", ONE_HUNDRED_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse100MB16384T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(16384, "data/100MB.txt", ONE_HUNDRED_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse1GB1T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(1, "data/1GB.txt", 100*TEN_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse1GB2048T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(2048, "data/1GB.txt", 100*TEN_MB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParse1GB4096T(b *testing.B) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	for i := 0; i < b.N; i++ {
		err := paraParseFile(4096, "data/1GB.txt", 100*TEN_MB)
		if err != nil {
			b.Error(err)
		}
	}
}
