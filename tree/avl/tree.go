package avl

import (
	"container/list"
	"fmt"
	"github.com/godaner/func/tree"
	"strconv"
)

type Tree struct {
	Left, Right *Tree
	Data        int
	High        int
}

// Build
func Build(datas []int) (t tree.Tree) {
	if len(datas) <= 0 {
		return
	}
	var root tree.Addable
	root = newTreeNode(datas[0])
	for i := 1; i < len(datas); i++ {
		root = root.Add(datas[i])
	}
	return root.(tree.Tree)
}

// newTreeNode
func newTreeNode(data int) (tn tree.Addable) {
	return &Tree{
		Left:  nil,
		Right: nil,
		Data:  data,
	}
}

func (t *Tree) Rm(data int) (tt tree.Tree) {
	if t == nil {
		return nil
	}

	// first , rm node
	if data < t.Data { // left child
		t.Left, _ = t.Left.Rm(data).(*Tree)
	} else if data > t.Data { // right child
		t.Right, _ = t.Right.Rm(data).(*Tree)
	} else { // ==
		if t.Left != nil && t.Right != nil {
			t.Data = t.Right.Min().Elem()
			t.Right, _ = t.Right.Rm(t.Data).(*Tree)
		} else if t.Left == nil {
			t = t.Right
		} else if t.Right == nil {
			t = t.Left
		}
	}

	// second , balance it
	t.refreshHigh()

	bf := t.getBalancedFactor() // bf
	if bf >= 2 {
		leftBF := t.Left.getBalancedFactor() // child bf
		if leftBF >= 0 {                     // LL
			t = t.rightRotation()
		} else { // LR
			t = t.leftRightRotation()
		}
	}
	if bf <= -2 {
		rightBF := t.Right.getBalancedFactor() // child bf
		if rightBF > 0 {                       // RL
			t = t.rightLeftRotation()
		} else { // RR
			t = t.leftRotation()
		}
	}
	return t
}

func (t *Tree) Add(data int) (tt tree.Addable) {
	if t == nil {
		return newTreeNode(data)
	}
	if data < t.Data {
		t.Left, _ = t.Left.Add(data).(*Tree)
		t.refreshHigh()
		bf := t.getBalancedFactor()
		if bf >= 2 || bf <= -2 { // 不平衡
			if data < t.Left.Elem() { // LL
				return t.rightRotation()
			} else { // LR
				return t.leftRightRotation()
			}
		}
	} else {
		t.Right, _ = t.Right.Add(data).(*Tree)
		t.refreshHigh()
		bf := t.getBalancedFactor()
		if bf >= 2 || bf <= -2 { // 不平衡
			if data > t.Right.Elem() { // RR
				return t.leftRotation()
			} else { // RL
				return t.rightLeftRotation()
			}
		}
	}
	return t
}

// getBalancedFactor
//   t.Left.getHigh() - t.Right.getHigh()
func (t *Tree) getBalancedFactor() (bf int) {
	if t == nil {
		return 0
	}
	return t.Left.getHigh() - t.Right.getHigh()
}
func (t *Tree) rightRotation() (r *Tree) {
	newRoot := t.Left
	t.Left = newRoot.Right
	newRoot.Right = t

	t.refreshHigh()
	newRoot.refreshHigh()
	return newRoot
}
func (t *Tree) leftRotation() (r *Tree) {
	newRoot := t.Right
	t.Right = newRoot.Left
	newRoot.Left = t

	t.refreshHigh()
	newRoot.refreshHigh()
	return newRoot
}
func (t *Tree) getHigh() (h int) {
	if t == nil {
		return -1
	}
	return t.High
}
func (t *Tree) refreshHigh() {
	if t == nil {
		return
	}
	t.High = maxN(t.Left.getHigh(), t.Right.getHigh()) + 1
}
func maxN(a, b int) (m int) {
	if a >= b {
		return a
	}
	return b
}
func (t *Tree) leftRightRotation() (r *Tree) {
	if t == nil {
		return nil
	}
	t.Left = t.Left.leftRotation()
	return t.rightRotation()
}
func (t *Tree) rightLeftRotation() (r *Tree) {
	if t == nil {
		return nil
	}
	t.Right = t.Right.rightRotation()
	return t.leftRotation()
}
func (t *Tree) Elem() (data int) {
	return t.Data
}

func (t *Tree) Width() (w int) {
	// 计算输出数组的规模
	h := uint(t.Depth() + 1)
	// 宽度（最底层的宽度）， 2^n - 1
	return (1 << h) - 1
}

