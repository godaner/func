package sort

// swap
func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

// compare
// 1 : a > b
// 0 : a = b
// -1 : a < b
type compare func(a, b interface{}) (res int)

var intsCompare compare

// testCase
type testCase struct {
	src  []int
	res  []int
	comp compare
}

var testCases []testCase

func init() {
	// intsCompare
	intsCompare = func(a, b interface{}) (res int) {
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

	// testCases
	testCases = []testCase{
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
}
