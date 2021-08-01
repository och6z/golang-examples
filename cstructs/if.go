package cstructs

import (
	"example.com/m/names"
)

func HelloIf() string {
	name := ""
	if name == "" {
		name = "World!"
	}
	return names.HelloWorldPrefix + name
}

// func HelloIfInitialization() string {
// 	if name := "World!"; name != "" {
// 		return names.HelloWorldPrefix + name
// 	}
// }
func HelloIfElse() string {
	name := "World!"
	if name == "" {
		return names.HelloWorldPrefix + ""
	} else {
		return names.HelloWorldPrefix + name
	}
}
func HelloIfElseIf() string {
	name := "World!"
	if name == "" {
		return names.HelloWorldPrefix + ""
	} else if name == "World!" {
		return names.HelloWorldPrefix + name
	} else {
		return names.HelloWorldPrefix + "Golang!"
	}
}