func (t *Tree) Compare(tar tree.Tree) (res int) {
	if t == tar {
		return tree.TREE_COMPARE_SAME
	}
	if t.Elem() > tar.Elem() {
		return tree.TREE_COMPARE_G
	}
	return tree.TREE_COMPARE_L
}

func (t *Tree) Find(data int) (r tree.Tree) {

	return find(t, data)
}
func find(root *Tree, data int) (r tree.Tree) {
	if root == nil {
		return
	}
	if data == root.Data {
		return root
	}
	if data <= root.Data {
		r := find(root.Left, data)
		if r != nil {
			return r
		}
	} else {
		r := find(root.Right, data)
		if r != nil {
			return r
		}
	}
	return nil
}

func (t *Tree) FindParent(data int) (r tree.Tree) {
	return findParent(t, data)
}

// findParent
func findParent(root *Tree, data int) (r tree.Tree) {
	if root == nil {
		return nil
	}
	if root.Left != nil && root.Left.Data == data {
		return root
	}
	if root.Right != nil && root.Right.Data == data {
		return root
	}
	if data <= root.Data {
		lr := findParent(root.Left, data)
		if lr != nil {
			return lr
		}
	} else {
		rr := findParent(root.Right, data)
		if rr != nil {
			return rr
		}
	}
	return nil
}

func (t *Tree) Depth() (dep int) {
	return depth(t, -1)
}

func depth(root *Tree, curDep int) (dep int) {
	if root == nil {
		return curDep
	}
	curDep++
	return maxN(depth(root.Left, curDep), depth(root.Right, curDep))
}
func (t *Tree) Print() {
	printPos(getTreePos(t))
}
func printPos(pos [][]string) {
	for i := 0; i < len(pos); i++ {
		ts := pos[i]
		for j := 0; j < len(ts); j++ {
			if ts[j] == "" {
				fmt.Print(" ")
			} else {
				fmt.Print("", ts[j])
			}
		}
		fmt.Println()
	}
	fmt.Println("--------------")
}

// getTreePos
//  计算树的各个节点的位置
func getTreePos(root *Tree) [][]string {
	// 计算输出数组的规模
	h := uint(root.Depth() + 1)
	// 宽度（最底层的宽度）， 2^n - 1
	w := (1 << h) - 1
	// 构造答案二维数组（slice）
	ans := make([][]string, h)
	for i := range ans {
		ans[i] = make([]string, w)
	}
	fill(root, &ans, 0, 0, w-1)
	return ans
}

// fill 递归的填充。
//
// 对当前传入的数，如果是空树，则返回。
// 我们获得当前可填充范围，计算出中位数mid，在二维数组的[h][mid]填充节点值，
// 然后递归对左节点，右节点调用fill方法。
func fill(root *Tree, ans *[][]string, h, l, r int) {
	if root == nil {
		return
	}
	mid := (l + r) / 2
	(*ans)[h][mid] = strconv.Itoa(root.Data)
	fill(root.Left, ans, h+1, l, mid-1)
	fill(root.Right, ans, h+1, mid+1, r)
}

func (t *Tree) Min() (r tree.Tree) {
	if t == nil {
		return nil
	}
	min := t
	for {
		if min.Left == nil {
			return min
		}
		min = min.Left
	}
}

func (t *Tree) Max() (r tree.Tree) {
	if t == nil {
		return nil
	}
	max := t
	for {
		if max.Right == nil {
			return max
		}
		max = max.Right
	}
}

func (t *Tree) Pre() (p []int) {
	p = []int{}
	pre(t, &p)
	return p
}

// pre
func pre(root *Tree, p *[]int) {
	if root == nil {
		return
	}
	*p = append(*p, root.Data)
	pre(root.Left, p)
	pre(root.Right, p)

}

func (t *Tree) Mid() (m []int) {
	m = []int{}
	mid(t, &m)
	return m
}

// mid
func mid(root *Tree, m *[]int) {
	if root == nil {
		return
	}
	mid(root.Left, m)
	*m = append(*m, root.Data)
	mid(root.Right, m)

}
func (t *Tree) Post() (p []int) {
	p = []int{}
	post(t, &p)
	return p
}

// post
func post(root *Tree, p *[]int) {
	if root == nil {
		return
	}
	post(root.Left, p)
	post(root.Right, p)
	*p = append(*p, root.Data)

}

func (t *Tree) BFS() (n []int) {
	if t == nil {
		return
	}
	queue := list.New()
	queue.PushBack(t)
	for queue.Len() > 0 {
		// get father
		val := queue.Front()
		queue.Remove(val)

		node, _ := val.Value.(*Tree)
		n = append(n, node.Data)

		//put child
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return n
}
