package binary

import (
	"fmt"
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
func TestW2N(t *testing.T) {
	sss := []string{"ABDH##I##E##CF#J##G##"}
	for index, ss := range sss {
		fmt.Print(index, " : ")
		for i, _ := range ss {
			if ss[i] >= 65 {
				fmt.Print(ss[i] - 65)
			} else {
				fmt.Print("#")
			}
		}

	}
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
		tree := Build(testCase.preSrc)
		actuallyPreRes := tree.Pre()
		fmt.Println(tree.Mid())
		fmt.Println(tree.Post())
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
		tree := Build(testCase.preSrc)
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
		tree := Build(testCase.preSrc)
		actuallyPreRes := tree.Post()
		c := testCase.comp(testCase.wannaRes, actuallyPreRes)
		if c != 0 {
			t.Fatalf("TestTree_Post : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, actuallyPreRes, testCase.wannaRes)
		}
	}
}

func TestTree_Max(t *testing.T) {
	c := func(a, b interface{}) (res int) {
		ai, _ := a.(int)
		bi, _ := b.(int)
		if ai==bi{
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
			wannaRes: 9,
			comp:     c,
		},
	}
	for i, testCase := range testCases {
		tree := Build(testCase.preSrc)
		actuallyRes := tree.Max()
		c := testCase.comp(testCase.wannaRes, actuallyRes)
		if c != 0 {
			t.Fatalf("TestTree_Max : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, actuallyRes, testCase.wannaRes)
		}
	}
}

func TestTree_Min(t *testing.T) {
	c := func(a, b interface{}) (res int) {
		ai, _ := a.(int)
		bi, _ := b.(int)
		if ai==bi{
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
			wannaRes: 0,
			comp:     c,
		},
	}
	for i, testCase := range testCases {
		tree := Build(testCase.preSrc)
		actuallyRes := tree.Min()
		c := testCase.comp(testCase.wannaRes, actuallyRes)
		if c != 0 {
			t.Fatalf("TestTree_Min : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, actuallyRes, testCase.wannaRes)
		}
	}
}