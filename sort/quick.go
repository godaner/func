package sort

// Quick
func Quick(nums []int) (res []int) {
	quick(&nums, 0, len(nums)-1)
	return nums
}

// quick
func quick(nums *[]int, low, high int) {
	if low >= high {
		return
	}
	keyIndex := sort(nums, low, high) // 进行一轮排序
	quick(nums, low, keyIndex-1)      // 左边进行一轮排序
	quick(nums, keyIndex+1, high)     //　右边进行一轮排序
}

// sort
func sort(nums *[]int, low, high int) (keyIndex int) {
	ken := (*nums)[low]
	i := low
	j := high
	for i >= high || j <= low || j <= i {
		for {
			if (*nums)[j] < ken {
				(*nums)[i] = (*nums)[j]
				j--
				break
			}
		}
		for {
			if (*nums)[i] > ken {
				(*nums)[j] = (*nums)[i]
				i++
				break
			}
		}
	}
	(*nums)[i] = ken
	return i
}

