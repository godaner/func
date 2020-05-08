package sort

import (
	"testing"
)


func TestMerge(t *testing.T) {
	for i, testCase := range testCases {
		res := Merge(testCase.src)
		c := testCase.comp(testCase.res, res)
		t.Logf("TestMerge : testCase[%v] actually res is : %v , wanna res is : %v !", i, res, testCase.res)
		if c != 0 {
			t.Errorf("TestMerge : testCase[%v] fail !", i)
		}
	}
}
