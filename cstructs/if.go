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
