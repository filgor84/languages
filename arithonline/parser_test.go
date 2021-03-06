package arithonline

import (
	"testing"
)

func TestParseString(t *testing.T) {
	testString := "3 + 2 * (4 + 5)"
	res, err := ParseString([]byte(testString))
	if err != nil {
		t.Error(err)
	}
	if res != 21 {
		t.Errorf("Expected: 21 Found: %d", res)
	}

}
