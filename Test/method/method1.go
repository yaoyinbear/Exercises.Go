package main

import "github.com/Exercises.Go/Test/method/pkg1"

func main() {
    var s pkg1.MyStruct
    s.SetName("Alice") // 编译错误：无法访问未导出的字段
}
