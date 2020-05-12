package binary

import (
	"testing"
)

// compare
// 1 : a > b
// 0 : a = b
// -1 : a < b
type compare func(a, b interface{}) (res int)

func TestTree_Pre(t *testing.T) {
	c := func(a, b interface{}) (res int) {
		ai, _ := a.(string)
		bi, _ := b.(string)
		if ai == bi {
			return 0
		}
		return -1
	}
	testCases := []struct {
		preSrc   string
		wannaRes string
		comp     compare
	}{
		{
			preSrc:   "ABDH##I##E##CF#J##G##",
			wannaRes: "ABDHIECFJG",
			comp:     c,
		},
	}
	for i, testCase := range testCases {
		tree := Build(testCase.preSrc)
		actuallyPreRes := tree.Pre()
		c := testCase.comp(testCase.wannaRes, actuallyPreRes)
		if c != 0 {
			t.Fatalf("TestTree_Pre : testCase[%v] fail , actually res is : %v , wanna res is : %v !", i, actuallyPreRes, testCase.wannaRes)
		}
	}
}
func TestTree_Mid(t *testing.T) {
	c := func(a, b interface{}) (res int) {
		ai, _ := a.(string)
		bi, _ := b.(string)
		if ai == bi {
			return 0
		}
		return -1
	}
	testCases := []struct {
		preSrc   string
		wannaRes string
		comp     compare
	}{
		{
			preSrc:   "ABDH##I##E##CF#J##G##",
			wannaRes: "HDIBEAFJCG",
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
		ai, _ := a.(string)
		bi, _ := b.(string)
		if ai == bi {
			return 0
		}
		return -1
	}
	testCases := []struct {
		preSrc   string
		wannaRes string
		comp     compare
	}{
		{
			preSrc:   "ABDH##I##E##CF#J##G##",
			wannaRes: "HIDEBJFGCA",
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
