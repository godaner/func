package sort

// Bub
func Bub(nums []int) (res []int) {
	for i := 0; i < len(nums)-1; i++ { //处理len(nums)-1趟
		for j := 0; j < len(nums)-1-i; j++ { // 处理第i个数据
			if nums[j] > nums[j+1] {
				swap(&nums[j], &nums[j+1])
			}
		}
	}
	return nums
}
