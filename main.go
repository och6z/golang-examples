package main

import (
	"fmt"

	"example.com/m/cstructs"
	"example.com/m/names"
)

func main() {
	fmt.Println(names.HelloWorldPrefix + names.HelloWorld)
	fmt.Println(names.HelloWorldPrefix + names.World())
	fmt.Println(cstructs.IfTrue())
	fmt.Println(cstructs.IfFalse())
	fmt.Println(cstructs.IfInitializationTrue())
	fmt.Println(cstructs.IfInitializationFalse())
	fmt.Println(cstructs.IfExpression())
	fmt.Println(cstructs.IfElseTrue())
	fmt.Println(cstructs.IfElseFalse())
	fmt.Println(cstructs.IfElseInitializationFalse())
	fmt.Println(cstructs.IfElseVariableFalse())
	fmt.Println(cstructs.SwitchExpressionLeftToRightDefault())
	fmt.Println(cstructs.SwitchExpressionTopToBottomDefault())
	fmt.Println(cstructs.SwitchExpressionHelloDefault())
}
