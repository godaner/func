package sort

import (
	"testing"
)

func TestCount(t *testing.T) {
	for i, testCase := range testCases {
		res := Count(testCase.src)
		c := testCase.comp(testCase.res, res)
		//t.Logf("TestCount : testCase[%v] actually res is : %v , wanna res is : %v !", i, res, testCase.res)
		if c != 0 {
			t.Fatalf("TestCount : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, res, testCase.res)
			//t.Fatalf("TestCount : testCase[%v] fail !", i)
		}
	}
}
