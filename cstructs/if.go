package cstructs

func IfTrue() int {
	count := 0
	if true {
		count = count + 1
	}
	return count
}
func IfFalse() int {
	count := 0
	if false {
		count = count + 1
	}
	return count
}
func IfInitializationTrue() int {
	count := 0
	if one := 1; true {
		count = count + one
	}
	return count
}
func IfInitializationFalse() int {
	count := 0
	if one := 1; false {
		count = count + one
	}
	return count
}
func IfExpression() int {
	count := 0
	if i5 < i7 {
		count = count + 1
	}
	return count
}
func IfElseTrue() int {
	count := 0
	if true {
		count = count + 1
	} else {
		count = count - 1
	}
	return count
}
func IfElseFalse() int {
	count := 0
	if false {
		count = count + 1
	} else {
		count = count - 1
	}
	return count
}
func IfElseInitializationFalse() int {
	count := 0
	if t := 1; false {
		count = count + 1
		_ = t
		t := 7
		_ = t
	} else {
		count = count - t
	}
	return count
}
func IfElseVariableFalse() int {
	count := 0
	t := 1
	if false {
		count = count + 1
		t := 7
		_ = t
	} else {
		count = count - t
	}
	return count
}
