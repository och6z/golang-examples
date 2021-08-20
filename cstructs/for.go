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
