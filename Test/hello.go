package main  // 声明 main 包，表明当前是一个可执行程序

import "fmt"  // 导入内置 fmt 
import "os"
import "log"
import "time"

func main(){  // main函数，是程序执行的入口
    fmt.Println("Hello World!")  // 在终端打印 Hello World!


    log.Printf("cwd1: %v", cwd)
    log.Printf("cwd2: %v", cwd)
}

var cwd string = "3344"

func init() {
    q := [3]int{1, 2, 3}
    log.Printf("q1: %v", q)
    q = [...]int{2, 3, 4} // compile error: cannot assign [4]int to [3]int
    log.Printf("q2: %v", q)


    cwd, err := os.Getwd() 
    if err != nil {
        log.Fatalf("os.Getwd failed: %v", err)
    }
    log.Printf("Working directory = %s", cwd)

    cwd = "123"
    
    cwd, err1 := os.Getwd()
    log.Printf("error = %v", err1)


    const day = 24 * time.Hour
    fmt.Printf("%T %[1]v\n", day.Seconds()) // "86400"
    fmt.Printf("%T\n", day.Seconds)
}
