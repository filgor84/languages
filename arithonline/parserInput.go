package arithonline

var rule2InputNumber = map[int]int{
	0:  0,
	1:  2,
	2:  2,
	3:  2,
	4:  0,
	5:  2,
	6:  2,
	7:  0,
	8:  2,
	9:  2,
	10: 2,
	11: 0,
	12: 0,
	13: 0,
	14: 0,
}

func parseExecutorInput(ruleNum uint16) int {

	return 0
}

func parseExecutorNew(ruleNum uint16, x int64, y int64) (bool, int64) {
	var LHS int64
	switch ruleNum {
	case 0:
		return false, 0
	case 1:

		LHS := x + y

		return true, LHS
	case 2:

		LHS = x + y
		return true, LHS
	case 3:

		LHS = x * y
		return true, LHS

	case 4:
		return false, 0
	case 5:

		LHS = x + y
		return true, LHS

	case 6:

		LHS = x + y
		return true, LHS

	case 7:
		return false, 0
	case 8:

		LHS = x + y
		return true, LHS

	case 9:

		LHS = x + y
		return true, LHS

	case 10:

		LHS = x * y
		return true, LHS

	case 11:
		return false, 0
	case 12:
		return false, 0

	case 13:
		return false, 0
	case 14:
		return false, 0
	}
	return false, 0
}

func parseExecutor(ruleNum uint16, mem *StackInt64) error {
	var err error
	switch ruleNum {
	case 0:
		return nil
	case 1:
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
		return nil
	case 5:
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

	case 7:
		return nil
	case 8:
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
	case 10:
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
	case 11:
		return nil
	case 12:
		return nil
	case 13:
		return nil
	case 14:
		return nil
	}
	return err
}
