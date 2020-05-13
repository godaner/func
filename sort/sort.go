package sort

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	randomTestCaseNum       = 1
	randomTestCaseArrLen    = 50000
	randomTestCaseMaxNumber = 100000000
)

// compare
// 1 : a > b
// 0 : a = b
// -1 : a < b
type compare func(a, b interface{}) (res int)

var intsCompare compare

// testCase
type testCase struct {
	src  []int
	res  []int
	comp compare
}

func (tc *testCase) String() string {
	return "src is : " + fmt.Sprint(tc.src) + " , res is : " + fmt.Sprint(tc.res)
}

var testCases []testCase

func init() {

	// intsCompare
	intsCompare = func(a, b interface{}) (res int) {
		ai, _ := a.([]int)
		bi, _ := b.([]int)
		if len(ai) != len(bi) {
			return 1
		}
		for i, an := range ai {
			if bi[i] != an {
				return 1
			}
		}
		return 0
	}

	// testCases
	testCases = []testCase{
		{
			src:  []int{-250, -9999, -1, 0, 100, 200},
			res:  []int{-9999, -250, -1, 0, 100, 200},
			comp: intsCompare,
		},
		{
			src:  []int{-1, 0, -1},
			res:  []int{-1, -1, 0},
			comp: intsCompare,
		},
		{
			src:  []int{-0, 0, 0},
			res:  []int{-0, 0, 0},
			comp: intsCompare,
		},
		{
			src:  []int{88, 1, 55, 45, 4546, 564, 68, 13, 1, 3654, 896},
			res:  []int{1, 1, 13, 45, 55, 68, 88, 564, 896, 3654, 4546},
			comp: intsCompare,
		},
		{
			src:  []int{-1},
			res:  []int{-1},
			comp: intsCompare,
		},
		{
			src:  []int{1},
			res:  []int{1},
			comp: intsCompare,
		},
		{
			src:  []int{0},
			res:  []int{0},
			comp: intsCompare,
		},
		{
			src:  []int{1, -1, 0},
			res:  []int{-1, 0, 1},
			comp: intsCompare,
		},
		{
			src:  []int{-1, -1, -1},
			res:  []int{-1, -1, -1},
			comp: intsCompare,
		},
	}
	initRandomTestCases(&testCases, randomTestCaseNum, randomTestCaseMaxNumber, randomTestCaseArrLen)
}

// initRandomTestCases
func initRandomTestCases(testCases *[]testCase, num, max, l int) {
	for i := 0; i < num; i++ {
		arr := []int{}
		for i := 0; i < l; i++ {
			rand.Seed(time.Now().UnixNano())
			arr = append(arr, rand.Intn(max))
		}
		*testCases = append(*testCases, testCase{
			src:  arr,
			res:  Shell(arr, len(arr)/2),
			comp: intsCompare,
		})
	}
}

// swap
func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

// maxNumber
func maxNumber(nums []int) (maxN int) {
	for _, num := range nums {
		if num > maxN {
			maxN = num
		}
	}
	return maxN
}

// minNumber
func minNumber(nums []int) (minN int) {
	for _, num := range nums {
		if num < minN {
			minN = num
		}
	}
	return minN
}
