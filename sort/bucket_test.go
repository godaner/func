package sort

import (
	"testing"
)

func TestBucket(t *testing.T) {
	for i, testCase := range testCases {
		res := Bucket(testCase.src,10)
		c := testCase.comp(testCase.res, res)
		//t.Logf("TestBucket : testCase[%v] actually res is : %v , wanna res is : %v !", i, res, testCase.res)
		if c != 0 {
			t.Fatalf("TestBucket : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, res, testCase.res)
			//t.Fatalf("TestBucket : testCase[%v] fail !", i)
		}
	}
}
