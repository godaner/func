package sort

func Heap(nums []int) (res []int) {
	arrLen := len(nums)
	for i := 0; i < len(nums); i++ {
		maxHeap(&nums, arrLen-i)
		swap(&(nums[0]), &(nums[len(nums)-i-1]))
	}
	return nums
}

// 构建根堆
func maxHeap(nums *[]int, arrLen int) {
	for nodeIndex := len(*nums) - 1; nodeIndex >= 0; nodeIndex-- {
		oneMaxHeap(nums, nodeIndex, arrLen)
	}
}

// 构建一次根堆
func oneMaxHeap(nums *[]int, nodeIndex, arrLen int) {
	if nodeIndex >= arrLen {
		return
	}
	leftChildIndex := 2*nodeIndex + 1
	rightChildIndex := 2*nodeIndex + 2
	maxNumberIndex := nodeIndex
	if leftChildIndex < arrLen {
		maxNumberIndex = maxIndex(*nums, maxNumberIndex, leftChildIndex)
	}
	if rightChildIndex < arrLen {
		maxNumberIndex = maxIndex(*nums, maxNumberIndex, rightChildIndex)
	}
	// 从子节点找到了新的最大数
	if maxNumberIndex != nodeIndex {
		swap(&((*nums)[maxNumberIndex]), &((*nums)[nodeIndex]))
		oneMaxHeap(nums, maxNumberIndex, arrLen)
	}
}

// maxIndex
func maxIndex(arr []int, index1, index2 int) (index int) {
	if arr[index1] >= arr[index2] {
		return index1
	}
	return index2
}
