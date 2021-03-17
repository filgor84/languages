package arithperfect

import (
	"errors"
	"sync"
)

func parseWhole(data []byte, threads int) error {
	var err error
	stackOfStackSymbol := make([]uint16, threads*STACKSIZE)
	stackSymbolTops := make([]int, threads)
	stackOfStackData := make([]int64, threads*STACKSIZE)
	stackDataTops := make([]int, threads)
	splitPoint := make([]int, threads+1)
	var startSplit int
	var wg sync.WaitGroup
	splitPoint[0] = 0
	splitPoint[threads] = len(data)
	for i := 1; i < threads; i++ {
		startSplit = (len(data) * i) / threads
		splitPoint[i], err = findSplitPoint(&data, startSplit)
		if err != nil {
			return err
		}
	}
	for i := 0; i < threads; i++ {
		go parseSequenceSync(&wg, data[splitPoint[i]:splitPoint[i+1]], stackOfStackData[(i*STACKSIZE):(i+1)*STACKSIZE])
	}

}

func findSplitPoint(dataPtr *[]byte, start int) (int, error) {
	if start > len(*dataPtr) {
		return -1, errors.New("findSplitPoint: start from irregular position")
	}
	for i := start; i < len(*dataPtr); i++ {
		if (*dataPtr)[i] == '+' {
			return i, nil
		}

	}
	return -1, errors.New("Split point not found")
}

func allocateStacks(stackOfStacks *[][]uint16, nStacks int) {
	for i := 0; i < nStacks; i++ {
		(*stackOfStacks)[i] = make([]uint16, STACKSIZE)
	}
}
