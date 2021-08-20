package cstructs

func ForSingleConditionAbsent() int {
	i = 0
	for {
		i = i + 1
		if i > 5 {
			break
		}
	}
	return i
}
