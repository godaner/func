package avl

import (
	"testing"
)

// compare
// 1 : a > b
// 0 : a = b
// -1 : a < b
type compare func(a, b interface{}) (res int)

func TestBuild(t *testing.T) {
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
		src      []int
		wannaRes []int
		comp     compare
	}{
		{
			src:      []int{0, 1, 3, 7, 8, 4, 2, 5, 9, 6},
			wannaRes: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			comp:     c,
		},
		{
			src:      []int{1, 0, 3, 7, 8, 4, 2, 5, 9, 6, -1},
			wannaRes: []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			comp:     c,
		},
		{
			src:      []int{-1, -2, 0, 1},
			wannaRes: []int{-2, -1, 0, 1},
			comp:     c,
		},
	}

	for i, testCase := range testCases {
		tree := Build(testCase.src)
		actuallyPreRes := tree.Mid()
		c := testCase.comp(testCase.wannaRes, actuallyPreRes)
		if c != 0 {
			t.Fatalf("TestBuild : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, actuallyPreRes, testCase.wannaRes)
		}
	}
}

func TestTree_Find(t *testing.T) {

	testCases := []struct {
		src           []int
		exits         bool
		wannaFindDate int
	}{
		{
			src:           []int{0, 1, 3, 7, 8, 4, 2, 5, 9, 6},
			wannaFindDate: 8,
			exits:         true,
		},
		{
			src:           []int{0, 1, 3, 7, 8, 4, 2, 5, 9, 6},
			wannaFindDate: -1,
			exits:         false,
		},
	}
	for i, testCase := range testCases {
		tree := Build(testCase.src)
		node := tree.Find(testCase.wannaFindDate)
		actuallyRes := node != nil
		if testCase.exits != actuallyRes && testCase.wannaFindDate == node.Data() {
			t.Fatalf("TestTree_Find : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, actuallyRes, testCase.exits)
		}
	}
}

func TestTree_FindParent(t *testing.T) {

	testCases := []struct {
		src           []int
		exits         bool
		wannaFindDate int
	}{
		{
			src:           []int{0, 1, 3, 7, 8, 4, 2, 5, 9, 6},
			wannaFindDate: 7,
			exits:         true,
		},
		{
			src:           []int{0, 1, 3, 7, 8, 4, 2, 5, 9, 6},
			wannaFindDate: 1,
			exits:         true,
		},
	}
	for i, testCase := range testCases {
		tree := Build(testCase.src)
		node := tree.FindParent(testCase.wannaFindDate)
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
		preSrc   []int
		wannaRes int
		comp     compare
	}{
		{
			preSrc:   []int{0, -1, 3, 7, 8, 4, 2, 5, 9, 6},
			wannaRes: 3,
			comp:     c,
		},
	}
	for i, testCase := range testCases {
		tree := Build(testCase.preSrc)
		actuallyRes := tree.Depth()
		c := testCase.comp(testCase.wannaRes, actuallyRes)
		if c != 0 {
			t.Fatalf("TestTree_Depth : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, actuallyRes, testCase.wannaRes)
		}
	}
}

func TestTree_Max(t *testing.T) {
	testCases := []struct {
		preSrc   []int
		wannaRes int
	}{
		{
			preSrc:   []int{0, 1, 3, 7, 8, 4, 2, 5, 9, 6},
			wannaRes: 9,
		},
	}
	for i, testCase := range testCases {
		tree := Build(testCase.preSrc)
		node := tree.Max()
		if node == nil {
			t.Fatalf("TestTree_Max : testCase[%v] fail , res is : nil !", i)
		}
		actuallyRes := node.Data()
		if testCase.wannaRes != actuallyRes {
			t.Fatalf("TestTree_Max : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, actuallyRes, testCase.wannaRes)
		}
	}
}

func TestTree_Min(t *testing.T) {
	testCases := []struct {
		preSrc   []int
		wannaRes int
	}{
		{
			preSrc:   []int{0, 1, 3, 7, 8, 4, 2, 5, 9, 6},
			wannaRes: 0,
		},
	}
	for i, testCase := range testCases {
		tree := Build(testCase.preSrc)
		node := tree.Min()
		if node == nil {
			t.Fatalf("TestTree_Min : testCase[%v] fail , res is : nil !", i)
		}
		actuallyRes := node.Data()
		if testCase.wannaRes != actuallyRes {
			t.Fatalf("TestTree_Min : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, actuallyRes, testCase.wannaRes)
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
		preSrc   []int
		wannaRes []int
		comp     compare
	}{
		{
			preSrc:   []int{0, -1, 3, 7, 8, 4, 2, 5, 9, 6},
			wannaRes: []int{3, 0, 7, -1, 2, 5, 8, 4, 6, 9},
			comp:     c,
		},
	}
	for i, testCase := range testCases {
		tree := Build(testCase.preSrc)
		actuallyRes := tree.BFS()
		c := testCase.comp(testCase.wannaRes, actuallyRes)
		if c != 0 {
			t.Fatalf("TestTree_BFS : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, actuallyRes, testCase.wannaRes)
		}
	}
}

func TestTree_Width(t *testing.T) {
	testCases := []struct {
		preSrc   []int
		wannaRes int
	}{
		{
			preSrc:   []int{0, 1, 3, 7, 8, 4, 2, 5, 9, 6},
			wannaRes: 15,
		},
	}
	for i, testCase := range testCases {
		tree := Build(testCase.preSrc)
		actuallyRes := tree.Width()
		if actuallyRes != testCase.wannaRes {
			t.Fatalf("TestTree_Width : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, actuallyRes, testCase.wannaRes)
		}
	}
}

func TestTree_Print(t *testing.T) {
	testCases := []struct {
		preSrc []int
	}{
		{
			preSrc: []int{0, 1, 3, 7, 8, 4, 2, 5, 9, 6},
			//preSrc: []int{11, 12, 33, 71, 8, 43, 2, 5, 9, 6},
		},
		{
			preSrc: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			//preSrc: []int{11, 12, 33, 71, 8, 43, 2, 5, 9, 6},
		},
		{
			//preSrc: []int{0, 1, 3, 7, 8, 4, 2, 5, 9, 6},
			preSrc: []int{50, 30, 80, 20, 35, 34, 32, 40, 70, 75, 100},
		},
	}
	for _, testCase := range testCases {
		tree := Build(testCase.preSrc)
		tree.Print()
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
		preSrc   []int
		wannaRes []int
		wannaRm  int
		comp     compare
	}{
		{
			preSrc:   []int{50, 30, 20, 35, 34, 32, 40, 80, 75, 100},
			wannaRes: []int{34, 30, 20, 32, 80, 40, 35, 50, 100},
			wannaRm:  75,
			comp:     c,
		},
	}
	for i, testCase := range testCases {
		tree := Build(testCase.preSrc)
		tree = tree.Rm(testCase.wannaRm)
		actuallyRes := tree.Pre()
		if c(testCase.wannaRes, actuallyRes) != 0 {
			t.Fatalf("TestTree_Rm : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, actuallyRes, testCase.wannaRes)
		}
	}
}
