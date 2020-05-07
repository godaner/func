package choice

func Choice(nums []int) (res []int) {
	for i:=0;i<len(nums)-1;i++{
		minP:=i
		for j:=i+1;j<len(nums);j++{
			if nums[minP]>nums[j]{
				minP=j
			}
		}
		swap(&nums[minP],&nums[i])
	}
	return nums
}
// swap
func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}
