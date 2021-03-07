package arithonline

//ParseStringNew blabla
func ParseStringNew(data []byte) (int64, error) {
	dataMemory := StackInt64{}
	var stackSymbol StackUInt16
	//var stackPrecedence StackUInt16
	var stackTerminalPos StackInt
	var stackYieldPPos StackInt
	var lastTerminalSymbol uint16
	var tPos int
	var sPos int
	lexerReader := LexerReader{
		dfaLanguage,
		data,
		0}
	var inputNumber int
	var startRule int
	var x int64
	var y int64
	var output int64
	var hasOutput bool
	var precedence uint16
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
				if !stackTerminalPos.stackEmpty() {
					lastTerminalSymbol = stackSymbol.data[stackTerminalPos.data[stackTerminalPos.Top-1]]
					precedence = getPrecedence(lastTerminalSymbol, symbolID)
					//curSymbol = Symbol{symbolId, precedence}

					for precedence == _TAKES_PREC {
						tPos, err = stackYieldPPos.pop()
						sPos, err = stackTerminalPos.read(tPos)
						stackTerminalPos.Top = tPos

						if sPos == 0 || isTerminal(stackSymbol.data[sPos-1]) {
							startRule = sPos
						} else {
							startRule = sPos - 1
						}
						lhs, pRule := findMatch(stackSymbol.data[startRule:stackSymbol.Top])
						stackSymbol.Top = startRule
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
						stackSymbol.push(lhs)
						//stackPrecedence.push(_NO_PREC)

						//executefunction associated to candRule and push nonterminal
						if !stackTerminalPos.stackEmpty() {
							lastTerminalSymbol = stackSymbol.data[stackTerminalPos.data[stackTerminalPos.Top-1]]

							precedence = getPrecedence(lastTerminalSymbol, symbolID)
						} else {
							//just in case of sequential parsing or initial string, else NO_PREC
							precedence = _YIELDS_PREC
						}
					}
					//err = pStack.pushSymbol(Symbol{symbolID, precedence})
					err = stackSymbol.push(symbolID)
					//err = stackPrecedence.push(precedence)
					err = stackTerminalPos.push(stackSymbol.Top - 1)
					if precedence == _YIELDS_PREC {
						err = stackYieldPPos.push(stackTerminalPos.Top - 1)
					}
				} else {
					err = stackSymbol.push(symbolID)
					//err = stackPrecedence.push(_YIELDS_PREC)
					err = stackTerminalPos.push(stackSymbol.Top - 1)
					err = stackYieldPPos.push(stackTerminalPos.Top - 1)
				}
			} else {
				err = stackSymbol.push(symbolID)
			}

		}
		if lexerReader.eof() {
			//just in case of sequential parsing or ending string
			for !stackTerminalPos.stackEmpty() {
				tPos, err = stackYieldPPos.pop()
				sPos, err = stackTerminalPos.read(tPos)
				stackTerminalPos.Top = tPos

				if sPos == 0 || isTerminal(stackSymbol.data[sPos-1]) {
					startRule = sPos
				} else {
					startRule = sPos - 1
				}
				lhs, pRule := findMatch(stackSymbol.data[startRule:stackSymbol.Top])
				stackSymbol.Top = startRule
				inputNumber = rule2InputNumber[int(pRule)]
				if inputNumber == 2 {
					y, err = dataMemory.pop()
					x, err = dataMemory.pop()
				}
				hasOutput, output = parseExecutorNew(pRule, x, y)
				if hasOutput {
					err = dataMemory.push(output)
				}
				stackSymbol.push(lhs)
				//executefunction associated to candRule and push nonterminal
			}

		}

	}
	res, err = dataMemory.pop()
	return res, err
}
