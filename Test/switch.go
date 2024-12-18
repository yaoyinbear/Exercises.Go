package main

import (
	"fmt"
)

func main() {
	switch true {
	default:
		fmt.Println("D")
	case true:
		fmt.Println("T1")
	case true:
		fmt.Println("T2")

	}


	switch x:=false
	{
	case true:
		fmt.Printf("T3: %v\n", x)
	case false:
		fmt.Printf("T4: %v\n", x)
	}
}

func False() bool {
	return false
}