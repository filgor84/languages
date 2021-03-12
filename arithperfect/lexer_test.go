package arithperfect

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func printToken(data *[]byte, start int, end int, rule int) string {
	return fmt.Sprintf("%s: %s", (*data)[start:end], dfaToString[rule])

}

func lexString(data []byte) error {
	start := 0
	end := 0
	//rule := 0
	//var dataMemory StackInt64
	var err error
	for start < len(data) {
		end, _, err = yyLex(&data, start)
		//lexerExecutor(rule, start, end, &data, &dataMemory)
		//dataMemory.top = 0
		if err != nil {
			return errors.New("yyLex failed")
		}
		start = end
	}
	return nil
}

func lexStringWithExec(data []byte) error {
	start := 0
	end := 0
	rule := 0
	var dataMemory = make([]int64, STACKSIZE)
	var dmemTop int
	var err error
	for start < len(data) {
		end, rule, err = yyLex(&data, start)
		lexerExecutor(rule, start, end, &data, &dataMemory, &dmemTop)
		dmemTop = 0
		if err != nil {
			return errors.New("yyLex failed")
		}
		start = end
	}
	return nil
}

func lexFile(fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("File %s not found", fileName)
	}
	err = lexString(data)
	return err
}

func lexFileWithExec(fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("File %s not found", fileName)
	}
	err = lexStringWithExec(data)
	return err
}

func TestSimpleExp(t *testing.T) {
	data := []byte("35 + 27")
	start := 0
	end := 0
	rule := -1
	exp := "35: NUMBER\n : SKIP\n+: PLUS\n : SKIP\n27: NUMBER"
	var res []string
	var err error
	for start < len(data) {
		end, rule, err = yyLex(&data, start)
		if err != nil {
			t.Log(err)
		}
		res = append(res, printToken(&data, start, end, rule))
		start = end
	}
	if strings.Join(res, "\n") != exp {
		t.Errorf("Expected:\n%s\n Found:\n%s", exp, strings.Join(res, "\n"))
	}

}

func BenchmarkJustRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ioutil.ReadFile("data/small.txt")
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkLexSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lexFile("data/small.txt")
	}
}

func BenchmarkLex1Mb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lexFile("data/1MB.txt")
	}
}

func BenchmarkLex10Mb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lexFile("data/10MB.txt")
	}
}

func BenchmarkLex20Mb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lexFile("data/20MB.txt")
	}
}

func BenchmarkLex1MbWithExec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lexFileWithExec("data/1MB.txt")
	}
}

func BenchmarkLex10MbWithExec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lexFileWithExec("data/10MB.txt")
	}
}

func BenchmarkLex20MbWithExec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lexFileWithExec("data/20MB.txt")
	}
}
