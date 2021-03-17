package arithperfect

import (
	"testing"
)

func TestSplitData(t *testing.T) {
	dataString := []byte("3*3+4")
	end, err := findSplitPoint(&dataString, 0)
	if err != nil {
		t.Error(err)
	}
	if end != 3 {
		t.Errorf("TestSplitData: expected 3, found %d", end)
	}
}
