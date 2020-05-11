package sort

// Heap
func Heap(nums []int) (res []int) {
	// 初始化大根堆
	buildMaxHeap(&nums)
	// 临时变量
	size := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		// 交换最大数到末尾
		swap(&(nums[0]), &(nums[len(nums)-i-1]))
		// 根节点被换，需要从顶部从新向下调整
		adjustHeap(&nums, 0, size)
		size--
	}
	return nums
}

// 构建大根堆
func buildMaxHeap(nums *[]int) {
	arrLen := len(*nums)
	//　从最底层开始构建大根堆
	for nodeIndex := arrLen - 1; nodeIndex >= 0; nodeIndex-- {
		adjustHeap(nums, nodeIndex, arrLen)
	}
}

//  寻找和处理某个节点在其子树大根堆的位置
//  nodeIndex:待处理节点索引
//  arrLen:总结点数，待处理的节点最大比较范围
func adjustHeap(nums *[]int, nodeIndex, arrLen int) {
	if nodeIndex >= arrLen {
		return
	}
	// 寻找待处理节点的子节点
	leftChildIndex := 2*nodeIndex + 1
	rightChildIndex := 2*nodeIndex + 2
	maxNumberIndex := nodeIndex
	// 寻找是否存在比起大的节点
	if leftChildIndex < arrLen {
		maxNumberIndex = maxIndex(*nums, maxNumberIndex, leftChildIndex)
	}
	if rightChildIndex < arrLen {
		maxNumberIndex = maxIndex(*nums, maxNumberIndex, rightChildIndex)
	}
	// 如果存在比其大的节点，需要交换节点，并且递归向下比较
	if maxNumberIndex != nodeIndex {
		swap(&((*nums)[maxNumberIndex]), &((*nums)[nodeIndex]))
		// 递归向下比较
		adjustHeap(nums, maxNumberIndex, arrLen)
	}
}

// maxIndex
func maxIndex(arr []int, index1, index2 int) (index int) {
	if arr[index1] >= arr[index2] {
		return index1
	}
	return index2
}
