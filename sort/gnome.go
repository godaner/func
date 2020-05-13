package sort

// Gnome
func Gnome(nums []int) (res []int) {
	if len(nums) <= 1 {
		return nums
	}
	for p := 1; p < len(nums); {
		if p > 0 && nums[p-1] > nums[p] {
			swap(&(nums[p-1]), &(nums[p]))
			p--
		} else {
			p++
		}
	}
	return nums
}
