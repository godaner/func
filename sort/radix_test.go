package sort

import (
	"testing"
)

func TestRadix(t *testing.T) {
	testCases := make([]testCase, 0)
	initRandomTestCases(&testCases, randomTestCaseNum, randomTestCaseMaxNumber, randomTestCaseArrLen)
	for i, testCase := range testCases {
		res := Radix(testCase.src)
		c := testCase.comp(testCase.res, res)
		//t.Logf("TestRadix : testCase[%v] actually res is : %v , wanna res is : %v !", i, res, testCase.res)
		if c != 0 {
			t.Fatalf("TestRadix : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, res, testCase.res)
			//t.Fatalf("TestRadix : testCase[%v] fail !", i)
		}
	}
}

