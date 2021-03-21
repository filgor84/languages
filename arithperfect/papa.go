package arithperfect

import (
	"errors"
	"sync"
)

func papaSequence(data []byte, stackData []int64, topData *int, stackSymbol []uint16, topSymbol *int, offset int, isStart bool, isEnd bool) error {
	var err error
	//***Variables for lexing part
	var startLex int
	var endLex int
	var lexRule int
	controlState := _LEX_ERROR
	//***Variable for parsing part
	stackYPos := make([]int, STACKSIZE) //Helper stack to keep in memory last terminal
	//which yield precedenc
	var topYPos int //Top of stackYPos
	var symbol uint16
	var lastTerminal uint16
	var lastYieldPrecPos int
	var ruleStart int
	var newNTSymbol uint16
	var parsingRule uint16
	var precedence uint16
	for startLex < len(data) {
		endLex, lexRule, err = yyLex(&data, startLex)
		if err != nil {
			return err
		}
		controlState, symbol, err = lexerExecutor(lexRule, startLex, endLex, &data, &stackData, topData)
		if err != nil {
			return err
		}
		if controlState == _LEX_CORRECT {
			//Find last terminal symbol on stack and calculate precedence
			if *topSymbol == 0 || (*topSymbol == 1 && !isTerminal(stackSymbol[0])) {
				//if there are no terminals assign by default _YIELDS_PREC(seq parsing only!)
				if isStart {
					precedence = _YIELDS_PREC
				} else {
					precedence = _NO_PREC
				}
			} else {
				//if there is a terminal (on the first two position, for sure) calculate precedence
				lastTerminal = stackSymbol[*topSymbol-1]
				if !isTerminal(lastTerminal) {
					lastTerminal = stackSymbol[*topSymbol-2]
				}
				precedence = getPrecedence(lastTerminal, symbol)
			}

			for precedence == _TAKES_PREC && topYPos > 0 {
				//Pop position of first yield prec terminal
				lastYieldPrecPos = stackYPos[topYPos-1]
				topYPos--

				if lastYieldPrecPos == 0 || isTerminal(stackSymbol[lastYieldPrecPos-1]) {
					//if the last yield prec symbol is the first symbol on stack or the symbol before
					//that is a terminal, then the rule will start from last yield prec symbol
					ruleStart = lastYieldPrecPos

				} else {
					ruleStart = lastYieldPrecPos - 1
				}
				newNTSymbol, parsingRule = findMatch(stackSymbol[ruleStart:*topSymbol])
				if newNTSymbol == _EMPTY {
					return errors.New("Unrecognized rule")
				}
				parserExecutorNoPtr(parsingRule, stackData, topData)

				stackSymbol[ruleStart] = newNTSymbol
				*topSymbol = ruleStart + 1

				//Find last terminal symbol on stack and calculate precedence
				if *topSymbol == 1 && !isTerminal(stackSymbol[0]) {
					//if there are no terminals assign by default _YIELDS_PREC(seq parsing only!)
					if isStart {
						precedence = _YIELDS_PREC
					} else {
						precedence = _NO_PREC
					}
				} else {
					//if there is a terminal (on the first two position, for sure) calculate precedence
					lastTerminal = stackSymbol[*topSymbol-1]
					if !isTerminal(lastTerminal) {
						lastTerminal = stackSymbol[*topSymbol-2]
					}
					precedence = getPrecedence(lastTerminal, symbol)
				}

			}
			stackSymbol[*topSymbol] = symbol
			if precedence == _YIELDS_PREC {
				stackYPos[topYPos] = *topSymbol
				topYPos++
			}
			*topSymbol++
		}
		startLex = endLex
	}
	if isEnd {
		for topYPos > 0 {
			lastYieldPrecPos = stackYPos[topYPos-1]
			topYPos--

			if lastYieldPrecPos == 0 || isTerminal(stackSymbol[lastYieldPrecPos-1]) {
				//if the last yield prec symbol is the first symbol on stack or the symbol before
				//that is a terminal, then the rule will start from last yield prec symbol
				ruleStart = lastYieldPrecPos

			} else {
				ruleStart = lastYieldPrecPos - 1
			}
			newNTSymbol, parsingRule = findMatch(stackSymbol[ruleStart:*topSymbol])
			if newNTSymbol == _EMPTY {
				return errors.New("Unrecognized rule")
			}
			parserExecutorNoPtr(parsingRule, stackData, topData)

			stackSymbol[ruleStart] = newNTSymbol
			*topSymbol = ruleStart + 1
		}
	}
	/*
		if *topData != 1 {
			return errors.New("More than one value on the stack")
		}*/
	return nil
}

func papaSequenceSync(wg *sync.WaitGroup, data []byte, stackData []int64, topData *int, stackSymbol []uint16, topSymbol *int, offset int, isStart bool, isEnd bool) {
	defer wg.Done()
	papaSequence(data, stackData, topData, stackSymbol, topSymbol, offset, isStart, isEnd)
}