package bst

import (
	"testing"
)

// compare
// 1 : a > b
// 0 : a = b
// -1 : a < b
type compare func(a, b interface{}) (res int)

func TestBuild(t *testing.T) {
	c := func(a, b interface{}) (res int) {
		ai, _ := a.([]int)
		bi, _ := b.([]int)
		for index, ain := range ai {
			if bi[index] != ain {
				return -1
			}
		}
		return 0
	}
	testCases := []struct {
		src      []int
		wannaRes []int
		comp     compare
	}{
		{
			src:      []int{0, 1, 3, 7, 8, 4, 2, 5, 9, 6},
			wannaRes: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			comp:     c,
		},
		{
			src:      []int{1, 0, 1, 3, 7, 8, 4, 2, 5, 9, 6, -1},
			wannaRes: []int{-1, 0, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			comp:     c,
		},
		{
			src:      []int{-1, -1, -2, 0, 1, 1, 1},
			wannaRes: []int{-2, -1, -1, 0, 1, 1, 1},
			comp:     c,
		},
	}

	for i, testCase := range testCases {
		tree := Build(testCase.src)
		actuallyPreRes := tree.Mid()
		c := testCase.comp(testCase.wannaRes, actuallyPreRes)
		if c != 0 {
			t.Fatalf("TestBuild : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, actuallyPreRes, testCase.wannaRes)
		}
	}
}
