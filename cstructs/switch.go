package cstructs

func SwitchTrue() string {
	switch true {
	case i5 < 5:
		assertion = assert()
	case i5 == 5:
		assertion = assert()
	case i5 > 5:
		assertion = assert()
	}
	return assertion
}
func Switch() string {
	switch {
	case i5 < 5:
		assertion = assert()
	case i5 == 5:
		assertion = assert()
	case i5 > 5:
		assertion = assert()
	}
	return assertion
}
func SwitchInitializationTrue() string {
	switch x := 5; true {
	case i5 < x:
		assertion = assert()
	case i5 == x:
		assertion = assert()
	case i5 > x:
		assertion = assert()
	}
	return assertion
}
func SwitchExpressionDefault() string {
	switch i5 {
	case 0:
		assertion = assert()
	case 1:
		assertion = assert()
	case 2:
		assertion = assert()
	case 3:
		assertion = assert()
	case 4:
		assertion = assert()
	case 5:
		assertion = assert()
	default:
		assertion = "assertion default"
	}
	return assertion
}
func SwitchExpressionLeftToRightDefault() string {
	switch i5 {
	case 0, 1, 2, 3:
		assertion = assert()
	case 4, 5, 6, 7:
		assertion = assert()
	case 8, 9:
		assertion = assert()
	default:
		assertion = "assertion default"
	}
	return assertion
}
func SwitchExpressionTopToBottomDefault() string {
	switch i5 {
	case 0:
	case 1:
	case 2:
	case 3:
	case 4:
		assertion = assert()
	case 5:
		assertion = assert()
	default:
		assertion = "assertion default"
	}
	return assertion
}
func SwitchExpressionFallthroughDefault() string {
	switch i5 {
	case 0:
		dummy := 0
		_ = dummy
		fallthrough
	case 1:
		dummy := 0
		_ = dummy
		fallthrough
	case 2:
		dummy := 0
		_ = dummy
		fallthrough
	case 3:
		dummy := 0
		_ = dummy
		fallthrough
	case 4:
		dummy := 0
		_ = dummy
		assertion = assert()
	case 5:
		dummy := 0
		_ = dummy
		fallthrough
	default:
		dummy := 0
		_ = dummy
		assertion = "assertion default"
	}
	return assertion
}
