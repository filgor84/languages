package arithonline

//ParseString blabla
func ParseString(data []byte) (int64, error) {
	dataMemory := StackInt64{}
	pStack := parserStack{}
	lexerReader := LexerReader{
		dfaLanguage,
		[]byte(data),
		0}
	var lastTerminal Symbol
	var precedence uint16
	var candRule []uint16
	var res int64
	var lexerRule LexerRule
	var yydata yyData
	var err error
	var symbolID uint16
	var controlState int

	for !lexerReader.eof() {
		lexerRule, yydata, err = lexerReader.yyLex()
		if err != nil {
			return -1, err
		}
		controlState, symbolID, err = lexerExecutor(lexerRule, yydata, &dataMemory)

		if controlState == LEX_CORRECT {
			if isTerminal(symbolID) {
				if pStack.hasTerminal() {
					lastTerminal, err = pStack.getLastTerminalSymbol()
					precedence = getPrecedence(lastTerminal.symbolId, symbolID)
					//curSymbol = Symbol{symbolId, precedence}

					for precedence == _TAKES_PREC {
						candRule, err = pStack.popCandidateRule()
						lhs, pRule := findMatch(candRule)
						err = parseExecutor(pRule, &dataMemory)
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
				err = parseExecutor(pRule, &dataMemory)
				pStack.pushSymbol(Symbol{lhs, _NO_PREC})
				//executefunction associated to candRule and push nonterminal
			}

		}

	}
	res, err = dataMemory.pop()
	return res, err
}
