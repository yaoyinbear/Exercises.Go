package qsort

import "math/rand"

func QSort(s []interface{}, cmpFunc func(interface{}, interface{}) bool) {
	l := 0
	r := len(s) - 1
	if l >= r {
		return
	}

	k := 0
	pivot := s[rand.Intn(len(s))]

	for k <= r {
		if cmpFunc(s[k], pivot) {
			s[k], s[l] = s[l], s[k]
			l++
			k++
		} else if cmpFunc(pivot, s[k]) {
			s[k], s[r] = s[r], s[k]
			r--
		} else {
			k++
		}
	}
	QSort(s[:l], cmpFunc)
	if r < len(s)-1 {
		QSort(s[r+1:], cmpFunc)
	}
}
