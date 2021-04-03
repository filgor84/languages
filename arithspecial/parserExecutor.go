package arithspecial

import "errors"

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

func parserExecutor(ruleNum uint16, stack []int64, top *int) error {
	var input1 int64
	var input2 int64
	switch ruleNum {
	case 0:
	case 1:
		input1 = stack[(*top)-1]
		*top--
		input2 = stack[(*top)-1]
		*top--
		stack[*top] = input1 + input2
		*top++
	case 2:
		input1 = stack[(*top)-1]
		*top--
		input2 = stack[(*top)-1]
		*top--
		stack[*top] = input1 + input2
		*top++
	case 3:
		input1 = stack[(*top)-1]
		*top--
		input2 = stack[(*top)-1]
		*top--
		stack[*top] = input1 * input2
		*top++
	case 4:
	case 5:
		input1 = stack[(*top)-1]
		*top--
		input2 = stack[(*top)-1]
		*top--
		stack[*top] = input1 + input2
		*top++
	case 6:
		input1 = stack[(*top)-1]
		*top--
		input2 = stack[(*top)-1]
		*top--
		stack[*top] = input1 + input2
		*top++
	case 7:
	case 8:

		input1 = stack[(*top)-1]
		*top--
		input2 = stack[(*top)-1]
		*top--
		stack[*top] = input1 + input2
		*top++
	case 9:

		input1 = stack[(*top)-1]
		*top--
		input2 = stack[(*top)-1]
		*top--
		stack[*top] = input1 + input2
		*top++
	case 10:
		input1 = stack[(*top)-1]
		*top--
		input2 = stack[(*top)-1]
		*top--
		stack[*top] = input1 * input2
		*top++
	case 11:
	case 12:
	case 13:
	case 14:
	default:
		return errors.New("Parsed unknown rule")
	}
	return nil
}

/*
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
*/
