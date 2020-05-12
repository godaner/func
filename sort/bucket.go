package sort

// Bucket
func Bucket(nums []int, bucketSize int) (res []int) {
	maxN := maxNumber(nums)
	minN := minNumber(nums)
	bucket := make([][]int, (maxN-minN)/bucketSize+1)

	// to bucket
	for _, num := range nums {
		p := num / bucketSize
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
