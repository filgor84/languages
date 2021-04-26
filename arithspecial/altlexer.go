package arithspecial

import (
	"errors"
	"unicode/utf8"
)

const (
	TAB_B     = 9
	NEWLINE_B = 10
	SPACE_B   = 32
	LPAR_B    = 40
	RPAR_B    = 41
	TIMES_B   = 42
	PLUS_B    = 43
)

func matchArith(s []byte) (end int, rule int) {
	end = -1
	rule = -1
	var r rune
	var rlen int

	i := 0

	r, rlen = utf8.DecodeRune(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r == LPAR_B:
		end = i
		rule = 0
	case r == RPAR_B:
		end = i
		rule = 1
	case r == TIMES_B:
		end = i
		rule = 2
	case r == PLUS_B:
		end = i
		rule = 3
	case r >= 49 && r <= 57:
		end = i
		rule = 4
		goto s3
	case r == TAB_B || r == NEWLINE_B || r == SPACE_B:
		end = i
		rule = 5
	}
	return
s3:
	r, rlen = utf8.DecodeRune(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57:
		end = i
		goto s4
	}
	return
s4:
	r, rlen = utf8.DecodeRune(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57:
		end = i
		goto s4
	}
	return
}

func matchArithByte(s []byte) (end int, rule int) {
	end = -1
	rule = -1
	var b byte
	i := 0
	if i == len(s) {
		return
	}
	b = s[i]
	i += 1
	switch {
	case b == LPAR_B:
		end = i
		rule = 0
	case b == RPAR_B:
		end = i
		rule = 1
	case b == TIMES_B:
		end = i
		rule = 2
	case b == PLUS_B:
		end = i
		rule = 3
	case b >= 49 && b <= 57:
		end = i
		rule = 4
		goto s3
	case b == TAB_B || b == NEWLINE_B || b == SPACE_B:
		end = i
		rule = 5
	}
	return
s3:
	if i == len(s) {
		return
	}
	b = s[i]

	i += 1
	switch {
	case b >= 48 && b <= 57:
		end = i
		goto s4
	}
	return
s4:
	if i == len(s) {
		return
	}
	b = s[i]

	i += 1
	switch {
	case b >= 48 && b <= 57:
		end = i
		goto s4
	}
	return
}

func altLexer(data []byte, start int) (int, int, error) {
	end, rule := matchArithByte(data[start:])
	if rule == -1 {
		return -1, -1, errors.New("Illegal token")
	}
	return end + start, rule, nil
}
