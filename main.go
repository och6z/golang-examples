package main

import (
	"fmt"

	"example.com/m/names"
)

func main() {
	fmt.Println(names.HelloWorldPrefix + names.HelloWorld)
	fmt.Println(names.HelloWorldPrefix + names.World())
}
