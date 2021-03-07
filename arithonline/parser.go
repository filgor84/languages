package arithonline

//ParseString blabla
func ParseString(data []byte) (int64, error) {
	dataMemory := StackInt64{}
	pStack := parserStack{}
	lexerReader := LexerReader{
		dfaLanguage,
		data,
		0}
	var inputNumber int
	var x int64
	var y int64
	var output int64
	var hasOutput bool
	var lastTerminal Symbol
	var precedence uint16
	var candRule []uint16
	var res int64
	var hasValue bool
	var lexerRule LexerRule
	var yydata yyData
	var err error
	var symbolID uint16
	var controlState int
	//user declared type
	var yyValue int64

	for !lexerReader.eof() {
		lexerRule, yydata, err = lexerReader.yyLex()
		if err != nil {
			return -1, err
		}
		controlState, hasValue, yyValue, symbolID, err = lexerExecutorNew(lexerRule, yydata)
		if hasValue {
			dataMemory.push(yyValue)
		}
		if controlState == _LEX_CORRECT {
			if isTerminal(symbolID) {
				if pStack.hasTerminal() {
					lastTerminal, err = pStack.getLastTerminalSymbol()
					precedence = getPrecedence(lastTerminal.symbolId, symbolID)
					//curSymbol = Symbol{symbolId, precedence}

					for precedence == _TAKES_PREC {
						candRule, err = pStack.popCandidateRule()
						lhs, pRule := findMatch(candRule)
						inputNumber = rule2InputNumber[int(pRule)]
						if inputNumber == 2 {
							y, err = dataMemory.pop()
							x, err = dataMemory.pop()
						}
						hasOutput, output = parseExecutorNew(pRule, x, y)
						if hasOutput {
							err = dataMemory.push(output)
						}

						//err = parseExecutor(pRule, &dataMemory)
						pStack.pushSymbol(Symbol{lhs, _NO_PREC})

						//executefunction associated to candRule and push nonterminal
						if pStack.hasTerminal() {
							lastTerminal, err = pStack.getLastTerminalSymbol()
							precedence = getPrecedence(lastTerminal.symbolId, symbolID)
						} else {
							//just in case of sequential parsing or initial string, else NO_PREC
							precedence = _YIELDS_PREC
						}
					}
					err = pStack.pushSymbol(Symbol{symbolID, precedence})
				} else {
					err = pStack.pushSymbol(Symbol{symbolID, _YIELDS_PREC})
				}
			} else {
				err = pStack.pushSymbol(Symbol{symbolID, _NO_PREC})
			}

		}
		if lexerReader.eof() {
			//just in case of sequential parsing or ending string
			for pStack.hasTerminal() {
				candRule, err = pStack.popCandidateRule()
				lhs, pRule := findMatch(candRule)
				inputNumber = rule2InputNumber[int(pRule)]
				if inputNumber == 2 {
					y, err = dataMemory.pop()
					x, err = dataMemory.pop()
				}
				hasOutput, output = parseExecutorNew(pRule, x, y)
				if hasOutput {
					err = dataMemory.push(output)
				}

				pStack.pushSymbol(Symbol{lhs, _NO_PREC})
				//executefunction associated to candRule and push nonterminal
			}

		}

	}
	res, err = dataMemory.pop()
	return res, err
}
