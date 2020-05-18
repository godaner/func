package binary

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

func (t *Tree) Add(data int) (tt tree.Tree) {
	panic("implement me")
}

func (t *Tree) Width() (w int) {
	// 计算输出数组的规模
	h := uint(t.Depth() + 1)
	// 宽度（最底层的宽度）， 2^n - 1
	return (1 << h) - 1
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

func (t *Tree) Rm(data int) (tt tree.Tree) {
	if t == nil {
		return nil
	}
	if t.Date == data {
		return nil
	}
	rm(t, data)
	return t
}
func rm(root *Tree, data int) {
	// rec
	if root.Left != nil {
		if root.Left.Date == data {
			root.Left = nil
			return
		} else {
			rm(root.Left, data)
		}
	}
	if root.Right != nil {
		if root.Right.Date == data {
			root.Right = nil
			return
		} else {
			rm(root.Right, data)
		}
	}
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

// Data
func (t *Tree) Data() (data int) {
	if t == nil {
		return 0
	}
	return t.Date
}

// BFS
//  广度优先遍历
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
	return maxN(depth(root.Left, curDep), depth(root.Right, curDep))
}
func maxN(a, b int) (m int) {
	if a >= b {
		return a
	}
	return b
}

// Find
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
	lr := findParent(root.Left, data)
	if lr != nil {
		return lr
	}
	rr := findParent(root.Right, data)
	if rr != nil {
		return rr
	}
	return nil
}

// Find
func (t *Tree) Find(data int) (r tree.Tree) {
	return find(t, data)
}

// find
func find(root *Tree, data int) (r tree.Tree) {
	if root == nil {
		return nil
	}
	if root.Date == data {
		return root
	}

	lr := find(root.Left, data)
	if lr != nil {
		return lr
	}
	rr := find(root.Right, data)
	if rr != nil {
		return rr
	}
	return nil
}

// Min
func (t *Tree) Min() (r tree.Tree) {
	return min(t)
}

// min
func min(root *Tree) (r tree.Tree) {
	if root == nil {
		return
	}
	var minN tree.Tree
	minN = root
	lc := min(root.Left)
	if lc != nil && lc.Data() < minN.Data() {
		minN = lc
	}
	rc := min(root.Right)
	if rc != nil && rc.Data() < minN.Data() {
		minN = rc
	}
	return minN
}

// Max
func (t *Tree) Max() (r tree.Tree) {
	return max(t)
}

// max
func max(root *Tree) (r tree.Tree) {
	if root == nil {
		return
	}
	var maxN tree.Tree
	maxN = root
	lc := max(root.Left)
	if lc != nil && lc.Data() > maxN.Data() {
		maxN = lc
	}
	rc := max(root.Right)
	if rc != nil && rc.Data() > maxN.Data() {
		maxN = rc
	}
	return maxN
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

// BuildFromPreMid
//  通过先序中序遍历建立二叉树
//  pre  [0 1 3 7 8 4 2 5 9 6]
//  mid  [7 3 8 1 4 0 5 9 2 6]
//  post [7 8 3 4 1 9 5 6 2 0]
func BuildFromPreMid(pre []int, mid []int) (t *Tree) {
	if len(pre) == 0 || len(mid) == 0 {
		return nil
	}
	d := pre[0]
	dp := pos(mid, d)
	root := newTreeNode(d)
	root.Left = BuildFromPreMid(pre[1:dp+1], mid[:dp])
	root.Right = BuildFromPreMid(pre[dp+1:], mid[dp+1:])
	return root
}

// pos
func pos(nums []int, num int) (p int) {
	for index, n := range nums {
		if num == n {
			return index
		}
	}
	return p
}

// BuildFromPre
//  通过前序遍历建立二叉树
//  ABDH##I##E##CF#J##G##
func BuildFromPre(nodes []*int) (t *Tree) {
	i := -1 // 第几次输入
	return buildFromPre(nodes, &i)
}

// buildFromPre
//  nodes　总的输入
//  i　第几次输入
func buildFromPre(nodes []*int, i *int) (t *Tree) {
	l := len(nodes)
	if l == 0 || l < *i {
		return nil
	}
	*i = (*i) + 1
	data := nodes[*i]
	if data == nil {
		return nil
	}
	node := newTreeNode(*data)
	node.Left = buildFromPre(nodes, i)
	node.Right = buildFromPre(nodes, i)
	return node
}

// newTreeNode
func newTreeNode(data int) (tn *Tree) {
	return &Tree{
		Left:  nil,
		Right: nil,
		Date:  data,
	}
}
