package insert

import "testing"

// compare
// 1 : a > b
// 0 : a = b
// -1 : a < b
type compare func(a, b interface{}) (res int)

func TestInsert(t *testing.T) {
	// compare
	intsCompare := func(a, b interface{}) (res int) {
		ai, _ := a.([]int)
		bi, _ := b.([]int)
		if len(ai) != len(bi) {
			return 1
		}
		for i, an := range ai {
			if bi[i] != an {
				return 1
			}
		}
		return 0
	}
	// testCases
	testCases := []struct {
		src  []int
		res  []int
		comp compare
	}{
		{
			src:  []int{88, 1, 55, 45, 4546, 564, 68, 13, 1, 3654, 896},
			res:  []int{1, 1, 13, 45, 55, 68, 88, 564, 896, 3654, 4546},
			comp: intsCompare,
		},
		{
			src:  []int{-1},
			res:  []int{-1},
			comp: intsCompare,
		},
		{
			src:  []int{-1, -1, -1},
			res:  []int{-1, -1, -1},
			comp: intsCompare,
		},
	}
	for i, testCase := range testCases {
		res := Insert(testCase.src)
		c := testCase.comp(testCase.res, res)
		t.Logf("TestInsert : testCase[%v] actually res is : %v , wanna res is : %v !", i, res, testCase.res)
		if c != 0 {
			t.Errorf("TestInsert : testCase[%v] fail !", i)
		}
	}
}
