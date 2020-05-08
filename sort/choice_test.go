package sort

import (
	"testing"
)

func TestChoice(t *testing.T) {
	for i, testCase := range testCases {
		res := Choice(testCase.src)
		c := testCase.comp(testCase.res, res)
		//t.Logf("TestChoice : testCase[%v] actually res is : %v , wanna res is : %v !", i, res, testCase.res)
		if c != 0 {
			t.Fatalf("TestChoice : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, res, testCase.res)
			//t.Fatalf("TestChoice : testCase[%v] fail !", i)
		}
	}
}
