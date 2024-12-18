package qsort

import (
	"reflect"
	"testing"
)

type CmpFunc func(interface{}, interface{}) bool

func lessFunc(a interface{}, b interface{}) bool {
	return a.(int) < b.(int)
}
func moreFunc(a interface{}, b interface{}) bool {
	return a.(int) > b.(int)
}

func TestQSort(t *testing.T) {
	type test struct {
		name string
		in   []int
		out  []int
		fn   CmpFunc
	}

	tests := []test{
		{name: "1", in: []int{3, 2, 4, 6, 8, 1, 9, 7, 5}, out: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, fn: lessFunc},
		{name: "2", in: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, out: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, fn: lessFunc},
		{name: "3", in: []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, out: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, fn: lessFunc},
		{name: "1.r", in: []int{3, 2, 4, 6, 8, 1, 9, 7, 5}, out: []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, fn: moreFunc},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := make([]interface{}, len(tc.in))
			for i, v := range tc.in {
				s[i] = v
			}

			QSort(s, tc.fn)

			o := make([]int, len(tc.in))
			for i, v := range s {
				o[i] = v.(int)
			}

			if !reflect.DeepEqual(o, tc.out) {
				t.Errorf("expected:%#v, got:%#v", tc.out, o)
			}
		})
	}
}
