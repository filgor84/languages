package arithspecial

func parseFinal(symbolsLeft []uint16, symNumber int, dataLeft []int64, dataNumber int, symbolStack []uint16) (int64, error) {
	memoryStack := make([]int64, STACKSIZE)
	dataCursor := 0
	var topMemory int
	yPrecStack := make([]int, STACKSIZE)
	var yTop int
	var symbol uint16
	var lastTerminal uint16
	var lastYieldPrecPos int
	var ruleStart int
	var newNTSymbol uint16
	var parsingRule uint16
	var precedence uint16
	topSymbol := 0
	for i := 0; i < symNumber; i++ {
		symbol = symbolsLeft[i]
		if symbol == NUMBER {
			memoryStack[topMemory] = dataLeft[dataCursor]
			dataCursor++
			topMemory++
		}
		precedence = _TAKES_PREC
		for precedence == _TAKES_PREC {
			//Find last terminal symbol on stack and calculate precedence
			if topSymbol == 0 || (topSymbol == 1 && !isTerminal(symbolStack[0])) {
				//if there are no terminals assign by default _YIELDS_PREC
				precedence = _YIELDS_PREC
			} else {
				//if there is a terminal (on the first two position, for sure) calculate precedence
				lastTerminal = symbolStack[topSymbol-1]
				if !isTerminal(lastTerminal) {
					lastTerminal = symbolStack[topSymbol-2]
				}
				precedence = getPrecedence(lastTerminal, symbol)
			}
			//Pop position of first yield prec terminal
			if precedence == _TAKES_PREC && yTop > 0 {
				lastYieldPrecPos = yPrecStack[yTop-1]
				yTop--

				if lastYieldPrecPos == 0 || isTerminal(symbolStack[lastYieldPrecPos-1]) {
					//if the last yield prec symbol is the first symbol on stack or the symbol before
					//that is a terminal, then the rule will start from last yield prec symbol
					ruleStart = lastYieldPrecPos

				} else {
					ruleStart = lastYieldPrecPos - 1
				}
				newNTSymbol, parsingRule = findMatch(symbolStack[ruleStart:topSymbol])
				if newNTSymbol != _EMPTY {
					//errors.New("Unrecognized rule")
					parserExecutor(parsingRule, &memoryStack, &topMemory)
					symbolStack[ruleStart] = newNTSymbol
					topSymbol = ruleStart + 1
					continue
				}
			} else if precedence == _TAKES_PREC && yTop == 0 {
				precedence = _NO_PREC
			}
		}
		//Push the symbol on the top of the stackSymbol
		//if yield precedence, push its position on the top of yPrecStack
		symbolStack[topSymbol] = symbol
		if precedence == _YIELDS_PREC {
			yPrecStack[yTop] = topSymbol
			yTop++
		}
		topSymbol++

	}
	for yTop > 0 {
		lastYieldPrecPos = yPrecStack[yTop-1]
		yTop--

		if lastYieldPrecPos == 0 || isTerminal(symbolStack[lastYieldPrecPos-1]) {
			//if the last yield prec symbol is the first symbol on stack or the symbol before
			//that is a terminal, then the rule will start from last yield prec symbol
			ruleStart = lastYieldPrecPos

		} else {
			ruleStart = lastYieldPrecPos - 1
		}
		newNTSymbol, parsingRule = findMatch(symbolStack[ruleStart:topSymbol])
		if newNTSymbol == _EMPTY {
			break
		}
		parserExecutor(parsingRule, &memoryStack, &topMemory)
		symbolStack[ruleStart] = newNTSymbol
		topSymbol = ruleStart + 1
	}
	return memoryStack[0], nil
}
