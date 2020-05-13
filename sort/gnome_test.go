package sort

import (
	"testing"
)

func TestGnome(t *testing.T) {
	for i, testCase := range testCases {
		res := Gnome(testCase.src)
		c := testCase.comp(testCase.res, res)
		//t.Logf("TestGnome : testCase[%v] actually res is : %v , wanna res is : %v !", i, res, testCase.res)
		if c != 0 {
			t.Fatalf("TestGnome : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, res, testCase.res)
			//t.Fatalf("TestGnome : testCase[%v] fail !", i)
		}
	}
}
