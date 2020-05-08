package sort

// Insert
func Insert(nums []int) (res []int) {
	for i := 1; i < len(nums); i++ {
		c := nums[i]
		p := i - 1 // 有序集合的最后一个数字索引
		for p >= 0 && c < nums[p] {
			nums[p+1] = nums[p]
			p--
		}
		nums[p+1] = c
	}
	return nums
}
