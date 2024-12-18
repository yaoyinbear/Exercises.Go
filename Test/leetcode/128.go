package main

func longestConsecutive(nums []int) int {
	hash := make(map[int]*struct{ left, right int })
	s := 0
	for idx, num := range nums {
		if _, ok := hash[num]; ok {
			continue
		}

		left := idx
		right := idx

		okL := false
		if v, ok := hash[num-1]; ok {
			left = v.left
			okL = ok
		}
		okR := false
		if v, ok := hash[num+1]; ok {
			right = v.right
			okR = ok
		}

		if okL {
			hash[nums[left]].right = right
		}
		if okR {
			hash[nums[right]].left = left
		}

		hash[num] = &struct{ left, right int }{left, right}

		s = max(s, nums[right]-nums[left]+1)
	}

	return s
}

func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	ans := longestConsecutive(nums)
	println(ans)
}
