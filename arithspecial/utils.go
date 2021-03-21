package arithspecial

import (
	"strconv"
	"strings"
)

func printStackSymbol(stack []uint16, top int) string {
	var res []string
	for i := 0; i < top; i++ {
		res = append(res, tokenToString(stack[i]))
	}
	return strings.Join(res, "\n")
}

func printStackInt64(stack []int64, top int) string {
	var res []string
	for i := 0; i < top; i++ {
		res = append(res, strconv.FormatInt(stack[i], 10))
	}
	return strings.Join(res, "\n")
}
