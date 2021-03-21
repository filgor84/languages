package arithspecial

import (
	"fmt"
	"sync"
)

const (
	STACKSIZE = 1024
)

func parseString(data []byte, memoryStack []int64, symbolStack []uint16, topMemory *int, topSymbol *int) error {
	startLex := 0
	endLex := 0
	rule := 0
	controlState := _LEX_ERROR
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
			return err
		}
		controlState, symbol, err = lexerExecutor(rule, startLex, endLex, &data, &memoryStack, topMemory)
		if err != nil {
			return err
		}
		if controlState == _LEX_CORRECT {
			precedence = _TAKES_PREC
			for precedence == _TAKES_PREC {
				//Find last terminal symbol on stack and calculate precedence
				if *topSymbol == 0 || (*topSymbol == 1 && !isTerminal(symbolStack[0])) {
					//if there are no terminals assign by default _YIELDS_PREC
					precedence = _YIELDS_PREC
				} else {
					//if there is a terminal (on the first two position, for sure) calculate precedence
					lastTerminal = symbolStack[*topSymbol-1]
					if !isTerminal(lastTerminal) {
						lastTerminal = symbolStack[*topSymbol-2]
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
					newNTSymbol, parsingRule = findMatch(symbolStack[ruleStart:*topSymbol])
					if newNTSymbol != _EMPTY {
						//errors.New("Unrecognized rule")
						parserExecutor(parsingRule, &memoryStack, topMemory)
						symbolStack[ruleStart] = newNTSymbol
						*topSymbol = ruleStart + 1
						continue
					}
				} else if precedence == _TAKES_PREC && yTop == 0 {
					//precedence = _NO_PREC
					precedence = _YIELDS_PREC

				}
			}
			//Push the symbol on the top of the stackSymbol
			//if yield precedence, push its position on the top of yPrecStack
			if *topSymbol > 1023 {
				fmt.Print("STOP!")
			}
			symbolStack[*topSymbol] = symbol
			if precedence == _YIELDS_PREC {
				if yTop > 1023 {
					fmt.Print("STOP!")
				}
				yPrecStack[yTop] = *topSymbol
				yTop++
			}
			*topSymbol++

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
		newNTSymbol, parsingRule = findMatch(symbolStack[ruleStart:*topSymbol])
		if newNTSymbol == _EMPTY {
			break
		}
		parserExecutor(parsingRule, &memoryStack, topMemory)
		symbolStack[ruleStart] = newNTSymbol
		*topSymbol = ruleStart + 1
	}
	return nil
}

func parseStringPara(
	wg *sync.WaitGroup,
	data []byte,
	memoryStack []int64,
	symbolStack []uint16,
	topMemory *int,
	topSymbol *int) {
	defer wg.Done()
	parseString(data, memoryStack, symbolStack, topMemory, topSymbol)
}
