package main

import (
	"fmt"

	"example.com/m/cstructs"
	"example.com/m/names"
)

func main() {
	fmt.Println(names.HelloWorldPrefix + names.HelloWorld)
	fmt.Println(names.HelloWorldPrefix + names.World())

	fmt.Println(cstructs.Hello())
}
