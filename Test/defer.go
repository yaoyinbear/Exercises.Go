package main

import (
	"fmt"
)

type field struct{
    num int
}
func(t *field) print(n int){
    fmt.Println(t.num, n)
}

func test() {
    x, y := 10, 20

    defer func(i int) {
        println("defer:", i, y) // y 闭包引用
    }(x) // x 被复制

    x += 10
    y += 100
    println("x =", x, "y =", y)
}

func main() {    
    var i int = 1
    defer func() {fmt.Println("result2 =>",func() int { return i * 2 }()) }()
    i++

    v := field{1}
    defer v.print(func() int { return i * 2 }())
    v = field{5}
    i++

    // prints: 
    // 5 4
    // result => 6 (not ok if you expected 2)

	test()

	var whatever [5]struct{}
    for i := range whatever {
        defer func() { fmt.Println(i) }()
    }

    x := 10

    defer func() {
        fmt.Println("Deferred anonymous function:", x)
    }()

    x = 20
    fmt.Println("Before return1:", x)

    y := 10

    defer printX(y)

    y = 20
    fmt.Println("Before return2:", y)
}

func printX(x int) {
    fmt.Println("Deferred function:", x)
}