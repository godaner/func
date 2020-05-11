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
// buildMaxHeap
//  构建大根堆
func buildMaxHeap(nums *[]int) {
	arrLen := len(*nums)
	//　从最底层开始构建大根堆
	for nodeIndex := arrLen - 1; nodeIndex >= 0; nodeIndex-- {
		adjustHeap(nums, nodeIndex, arrLen)
	}
}
// adjustHeap
//  寻找某一个指定节点在大根堆的位置
//  即与起子节点比较大小，小于子节点，那么沉降，再利用递归的方式继续和子孙节点比较
//  nodeIndex:待处理节点索引
//  arrLen:总结点数，待处理的节点最大比较范围
func adjustHeap(nums *[]int, nodeIndex, arrLen int) {
	if nodeIndex >= arrLen {
		return
	}
	// 寻找当前节点，左子节点，右子节点三者中最大者的索引
	leftChildIndex := 2*nodeIndex + 1
	rightChildIndex := 2*nodeIndex + 2
	maxNumberIndex := nodeIndex
	if leftChildIndex < arrLen {
		maxNumberIndex = maxIndex(*nums, maxNumberIndex, leftChildIndex)
	}
	if rightChildIndex < arrLen {
		maxNumberIndex = maxIndex(*nums, maxNumberIndex, rightChildIndex)
	}
	// 如果当前节点小于其子节点，需要与子节点交换位置；并且递归向下比较当前节点是否比子孙节点小？
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
