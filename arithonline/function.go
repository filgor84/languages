package arithonline

/*
function is the semantic function of the parser.
*/
//func function(thread int, ruleNum uint16, lhs *symbol, rhs []*symbol) {
func parseExecutor(ruleNum uint16, mem *StackInt64) error {
	var err error
	switch ruleNum {
	case 0:
		/*NEW_AXIOM0 := lhs
		E_F_S_T1 := rhs[0]

		NEW_AXIOM0.Child = E_F_S_T1

		{
			NEW_AXIOM0.Value = E_F_S_T1.Value
		}*/
		return nil
	case 1: /*
			E_S0 := lhs
			E_F_S_T1 := rhs[0]
			PLUS2 := rhs[1]
			E_F_S_T3 := rhs[2]

			E_S0.Child = E_F_S_T1
			E_F_S_T1.Next = PLUS2
			PLUS2.Next = E_F_S_T3

			{
				newValue := parserInt64Pools[thread].Get()
				*newValue = *E_F_S_T1.Value.(*int64) + *E_F_S_T3.Value.(*int64)
				E_S0.Value = newValue
			}*/
		var E_F_S_T3 int64
		var E_F_S_T1 int64
		var LHS int64
		E_F_S_T3, err = mem.pop()
		if err != nil {
			return err
		}
		E_F_S_T1, err = mem.pop()
		if err != nil {
			return err
		}
		LHS = E_F_S_T1 + E_F_S_T3
		err = mem.push(LHS)
		return err
	case 2:
		/*
			E_S0 := lhs
			E_F_S_T1 := rhs[0]
			PLUS2 := rhs[1]
			E_S_T3 := rhs[2]

			E_S0.Child = E_F_S_T1
			E_F_S_T1.Next = PLUS2
			PLUS2.Next = E_S_T3

			{
				newValue := parserInt64Pools[thread].Get()
				*newValue = *E_F_S_T1.Value.(*int64) + *E_S_T3.Value.(*int64)
				E_S0.Value = newValue
			}*/
		var E_F_S_T1 int64
		var E_S_T3 int64
		var LHS int64
		E_S_T3, err = mem.pop()
		if err != nil {
			return err
		}
		E_F_S_T1, err = mem.pop()
		if err != nil {
			return err
		}
		LHS = E_F_S_T1 + E_S_T3
		err = mem.push(LHS)
		return err
	case 3:
		/*
			E_S_T0 := lhs
			E_F_S_T1 := rhs[0]
			TIMES2 := rhs[1]
			E_F_S_T3 := rhs[2]

			E_S_T0.Child = E_F_S_T1
			E_F_S_T1.Next = TIMES2
			TIMES2.Next = E_F_S_T3

			{
				newValue := parserInt64Pools[thread].Get()
				*newValue = *E_F_S_T1.Value.(*int64) * *E_F_S_T3.Value.(*int64)
				E_S_T0.Value = newValue
			}*/
		var E_F_S_T1 int64
		var E_F_S_T3 int64
		var LHS int64
		E_F_S_T3, err = mem.pop()
		if err != nil {
			return err
		}
		E_F_S_T1, err = mem.pop()
		if err != nil {
			return err
		}
		LHS = E_F_S_T1 * E_F_S_T3
		err = mem.push(LHS)
		return err

	case 4:
		/*
			NEW_AXIOM0 := lhs
			E_S1 := rhs[0]

			NEW_AXIOM0.Child = E_S1

			{
				NEW_AXIOM0.Value = E_S1.Value
			}*/
		return nil
	case 5:
		/*E_S0 := lhs
		E_S1 := rhs[0]
		PLUS2 := rhs[1]
		E_F_S_T3 := rhs[2]

		E_S0.Child = E_S1
		E_S1.Next = PLUS2
		PLUS2.Next = E_F_S_T3

		{
			newValue := parserInt64Pools[thread].Get()
			*newValue = *E_S1.Value.(*int64) + *E_F_S_T3.Value.(*int64)
			E_S0.Value = newValue
		}*/
		var E_S1 int64
		var E_F_S_T3 int64
		var LHS int64
		E_F_S_T3, err = mem.pop()
		if err != nil {
			return err
		}
		E_S1, err = mem.pop()
		if err != nil {
			return err
		}
		LHS = E_S1 + E_F_S_T3
		err = mem.push(LHS)
		return err
	case 6:
		/*
			E_S0 := lhs
			E_S1 := rhs[0]
			PLUS2 := rhs[1]
			E_S_T3 := rhs[2]

			E_S0.Child = E_S1
			E_S1.Next = PLUS2
			PLUS2.Next = E_S_T3

			{
				newValue := parserInt64Pools[thread].Get()
				*newValue = *E_S1.Value.(*int64) + *E_S_T3.Value.(*int64)
				E_S0.Value = newValue
			}*/
		var E_S1 int64
		var E_S_T3 int64
		var LHS int64
		E_S_T3, err = mem.pop()
		if err != nil {
			return err
		}
		E_S1, err = mem.pop()
		if err != nil {
			return err
		}
		LHS = E_S1 + E_S_T3
		err = mem.push(LHS)
		return err

	case 7: /*
			NEW_AXIOM0 := lhs
			E_S_T1 := rhs[0]

			NEW_AXIOM0.Child = E_S_T1

			{
				NEW_AXIOM0.Value = E_S_T1.Value
			}*/
		return nil
	case 8: /*
			E_S0 := lhs
			E_S_T1 := rhs[0]
			PLUS2 := rhs[1]
			E_F_S_T3 := rhs[2]

			E_S0.Child = E_S_T1
			E_S_T1.Next = PLUS2
			PLUS2.Next = E_F_S_T3

			{
				newValue := parserInt64Pools[thread].Get()
				*newValue = *E_S_T1.Value.(*int64) + *E_F_S_T3.Value.(*int64)
				E_S0.Value = newValue
			}*/
		var E_S_T1 int64
		var E_F_S_T3 int64
		var LHS int64
		E_F_S_T3, err = mem.pop()
		if err != nil {
			return err
		}
		E_S_T1, err = mem.pop()
		if err != nil {
			return err
		}
		LHS = E_S_T1 + E_F_S_T3
		err = mem.push(LHS)
		return err
	case 9:
		/*
			E_S0 := lhs
			E_S_T1 := rhs[0]
			PLUS2 := rhs[1]
			E_S_T3 := rhs[2]

			E_S0.Child = E_S_T1
			E_S_T1.Next = PLUS2
			PLUS2.Next = E_S_T3

			{
				newValue := parserInt64Pools[thread].Get()
				*newValue = *E_S_T1.Value.(*int64) + *E_S_T3.Value.(*int64)
				E_S0.Value = newValue
			}*/
		var E_S_T1 int64
		var E_S_T3 int64
		var LHS int64
		E_S_T3, err = mem.pop()
		if err != nil {
			return err
		}
		E_S_T1, err = mem.pop()
		if err != nil {
			return err
		}
		LHS = E_S_T1 + E_S_T3
		err = mem.push(LHS)
		return err
	case 10: /*
			E_S_T0 := lhs
			E_S_T1 := rhs[0]
			TIMES2 := rhs[1]
			E_F_S_T3 := rhs[2]

			E_S_T0.Child = E_S_T1
			E_S_T1.Next = TIMES2
			TIMES2.Next = E_F_S_T3

			{
				newValue := parserInt64Pools[thread].Get()
				*newValue = *E_S_T1.Value.(*int64) * *E_F_S_T3.Value.(*int64)
				E_S_T0.Value = newValue
			}*/
		var E_S_T1 int64
		var E_F_S_T3 int64
		var LHS int64
		E_F_S_T3, err = mem.pop()
		if err != nil {
			return err
		}
		E_S_T1, err = mem.pop()
		if err != nil {
			return err
		}
		LHS = E_S_T1 * E_F_S_T3
		err = mem.push(LHS)
		return err
	case 11: /*
			E_F_S_T0 := lhs
			LPAR1 := rhs[0]
			E_F_S_T2 := rhs[1]
			RPAR3 := rhs[2]

			E_F_S_T0.Child = LPAR1
			LPAR1.Next = E_F_S_T2
			E_F_S_T2.Next = RPAR3

			{
				E_F_S_T0.Value = E_F_S_T2.Value
			}*/
		return nil
	case 12: /*
			E_F_S_T0 := lhs
			LPAR1 := rhs[0]
			E_S2 := rhs[1]
			RPAR3 := rhs[2]

			E_F_S_T0.Child = LPAR1
			LPAR1.Next = E_S2
			E_S2.Next = RPAR3

			{
				E_F_S_T0.Value = E_S2.Value
			}*/
		return nil
	case 13: /*
			E_F_S_T0 := lhs
			LPAR1 := rhs[0]
			E_S_T2 := rhs[1]
			RPAR3 := rhs[2]

			E_F_S_T0.Child = LPAR1
			LPAR1.Next = E_S_T2
			E_S_T2.Next = RPAR3

			{
				E_F_S_T0.Value = E_S_T2.Value
			}*/
		return nil
	case 14:
		/*
			E_F_S_T0 := lhs
			NUMBER1 := rhs[0]

			E_F_S_T0.Child = NUMBER1

			{
				E_F_S_T0.Value = NUMBER1.Value
			}*/
		return nil
	}
	return err
}
