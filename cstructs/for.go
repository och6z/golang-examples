package cstructs

func ForConditionAbsent() int {
	i = 0
	for {
		i = i + 1
		if i > 5 {
			break
		}
	}
	return i
}

func ForCondition() int {
	sum = 0
	for sum < 100 {
		sum = sum + 9
	}
	return sum
}

func ForClause() int {
	sum = 0
	for i := 0; i <= 10; i++ {
		sum = sum + i
	}
	return sum
}

func ForClausePostStmtAbsent() int {
	sum = 0
	for i := 0; i <= 10; {
		sum = sum + i
		i++
	}
	return sum
}

func ForRangeClauseArray() int {
	sum = 0
	for _, value := range [5]int{0, 1, 2, 3, 4} {
		sum += value
	}
	return sum
}

func ForRangeClauseSlice() int {
	sum = 0
	arr := [6]int{0, 1, 2, 3, 4, 5}
	for _, value := range arr[0:5] {
		sum += value
	}
	return sum
}

func ForRangeClauseString() string {
	var assertion []byte
	for _, rn := range assert() {
		assertion = append(assertion, byte(rn))
	}
	return string(assertion)
}
