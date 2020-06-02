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
		{
			find: 2,
			code: "",
			src: []WeightData{
				{
					Data:   2,
					Weight: 56160,
				},
			},
		},
	}
	for i, testCase := range testCases {
		tree := Build(testCase.src...)
		//tree.Print()
		code := tree.Code(testCase.find)
		if code != testCase.code {
			t.Fatalf("TestTree_Code : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, code, testCase.code)
		}
	}
}
func TestTree_Codes(t *testing.T) {
	testCases := []struct {
		src   []WeightData
		codes map[int]string
	}{
		{
			codes: map[int]string{
				1: "1",
				2: "01",
				3: "00",
			},
			src: []WeightData{
				{
					Data:   1,
					Weight: 3,
				},
				{
					Data:   2,
					Weight: 2,
				},
				{
					Data:   3,
					Weight: 1,
				},
			},
		},
	}
	for i, testCase := range testCases {
		tree := Build(testCase.src...)
		//tree.Print()
		codes := tree.Codes()
		if !sameCodes(testCase.codes, codes) {
			t.Fatalf("TestTree_Codes : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, codes, testCase.codes)
		}
	}
}

func sameCodes(cs1 map[int]string, cs2 map[int]string) (b bool) {
	if len(cs1) != len(cs2) {
		return false
	}
	for k1, v1 := range cs1 {
		v2, ok := cs2[k1]
		if !ok {
			return false
		}
		if v2 != v1 {
			return false
		}
	}
	return true
}
