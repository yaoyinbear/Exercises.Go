package skiplist

import (
	"fmt"
	"reflect"
	"testing"
)

type MyInt int

func Less(a, b Key) bool {
	return a.(int) < b.(int)
}

func More(a, b Key) bool {
	return a.(int) > b.(int)
}

func TestSkipList(t *testing.T) {
	type test struct {
		name string
		in   []int
		out  []int
		fn   LessFunc
	}

	tests := []test{
		{name: "1", in: []int{3, 2, 4, 6, 8, 1, 9, 7, 5}, out: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, fn: Less},
		{name: "2", in: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, out: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, fn: Less},
		{name: "3", in: []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, out: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, fn: Less},
		{name: "1.r", in: []int{3, 2, 4, 6, 8, 1, 9, 7, 5}, out: []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, fn: More},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sl := NewSkipList(tc.fn)
			for _, num := range tc.in {
				sl.Insert(num, num*10)
			}

			res := make([]int, 0)
			sl.Range(func(key Key, value Value) bool {
				fmt.Printf("%v(%v) ", key, value)
				res = append(res, key.(int))
				return true
			})
			fmt.Println("")

			if !reflect.DeepEqual(res, tc.out) {
				t.Errorf("excepted:%#v, got:%#v", tc.out, res)
			}
		})
	}
}
