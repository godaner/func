package sort

import (
	"testing"
)

func TestBucket(t *testing.T) {
	testCases := make([]testCase, 0)
	initRandomTestCases(&testCases, randomTestCaseNum, randomTestCaseMaxNumber, randomTestCaseArrLen)
	for i, testCase := range testCases {
		res := Bucket(testCase.src, 10)
		c := testCase.comp(testCase.res, res)
		//t.Logf("TestBucket : testCase[%v] actually res is : %v , wanna res is : %v !", i, res, testCase.res)
		if c != 0 {
			t.Fatalf("TestBucket : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, res, testCase.res)
			//t.Fatalf("TestBucket : testCase[%v] fail !", i)
		}
	}
}
func TestBucketV2(t *testing.T) {
	testCases := make([]testCase, 0)
	initRandomTestCases(&testCases, randomTestCaseNum, randomTestCaseMaxNumber, randomTestCaseArrLen)
	for i, testCase := range testCases {
		res := BucketV2(testCase.src, 10)
		c := testCase.comp(testCase.res, res)
		//t.Logf("TestBucketV2 : testCase[%v] actually res is : %v , wanna res is : %v !", i, res, testCase.res)
		if c != 0 {
			t.Fatalf("TestBucketV2 : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, res, testCase.res)
			//t.Fatalf("TestBucketV2 : testCase[%v] fail !", i)
		}
	}
}
