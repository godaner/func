package sort

import (
	"testing"
)


func TestQuick(t *testing.T) {

	for i, testCase := range testCases {
		res := Quick(testCase.src)
		c := testCase.comp(testCase.res, res)
		t.Logf("TestQuick : testCase[%v] actually res is : %v , wanna res is : %v !", i, res, testCase.res)
		if c != 0 {
			t.Errorf("TestQuick : testCase[%v] fail !", i)
		}
	}
}
