package main

import (
	"fmt"
)

func varssMain() {

	var intVal int = 21
	var intVal32 int32 = 22
	var intVal64 int64 = 23

	fmt.Println("intVal", intVal)
	fmt.Println("intVal32", intVal32)
	fmt.Println("intVal64", intVal64)

	if intVal == 21 {
		fmt.Println("its 21!!!!!!!")
	} else if intVal32 == 22 {
		fmt.Println("but its 22 :(")
	} else {
		fmt.Printf("wtf is %v\n", intVal64)
	}

	switch {
	case intVal64 == 21:
		fmt.Print("its a 21")
	case intVal32 != 22:
		fmt.Print("its not 22")
	default:
		fmt.Print("its not 21 :(")
	}

	switch intVal {
	case 22:
		fmt.Println("its 22")
	case 21:
		fmt.Println("its 2111111")
	default:
		fmt.Println("Its not 2111")
	}
}

func main() {
	varssMain()
}
