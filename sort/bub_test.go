package sort

import (
	"testing"
)


func TestBub(t *testing.T) {
	for i, testCase := range testCases {
		res := Bub(testCase.src)
		c := testCase.comp(testCase.res, res)
		t.Logf("TestBub : testCase[%v] actually res is : %v , wanna res is : %v !", i, res, testCase.res)
		if c != 0 {
			t.Errorf("TestBub : testCase[%v] fail !", i)
		}
	}
}
