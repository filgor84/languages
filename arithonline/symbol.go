package arithonline

type Symbol struct {
	symbolId   uint16
	precedence uint16
}

/*func checkValue(value interface{}) uint8 {
	switch value.(type) {
	case int64:
		return VALUETYPE_INT64
	default:
		return NO_VALUETYPE
	}

}
*/
