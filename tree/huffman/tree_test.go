package bst

import (
	"testing"
)

// compare
// 1 : a > b
// 0 : a = b
// -1 : a < b
type compare func(a, b interface{}) (res int)



func TestTree_Print(t *testing.T) {
	testCases := []struct {
		src []WeightData
	}{
		{
			src: []WeightData{
				{
					Data:   2,
					Weight: 56160,
				},
				{
					Data:   213,
					Weight: 5611160,
				},
				{
					Data:   55,
					Weight: 160,
				},
				{
					Data:   11,
					Weight: 5660,
				},
				{
					Data:   22,
					Weight: 560,
				},
			},
		},

	}
	for _, testCase := range testCases {
		tree := Build(testCase.src...)
		tree.Print()
	}
}
