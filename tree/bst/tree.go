package bst

import (
	"container/list"
	"fmt"
	"github.com/godaner/func/tree"
	"strconv"
)

// Tree
type Tree struct {
	Left, Right *Tree
	Date        int
}

// Rm
/*删除元素
*1、如果被删除结点只有一个子结点，就直接将A的子结点连至A的父结点上，并将A删除
*2、如果被删除结点有两个子结点，将该结点右子数内的最小结点取代A。
 */
func (t *Tree) Rm(data int) (tt tree.Tree) {

	return t.rm(data)
}
func (t *Tree) rm(data int) (tt *Tree) {
	if t == nil {
		return t
	}
	compare := data - t.Data()
	if compare < 0 {
		t.Left = t.Left.rm(data)
	} else if compare > 0 {
		t.Right = t.Right.rm(data)
	} else { //找到结点,删除结点
		if t.Left != nil && t.Right != nil {
			t.Date = t.Right.Min().Data()
			t.Right = t.Right.rm(t.Data())
		} else if t.Left != nil {
			t = t.Left
		} else {
			t = t.Right
		}
	}
	return t
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
	(*ans)[h][mid] = strconv.Itoa(root.Date)
	fill(root.Left, ans, h+1, l, mid-1)
	fill(root.Right, ans, h+1, mid+1, r)
}

func (t *Tree) Data() (data int) {
	if t == nil {
		return 0
	}
	return t.Date
}

func (t *Tree) Width() (w int) {
	// 计算输出数组的规模
	h := uint(t.Depth() + 1)
	// 宽度（最底层的宽度）， 2^n - 1
	return (1 << h) - 1
}
func maxN(a, b int) (m int) {
	if a >= b {
		return a
	}
	return b
}

func (t *Tree) Compare(tar tree.Tree) (res int) {
	if t == tar {
		return tree.TREE_COMPARE_SAME
	}
	if t.Data() == tar.Data() {
		return tree.TREE_COMPARE_SAME
	}
	return tree.TREE_COMPARE_L
}

// Depth
//  树的深度
func (t *Tree) Depth() (maxDep int) {
	return depth(t, -1)
}

func depth(root *Tree, curDep int) (dep int) {
	if root == nil {
		return curDep
	}
	curDep++
	return max(depth(root.Left, curDep), depth(root.Right, curDep))
}
func max(a, b int) (m int) {
	if a >= b {
		return a
	}
	return b
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

		node := val.Value.(*Tree)
		n = append(n, node.Date)

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

func (t *Tree) FindParent(data int) (r tree.Tree) {
	return findParent(t, data)
}

// findParent
func findParent(root *Tree, data int) (r tree.Tree) {
	if root == nil {
		return nil
	}
	if root.Left != nil && root.Left.Date == data {
		return root
	}
	if root.Right != nil && root.Right.Date == data {
		return root
	}
	if data <= root.Date {
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

// Find
func (t *Tree) Find(data int) (r tree.Tree) {

	return find(t, data)
}
func find(root *Tree, data int) (r tree.Tree) {
	if root == nil {
		return
	}
	if data == root.Date {
		return root
	}
	if data <= root.Date {
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

// Add
func (t *Tree) Add(data int) {
	if t == nil {
		return
	}
	add(t, data)
}

// add
func add(curt *Tree, data int) {
	// Left
	if data <= curt.Date {
		if curt.Left == nil {
			curt.Left = newTreeNode(data)
			return
		}
		add(curt.Left, data)
		return
	}
	// Right
	if curt.Right == nil {
		curt.Right = newTreeNode(data)
		return
	}
	add(curt.Right, data)
}

// Build
func Build(datas []int) (t tree.Tree) {
	if len(datas) <= 0 {
		return
	}
	root := newTreeNode(datas[0])
	for i := 1; i < len(datas); i++ {
		root.Add(datas[i])
	}
	return root
}

// newTreeNode
func newTreeNode(data int) (tn *Tree) {
	return &Tree{
		Left:  nil,
		Right: nil,
		Date:  data,
	}
}

// Pre
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
	*p = append(*p, root.Date)
	pre(root.Left, p)
	pre(root.Right, p)

}

// Mid
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
	*m = append(*m, root.Date)
	mid(root.Right, m)

}

// Post
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
	*p = append(*p, root.Date)

}
