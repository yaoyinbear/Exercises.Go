package main

import "fmt"

type Key interface {
	Less(Key)
}

type MyInt int

// func (a *MyInt) Less(b *Key) {
// 	fmt.Printf("a: %v, b: %v\n", a, b)
// }

func (a MyInt) Less(b Key) {
	fmt.Printf("a: %V, b: %V\n", a, b)
}

func main() {
	x := MyInt(1)
	y := MyInt(2)
	var k1 Key = x
	var k2 Key = y
	fmt.Printf("1: %V, 2: %V\n", k1, k2)
	k1.Less(k2)

	var k3 Key = &x
	var k4 Key = &y
	fmt.Printf("3: %V, 4: %V\n", k3, k4)
	k3.Less(k4)
}
