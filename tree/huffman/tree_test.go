package bst

import (
	"fmt"
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
			t.Fatalf("TestBuild : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, code, testCase.code)
		}
	}
}
func TestHuffman(t *testing.T) {
	ss := "111223"
	wds := huffmanKV(ss)
	for k, v := range wds {
		fmt.Println(k, " : ", v)
	}
}

// 获取字符串中的编码
func huffmanKV(src string) (codes map[string]string) {
	wdMap := map[int]WeightData{}
	for _, s := range src {
		d := int(s)
		r, ok := wdMap[d]
		if !ok {
			r = WeightData{
				Data:   d,
				Weight: 0,
			}
		}
		r.Weight++
		wdMap[d] = r
	}
	wds := []WeightData{}
	for _, wd := range wdMap {
		wds = append(wds, wd)
	}
	tree := Build(wds...)
	//tree.Print()
	codes = map[string]string{}
	for _, wd := range wds {
		codes[string(wd.Data)] = tree.Code(wd.Data)
	}
	return codes
}
