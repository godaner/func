package sort

import (
	"testing"
)

func TestHeap(t *testing.T) {
	for i, testCase := range testCases {
		res := Heap(testCase.src)
		c := testCase.comp(testCase.res, res)
		//t.Logf("TestHeap : testCase[%v] actually res is : %v , wanna res is : %v !", i, res, testCase.res)
		if c != 0 {
			t.Fatalf("TestHeap : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, res, testCase.res)
			//t.Fatalf("TestHeap : testCase[%v] fail !", i)
		}
	}
}
