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

func TestTree_Code(t *testing.T) {
	testCases := []struct {
		src  []WeightData
		find int
		code string
	}{
		{
			find: 22,
			code: "0001",
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
		{
			find: 213,
			code: "1",
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
	for i, testCase := range testCases {
		tree := Build(testCase.src...)
		//tree.Print()
		code := tree.Code(testCase.find)
		if code != testCase.code {
			t.Fatalf("TestBuild : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, code, testCase.code)
		}
	}
}
