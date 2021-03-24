package arithspecial

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
		go parseStringPara(
			&wg,
			data[splitPoint[i]:splitPoint[i+1]],
			stackOfStackData[(i*STACKSIZE):(i+1)*STACKSIZE],
			stackOfStackSymbol[(i*STACKSIZE):(i+1)*STACKSIZE],
			&stackDataTops[i],
			&stackSymbolTops[i],
		)

	}
	wg.Wait()
	symbolsLeft := make([]uint16, 100*STACKSIZE)
	stackDataFinal := make([]int64, 100*STACKSIZE)
	dataLeft := make([]int64, 100*STACKSIZE)
	symbolNumber := 0
	dataNumber := 0
	for i := 0; i < len(stackDataTops); i++ {
		dataNumber += stackDataTops[i]
	}
	//fmt.Println("Symbols left to right")

	var stackBase int
	for i := 0; i < threads; i++ {
		stackBase = i * STACKSIZE
		curStackPos := 0
		for curStackPos < stackSymbolTops[i] {
			symbolsLeft[symbolNumber] = stackOfStackSymbol[curStackPos+stackBase]
			//fmt.Println(tokenToString(symbolsLeft[symbolNumber]))
			symbolNumber++
			curStackPos++
		}
	}

	//fmt.Println("Data left to right:")

	topData := 0
	for i := 0; i < threads; i++ {
		stackBase = i * STACKSIZE
		curStackPos := 0
		for curStackPos < stackDataTops[i] {
			dataLeft[topData] = stackOfStackData[stackBase+curStackPos]
			topData++
			curStackPos++
		}

	}

	//fmt.Println(dataLeft[:topData])

	for i := 0; i < topData; i++ {
		stackDataFinal[i] = dataLeft[topData-i-1]
		//fmt.Println(stackDataFinal[i])
	}
	//fmt.Println("Stack Data final:")
	//fmt.Println(stackDataFinal[:topData])

	//fmt.Printf("\nTop data:%d\n", topData)

	finalSymbolStack := make([]uint16, STACKSIZE)
	nonTermToTerm(symbolsLeft, symbolNumber)
	//fmt.Println("Symbols left to right:")

	//fmt.Println(printStackSymbol(symbolsLeft, symbolNumber))
	//fmt.Printf("Top Symbol: %d\n", symbolNumber)
	//fmt.Println(printStackInt64(stackDataFinal, topData))
	//fmt.Printf("Top Data: %d\n", topData)
	var res int64

	res, err = parseFinal(symbolsLeft, symbolNumber, dataLeft, dataNumber, finalSymbolStack)

	return res, err
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

func nonTermToTerm(symbols []uint16, top int) {
	for i := 0; i < top; i++ {
		if !isTerminal(symbols[i]) {
			symbols[i] = NUMBER
		}
	}
}
