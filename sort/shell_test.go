package sort

import (
	"testing"
)

func TestShell(t *testing.T) {
	for i, testCase := range testCases {
		res := Shell(testCase.src, len(testCase.src)/3)
		c := testCase.comp(testCase.res, res)
		//t.Logf("TestShell : testCase[%v] actually res is : %v , wanna res is : %v !", i, res, testCase.res)
		if c != 0 {
			t.Fatalf("TestShell : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, res, testCase.res)
			//t.Fatalf("TestShell : testCase[%v] fail !", i)
		}
	}
}
