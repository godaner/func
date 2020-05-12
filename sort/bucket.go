package sort

// Bucket
func Bucket(nums []int, bucketSize int) (res []int) {
	maxN := maxNumber(nums)
	minN := minNumber(nums)
	bucket := make([][]int, (maxN-minN)/bucketSize+1)

	// to bucket
	for _, num := range nums {
		p := (num - minN)/bucketSize // 映射关系
		val := bucket[p]
		if val == nil {
			val = make([]int, 0)
		}
		val = append(val, num)
		bucket[p] = val
	}
	// sort
	res = make([]int, 0)
	for _, val := range bucket {
		if len(val) == 0 {
			continue
		}
		val = Insert(val)
		res = append(res, val...)
	}
	return res
}
func BucketV2(nums []int, bucketSize int) (res []int) {
	maxN := maxNumber(nums)
	minN := minNumber(nums)
	bucket := make(map[int][]int)

	// to bucket
	for _, num := range nums {
		p := (num-minN) / bucketSize // 映射关系
		val, _ := bucket[p]
		if val == nil {
			val = make([]int, 0)
		}
		val = append(val, num)
		bucket[p] = val
	}
	// sort
	for i := 0; i < (maxN-minN)/bucketSize+1; i++ {
		val, _ := bucket[i]
		if len(val) == 0 {
			continue
		}
		val = Insert(val)
		res = append(res, val...)
	}
	return res
}
