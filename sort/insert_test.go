package sort

import (
	"testing"
)


func TestInsert(t *testing.T) {

	for i, testCase := range testCases {
		res := Insert(testCase.src)
		c := testCase.comp(testCase.res, res)
		t.Logf("TestInsert : testCase[%v] actually res is : %v , wanna res is : %v !", i, res, testCase.res)
		if c != 0 {
			t.Errorf("TestInsert : testCase[%v] fail !", i)
		}
	}
}
