package utils

// TernaryString operator
func TernaryString(condition bool, trueValue, falseValue string) string {
	return map[bool]string{true: trueValue, false: falseValue}[condition]
}
