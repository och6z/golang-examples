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
