package arithperfect

import "errors"

func papaString(data []byte) (int64, error) {
	startLex := 0
	endLex := 0
	rule := 0
	controlState := _LEX_ERROR
	memoryStack := make([]int64, STACKSIZE)
	var dmemTop int
	symbolStack := make([]uint16, STACKSIZE)
	var symTop int
	yPrecStack := make([]int, STACKSIZE)
	var yTop int
	var err error
	var symbol uint16
	var lastTerminal uint16
	var lastYieldPrecPos int
	var ruleStart int
	var newNTSymbol uint16
	var parsingRule uint16
	var precedence uint16
	for startLex < len(data) {
		endLex, rule, err = yyLex(&data, startLex)
		if err != nil {
			return -1, err
		}
		controlState, symbol, err = lexerExecutor(rule, startLex, endLex, &data, &memoryStack, &dmemTop)
		if err != nil {
			return -1, err
		}
		if controlState == _LEX_CORRECT {
			//Find last terminal symbol on stack and calculate precedence
			if symTop == 0 || (symTop == 1 && !isTerminal(symbolStack[0])) {
				//if there are no terminals assign by default _YIELDS_PREC(seq parsing only!)
				precedence = _YIELDS_PREC
			} else {
				//if there is a terminal (on the first two position, for sure) calculate precedence
				lastTerminal = symbolStack[symTop-1]
				if !isTerminal(lastTerminal) {
					lastTerminal = symbolStack[symTop-2]
				}
				precedence = getPrecedence(lastTerminal, symbol)
			}
			for precedence == _TAKES_PREC {
				//Pop position of first yield prec terminal
				lastYieldPrecPos = yPrecStack[yTop-1]
				yTop--

				if lastYieldPrecPos == 0 || isTerminal(symbolStack[lastYieldPrecPos-1]) {
					//if the last yield prec symbol is the first symbol on stack or the symbol before
					//that is a terminal, then the rule will start from last yield prec symbol
					ruleStart = lastYieldPrecPos

				} else {
					ruleStart = lastYieldPrecPos - 1
				}
				newNTSymbol, parsingRule = findMatch(symbolStack[ruleStart:symTop])
				if newNTSymbol == _EMPTY {
					return -1, errors.New("Unrecognized rule")
				}
				parserExecutor(parsingRule, &memoryStack, &dmemTop)

				symbolStack[ruleStart] = newNTSymbol
				symTop = ruleStart + 1

				//Find last terminal symbol on stack and calculate precedence
				if symTop == 1 && !isTerminal(symbolStack[0]) {
					//if there are no terminals assign by default _YIELDS_PREC(seq parsing only!)
					precedence = _YIELDS_PREC
				} else {
					//if there is a terminal (on the first two position, for sure) calculate precedence
					lastTerminal = symbolStack[symTop-1]
					if !isTerminal(lastTerminal) {
						lastTerminal = symbolStack[symTop-2]
					}
					precedence = getPrecedence(lastTerminal, symbol)
				}

			}
			symbolStack[symTop] = symbol
			if precedence == _YIELDS_PREC {
				yPrecStack[yTop] = symTop
				yTop++
			}
			symTop++
		}
		startLex = endLex
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
		newNTSymbol, parsingRule = findMatch(symbolStack[ruleStart:symTop])
		if newNTSymbol == _EMPTY {
			return -1, errors.New("Unrecognized rule")
		}
		parserExecutor(parsingRule, &memoryStack, &dmemTop)

		symbolStack[ruleStart] = newNTSymbol
		symTop = ruleStart + 1

	}
	if dmemTop != 1 {
		return -1, errors.New("More than one value on the stack")
	}
	return memoryStack[0], nil
}
