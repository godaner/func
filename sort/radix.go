package sort

import (
	"strconv"
)

// Radix
func Radix(nums []int) (res []int) {
	maxN := maxNumber(nums)
	wn := getWN(maxN)
	for i := 1; i <= wn; i++ { // 第n轮，根据最大数的位数来确定
		bucket := make([][]int, 10)
		// to bucket
		for _, n := range nums {
			w := getW(n, i) //获取指定位置的数字
			val := bucket[w]
			if val == nil {
				val = []int{}
			}
			val = append(val, n)
			bucket[w] = val
		}
		// get from bucket
		nums = []int{}
		for _, val := range bucket {
			nums = append(nums, val...)
		}
	}
	return nums
}

// getW
//  获取指定位置的数字
//  num:1324598 p:2 -> 9
//  num:1324598 p:1 -> 8
//  num:1324598 p:7 -> 1
//  num:-1324598 p:8 -> 0
func getW(num, p int) int {
	ns := strconv.FormatInt(int64(num), 10)
	nsL := len(ns)
	if nsL < p {
		return 0
	}
	p = nsL - p
	r, _ := strconv.ParseInt(ns[p:p+1], 10, 64)
	return int(r)
}

// getWN
//  获取数字的位数
func getWN(num int) int {
	return len(strconv.FormatInt(int64(num), 10))
}
