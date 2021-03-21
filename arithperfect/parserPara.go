package arithperfect

import (
	"errors"
	"sync"
)

func parseWhole(data []byte, threads int) (int64, error) {
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
		splitPoint[i], err = findSplitPoint(data, startSplit)
		if err != nil {
			return -1, err
		}
	}
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go papaSequenceSync(
			&wg,
			data[splitPoint[i]:splitPoint[i+1]],
			stackOfStackData[(i*STACKSIZE):(i+1)*STACKSIZE],
			&stackDataTops[i],
			stackOfStackSymbol[(i*STACKSIZE):(i+1)*STACKSIZE],
			&stackSymbolTops[i],
			splitPoint[i],
			true,
			true,
		)
	}
	wg.Wait()
	symbolsLeft := make([]uint16, STACKSIZE)
	dataLeft := make([]int64, STACKSIZE)
	topSymbol := 0
	topData := 0
	for i := 0; i < len(stackDataTops); i++ {
		topData += stackDataTops[i]
	}
	curData := topData - 1
	var stackBase int
	for i := 0; i < threads; i++ {
		stackBase = i * STACKSIZE
		for j := 0; j < stackSymbolTops[i]; j++ {
			symbolsLeft[topSymbol] = stackOfStackSymbol[stackBase+j]
			topSymbol++
		}
		for j := 0; j < stackDataTops[i]; j++ {
			dataLeft[curData] = stackOfStackData[stackBase+j]
			curData--
		}
	}
	finalSymbolStack := make([]uint16, STACKSIZE)
	topSymbolStack := 0

	parseSymbolData(symbolsLeft, dataLeft, &topData, finalSymbolStack, &topSymbolStack)

	return dataLeft[0], nil
}

func findSplitPoint(data []byte, start int) (int, error) {
	if start > len(data) {
		return -1, errors.New("findSplitPoint: start from irregular position")
	}
	for i := start; i < len(data); i++ {
		if data[i] == '+' {
			return i, nil
		}

	}
	return -1, errors.New("Split point not found")
}
