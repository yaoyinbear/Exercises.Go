package main

import (
    "fmt"
)


type IntArray []int
type IntList []int

func (array IntArray) sum() int {
    sum := 0
    for _, v := range array {
       sum += v;
    }
    return sum
}

func (array IntList) sum() int {
    sum := 0
    for _, v := range array {
       sum += v;
    }
    return sum
}

func main() {
	array := IntArray{1, 2, 3}
	var list IntList = IntList{1, 1, 1}
	fmt.Println(array.sum())
	fmt.Println(list.sum())
}