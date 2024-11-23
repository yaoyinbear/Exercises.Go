package main

import "fmt"

type People interface {
    
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
    if think == "sb" {
        talk = "你是个大帅比"
    } else {
        talk = "您好"
    }
    return
}

func main() {
	var peo People = &Stduent{}
    think := "bitch"
    fmt.Println(peo.Speak(think))

	    // 定义一个空接口x
		var x interface{}
		s := "pprof.cn"
		x = s
		fmt.Printf("type:%T value:%v\n", x, x)
		i := 100
		x = i
		fmt.Printf("type:%T value:%v\n", x, x)
		b := true
		x = b
		fmt.Printf("type:%T value:%v\n", x, x)
}