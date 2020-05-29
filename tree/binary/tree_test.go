package binary

import (
	"testing"
)

// compare
// 1 : a > b
// 0 : a = b
// -1 : a < b
type compare func(a, b interface{}) (res int)

func intP(i int) (p *int) {
	return &i
}
func TestTree_Pre(t *testing.T) {
	c := func(a, b interface{}) (res int) {
		ai, _ := a.([]int)
		bi, _ := b.([]int)
		for index, ain := range ai {
			if bi[index] != ain {
				return -1
			}
		}
		return 0
	}
	testCases := []struct {
		preSrc   []*int
		wannaRes []int
		comp     compare
	}{
		{
			//ABDH##I##E##CF#J##G##
			//0137##8##4##25#9##6##
			preSrc:   []*int{intP(0), intP(1), intP(3), intP(7), nil, nil, intP(8), nil, nil, intP(4), nil, nil, intP(2), intP(5), nil, intP(9), nil, nil, intP(6), nil, nil},
			wannaRes: []int{0, 1, 3, 7, 8, 4, 2, 5, 9, 6},
			comp:     c,
		},
	}
	for i, testCase := range testCases {
		tree := BuildFromPre(testCase.preSrc)
		actuallyPreRes := tree.Pre()
		c := testCase.comp(testCase.wannaRes, actuallyPreRes)
		if c != 0 {
			t.Fatalf("TestTree_Pre : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, actuallyPreRes, testCase.wannaRes)
		}
	}
}
func TestTree_Mid(t *testing.T) {
	c := func(a, b interface{}) (res int) {
		ai, _ := a.([]int)
		bi, _ := b.([]int)
		for index, ain := range ai {
			if bi[index] != ain {
				return -1
			}
		}
		return 0
	}
	testCases := []struct {
		preSrc   []*int
		wannaRes []int
		comp     compare
	}{
		{
			preSrc: []*int{intP(0), intP(1), intP(3), intP(7), nil, nil, intP(8), nil, nil, intP(4), nil, nil, intP(2), intP(5), nil, intP(9), nil, nil, intP(6), nil, nil},
			//"HDIBEAFJCG"
			wannaRes: []int{7, 3, 8, 1, 4, 0, 5, 9, 2, 6},
			comp:     c,
		},
	}
	for i, testCase := range testCases {
		tree := BuildFromPre(testCase.preSrc)
		actuallyPreRes := tree.Mid()
		c := testCase.comp(testCase.wannaRes, actuallyPreRes)
		if c != 0 {
			t.Fatalf("TestTree_Mid : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, actuallyPreRes, testCase.wannaRes)
		}
	}
}

func TestTree_Post(t *testing.T) {
	c := func(a, b interface{}) (res int) {
		ai, _ := a.([]int)
		bi, _ := b.([]int)
		for index, ain := range ai {
			if bi[index] != ain {
				return -1
			}
		}
		return 0
	}
	testCases := []struct {
		preSrc   []*int
		wannaRes []int
		comp     compare
	}{
		{
			preSrc: []*int{intP(0), intP(1), intP(3), intP(7), nil, nil, intP(8), nil, nil, intP(4), nil, nil, intP(2), intP(5), nil, intP(9), nil, nil, intP(6), nil, nil},
			//"HIDEBJFGCA"
			wannaRes: []int{7, 8, 3, 4, 1, 9, 5, 6, 2, 0},
			comp:     c,
		},
	}
	for i, testCase := range testCases {
		tree := BuildFromPre(testCase.preSrc)
		actuallyPreRes := tree.Post()
		c := testCase.comp(testCase.wannaRes, actuallyPreRes)
		if c != 0 {
			t.Fatalf("TestTree_Post : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, actuallyPreRes, testCase.wannaRes)
		}
	}
}

func TestTree_Max(t *testing.T) {
	testCases := []struct {
		preSrc   []*int
		wannaRes int
	}{
		{
			preSrc:   []*int{intP(0), intP(1), intP(3), intP(7), nil, nil, intP(8), nil, nil, intP(4), nil, nil, intP(2), intP(5), nil, intP(9), nil, nil, intP(6), nil, nil},
			wannaRes: 9,
		},
	}
	for i, testCase := range testCases {
		tree := BuildFromPre(testCase.preSrc)
		node := tree.Max()

		if node == nil {
			t.Fatalf("TestTree_Max : testCase[%v] fail , res is : nil !", i)
		}
		actuallyRes := node.Elem()
		if testCase.wannaRes != actuallyRes {
			t.Fatalf("TestTree_Max : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, actuallyRes, testCase.wannaRes)
		}
	}
}

func TestTree_Min(t *testing.T) {
	testCases := []struct {
		preSrc   []*int
		wannaRes int
		comp     compare
	}{
		{
			preSrc:   []*int{intP(0), intP(1), intP(3), intP(7), nil, nil, intP(8), nil, nil, intP(4), nil, nil, intP(2), intP(5), nil, intP(9), nil, nil, intP(6), nil, nil},
			wannaRes: 0,
		},
	}
	for i, testCase := range testCases {
		tree := BuildFromPre(testCase.preSrc)
		node := tree.Min()

		if node == nil {
			t.Fatalf("TestTree_Min : testCase[%v] fail , res is : nil !", i)
		}
		actuallyRes := node.Elem()
		if testCase.wannaRes != actuallyRes {
			t.Fatalf("TestTree_Min : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, actuallyRes, testCase.wannaRes)
		}
	}
}

func TestTree_Find(t *testing.T) {
	testCases := []struct {
		preSrc        []*int
		exits         bool
		wannaFindData int
	}{
		{
			preSrc:        []*int{intP(0), intP(1), intP(3), intP(7), nil, nil, intP(8), nil, nil, intP(4), nil, nil, intP(2), intP(5), nil, intP(9), nil, nil, intP(6), nil, nil},
			wannaFindData: 8,
			exits:         true,
		},
		{
			preSrc:        []*int{intP(0), intP(1), intP(3), intP(7), nil, nil, intP(8), nil, nil, intP(4), nil, nil, intP(2), intP(5), nil, intP(9), nil, nil, intP(6), nil, nil},
			wannaFindData: -1,
			exits:         false,
		},
	}
	for i, testCase := range testCases {
		tree := BuildFromPre(testCase.preSrc)
		node := tree.Find(testCase.wannaFindData)
		actuallyRes := node != nil
		if testCase.exits != actuallyRes && testCase.wannaFindData == node.Elem() {
			t.Fatalf("TestTree_Find : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, actuallyRes, testCase.exits)
		}
	}
}
func TestTree_FindParent(t *testing.T) {
	testCases := []struct {
		preSrc        []*int
		exits         bool
		wannaFindData int
	}{
		{
			preSrc:        []*int{intP(0), intP(1), intP(3), intP(7), nil, nil, intP(8), nil, nil, intP(4), nil, nil, intP(2), intP(5), nil, intP(9), nil, nil, intP(6), nil, nil},
			wannaFindData: 8,
			exits:         true,
		},
		{
			preSrc:        []*int{intP(0), intP(1), intP(3), intP(7), nil, nil, intP(8), nil, nil, intP(4), nil, nil, intP(2), intP(5), nil, intP(9), nil, nil, intP(6), nil, nil},
			wannaFindData: 0,
			exits:         false,
		},
	}
	for i, testCase := range testCases {
		tree := BuildFromPre(testCase.preSrc)
		node := tree.FindParent(testCase.wannaFindData)
		actuallyRes := node != nil
		if testCase.exits != actuallyRes {
			t.Fatalf("TestTree_FindParent : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, actuallyRes, testCase.exits)
		}
	}
}
func TestTree_Depth(t *testing.T) {
	c := func(a, b interface{}) (res int) {
		ai, _ := a.(int)
		bi, _ := b.(int)
		if ai == bi {
			return 0
		}
		return -1
	}
	testCases := []struct {
		preSrc   []*int
		wannaRes int
		comp     compare
	}{
		{
			preSrc: []*int{intP(0), intP(1), intP(3), intP(7), nil, nil, intP(8), nil, nil, intP(4), nil, nil, intP(2), intP(5), nil, intP(9), nil, nil, intP(6), nil, nil},
			//"HIDEBJFGCA"
			wannaRes: 3,
			comp:     c,
		},
	}
	for i, testCase := range testCases {
		tree := BuildFromPre(testCase.preSrc)
		actuallyRes := tree.Depth()
		c := testCase.comp(testCase.wannaRes, actuallyRes)
		if c != 0 {
			t.Fatalf("TestTree_Depth : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, actuallyRes, testCase.wannaRes)
		}
	}
}

func TestBuildFromPre(t *testing.T) {
	c := func(a, b interface{}) (res int) {
		ai, _ := a.([]int)
		bi, _ := b.([]int)
		if len(ai) != len(bi) {
			return -1
		}
		for i, ain := range ai {
			if bi[i] != ain {
				return -1
			}
		}
		return 0
	}
	testCases := []struct {
		preSrc      []*int
		wannaMidRes []int
		comp        compare
	}{
		{
			preSrc:      []*int{intP(0), intP(1), intP(3), intP(7), nil, nil, intP(8), nil, nil, intP(4), nil, nil, intP(2), intP(5), nil, intP(9), nil, nil, intP(6), nil, nil},
			wannaMidRes: []int{7, 3, 8, 1, 4, 0, 5, 9, 2, 6},
			comp:        c,
		},
	}
	for i, testCase := range testCases {
		tree := BuildFromPre(testCase.preSrc)
		actuallyRes := tree.Mid()
		c := testCase.comp(testCase.wannaMidRes, actuallyRes)
		if c != 0 {
			t.Fatalf("TestBuildFromPreMid : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, actuallyRes, testCase.wannaMidRes)
		}
	}
}

func TestBuildFromPreMid(t *testing.T) {
	c := func(a, b interface{}) (res int) {
		ai, _ := a.([]int)
		bi, _ := b.([]int)
		if len(ai) != len(bi) {
			return -1
		}
		for i, ain := range ai {
			if bi[i] != ain {
				return -1
			}
		}
		return 0
	}
	testCases := []struct {
		preSrc       []int
		midSrc       []int
		wannaPostRes []int
		comp         compare
	}{
		{
			preSrc:       []int{0, 1, 3, 7, 8, 4, 2, 5, 9, 6},
			midSrc:       []int{7, 3, 8, 1, 4, 0, 5, 9, 2, 6},
			wannaPostRes: []int{7, 8, 3, 4, 1, 9, 5, 6, 2, 0},
			comp:         c,
		},
	}
	for i, testCase := range testCases {
		tree := BuildFromPreMid(testCase.preSrc, testCase.midSrc)
		actuallyRes := tree.Post()
		c := testCase.comp(testCase.wannaPostRes, actuallyRes)
		if c != 0 {
			t.Fatalf("TestBuildFromPreMid : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, actuallyRes, testCase.wannaPostRes)
		}
	}
}

func TestTree_BFS(t *testing.T) {
	c := func(a, b interface{}) (res int) {
		ai, _ := a.([]int)
		bi, _ := b.([]int)
		if len(ai) != len(bi) {
			return -1
		}
		for i, ain := range ai {
			if bi[i] != ain {
				return -1
			}
		}
		return 0
	}
	testCases := []struct {
		preSrc   []*int
		wannaRes []int
		comp     compare
	}{
		{
			preSrc: []*int{intP(0), intP(1), intP(3), intP(7), nil, nil, intP(8), nil, nil, intP(4), nil, nil, intP(2), intP(5), nil, intP(9), nil, nil, intP(6), nil, nil},
			//"HIDEBJFGCA"
			wannaRes: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			comp:     c,
		},
	}
	for i, testCase := range testCases {
		tree := BuildFromPre(testCase.preSrc)
		actuallyRes := tree.BFS()
		c := testCase.comp(testCase.wannaRes, actuallyRes)
		if c != 0 {
			t.Fatalf("TestTree_BFS : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, actuallyRes, testCase.wannaRes)
		}
	}
}

func TestTree_Rm(t *testing.T) {
	c := func(a, b interface{}) (res int) {
		ai, _ := a.([]int)
		bi, _ := b.([]int)
		if len(ai) != len(bi) {
			return -1
		}
		for i, ain := range ai {
			if bi[i] != ain {
				return -1
			}
		}
		return 0
	}
	testCases := []struct {
		preSrc   []*int
		rm       int
		wannaRes []int
		comp     compare
	}{
		{
			preSrc: []*int{intP(0), intP(1), intP(3), intP(7), nil, nil, intP(8), nil, nil, intP(4), nil, nil, intP(2), intP(5), nil, intP(9), nil, nil, intP(6), nil, nil},
			//"HIDEBJFGCA"
			rm:       0,
			wannaRes: nil,
			comp:     c,
		},
		{
			preSrc: []*int{intP(0), intP(1), intP(3), intP(7), nil, nil, intP(8), nil, nil, intP(4), nil, nil, intP(2), intP(5), nil, intP(9), nil, nil, intP(6), nil, nil},
			//"HIDEBJFGCA"
			rm:       3,
			wannaRes: []int{0, 1, 2, 4, 5, 6, 9},
			comp:     c,
		},
	}
	for i, testCase := range testCases {
		tree := BuildFromPre(testCase.preSrc)
		//tree.Print()
		tt := tree.Rm(testCase.rm)
		var actuallyRes []int
		if tt != nil {
			actuallyRes = tt.BFS()
			//tt.Print()
		}
		c := testCase.comp(testCase.wannaRes, actuallyRes)
		if c != 0 {
			t.Fatalf("TestTree_Rm : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, actuallyRes, testCase.wannaRes)
		}
	}
}

func TestTree_Width(t *testing.T) {
	testCases := []struct {
		preSrc   []*int
		wannaRes int
	}{
		{
			preSrc:   []*int{intP(0), intP(1), intP(3), intP(7), nil, nil, intP(8), nil, nil, intP(4), nil, nil, intP(2), intP(5), nil, intP(9), nil, nil, intP(6), nil, nil},
			wannaRes: 15,
		},
	}
	for i, testCase := range testCases {
		tree := BuildFromPre(testCase.preSrc)
		actuallyRes := tree.Width()
		if actuallyRes != testCase.wannaRes {
			t.Fatalf("TestTree_Width : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, actuallyRes, testCase.wannaRes)
		}
	}
}

func TestTree_Print(t *testing.T) {
	testCases := []struct {
		preSrc []*int
	}{
		{
			preSrc: []*int{intP(0), intP(1), intP(3), intP(7), nil, nil, intP(8), nil, nil, intP(4), nil, nil, intP(2), intP(5), nil, intP(9), nil, nil, intP(6), nil, nil},
		},
	}
	for _, testCase := range testCases {
		tree := BuildFromPre(testCase.preSrc)
		tree.Print()
	}
}
