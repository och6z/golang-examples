package cstructs

func HelloIf() string {
	name := ""
	if name == "" {
		name = "World!"
	}
	return HelloWorldPrefix + name
}

// func HelloIfInitialization() string {
// 	if name := "World!"; name != "" {
// 		return HelloWorldPrefix + name
// 	}
// }
func HelloIfElse() string {
	name := "World!"
	if name == "" {
		return HelloWorldPrefix + ""
	} else {
		return HelloWorldPrefix + name
	}
}
func HelloIfElseIf() string {
	name := "World!"
	if name == "" {
		return HelloWorldPrefix + ""
	} else if name == "World!" {
		return HelloWorldPrefix + name
	} else {
		return HelloWorldPrefix + "Golang!"
	}
}
