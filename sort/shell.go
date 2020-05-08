package sort

// Shell
func Shell(nums []int, dk int) (res []int) {
	for { // 计算dk
		dk = dk / 2
		if dk == 0 {
			return nums
		}
		for i := 0; i < dk; i++ { // 循环各个分组
			for j := i; j < len(nums); j = j + dk { //循环每个分组 , 进行插入排序
				p := j - dk
				c := nums[j]
				for p >= 0 && c < nums[p] {
					nums[p+dk] = nums[p]
					p--
				}
				nums[p+dk] = c
			}
		}

	}

}
