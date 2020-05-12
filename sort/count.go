package sort

func Count(nums []int) (res []int) {
	maxN := maxNumber(nums)
	minN := minNumber(nums)
	m := make(map[int]int, maxN-minN)
	// collect
	for _, num := range nums {
		val, ok := m[num]
		if !ok {
			val = 0
		}
		val++
		m[num] = val
	}
	// sort
	res = make([]int, 0)
	for i := minN; i <= maxN; i++ {
		val, _ := m[i]
		if val <= 0 {
			continue
		}
		for j := 0; j < val; j++ {
			res = append(res, i)
		}
	}
	return res
}
