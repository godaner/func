package sort

// Merge
func Merge(nums []int) (res []int) {
	split(&nums, 0, len(nums)-1)
	return nums
}

// split
func split(nums *[]int, low, high int) {
	if low >= high {
		return
	}
	mid := (low + high) / 2
	split(nums, low, mid)
	split(nums, mid+1, high)
	merge(nums, low, mid, high)
}

// merge
func merge(nums *[]int, low, mid, high int) {
	tl := high - low + 1
	temp := make([]int, tl)
	i := low
	j := mid + 1
	k := 0
	for i <= mid && j <= high {
		if (*nums)[i] < (*nums)[j] {
			temp[k] = (*nums)[i]
			i++
		} else {
			temp[k] = (*nums)[j]
			j++
		}
		k++
	}
	for i <= mid {
		temp[k] = (*nums)[i]
		i++
		k++
	}
	for j <= high {
		temp[k] = (*nums)[j]
		j++
		k++
	}
	for i := 0; i < tl; i++ {
		(*nums)[low+i] = temp[i]
	}
}
