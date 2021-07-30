package cstructs

import "example.com/m/names"

func Hello() string {
	name := ""
	if name == "" {
		name = "World!"
	}
	return names.HelloWorldPrefix + name
}
