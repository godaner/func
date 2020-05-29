package bst

import (
	"container/list"
	"fmt"
	"github.com/godaner/func/tree"
)

// Tree
type Tree struct {
	Left, Right *Tree
	Data        int
	Weight      int64
}


func (t *Tree) Code(data int) (code string) {
	panic("implement me")
}

type WeightData struct {
	Data   int
	Weight int64
}

// Build
func Build(datas ...WeightData) (t tree.Tree) {
	if len(datas) <= 0 {
		return
	}
	// to nodes
	nodes := []*Tree{}
	for _, d := range datas {
		nodes = append(nodes, &Tree{
			Left:   nil,
			Right:  nil,
			Data:   d.Data,
			Weight: d.Weight,
		})
	}
	return build(nodes)
}

func build(nodes []*Tree) (t tree.Tree) {
	var min1,min2 *Tree
	for len(nodes) > 1 {
		min1, nodes = pickMin(nodes)
		min2, nodes = pickMin(nodes)
		nodes = append(nodes, &Tree{
			Left:   min1,
			Right:  min2,
			Data:   0,
			Weight: min1.Weight + min2.Weight,
		})
	}
	if len(nodes) == 0 {
		return nil
	}
	return nodes[0]
}

// pickMin
func pickMin(nodes []*Tree) (minNode *Tree, resNodes []*Tree) {
	if len(nodes) == 0 {
		return nil, nil
	}

	minIndex := 0
	minNode = nodes[0]
	for index, node := range nodes {
		if minNode.Compare(node) > 0 {
			minNode = node
			minIndex = index
		}
	}
	resNodes = []*Tree{}
	for index, node := range nodes {
		if index != minIndex {
			resNodes = append(resNodes, node)
		}
	}
	return minNode, resNodes
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
	(*ans)[h][mid] = fmt.Sprint(root.Data,"/",root.Weight)
	fill(root.Left, ans, h+1, l, mid-1)
	fill(root.Right, ans, h+1, mid+1, r)
}

func (t *Tree) Rm(data int) (tt tree.Tree) {
	if t == nil {
		return nil
	}
	if t.Data == data {
		return nil
	}
	rm(t, data)
	return t
}
func rm(root *Tree, data int) {
	// rec
	if root.Left != nil {
		if root.Left.Data == data {
			root.Left = nil
			return
		} else {
			rm(root.Left, data)
		}
	}
	if root.Right != nil {
		if root.Right.Data == data {
			root.Right = nil
			return
		} else {
			rm(root.Right, data)
		}
	}
}

func (t *Tree) Compare(tar tree.Tree) (res int) {
	if t.Weight == tar.(*Tree).Weight {
		return tree.TREE_COMPARE_SAME
	}
	if t.Weight  > tar.(*Tree).Weight {
		return tree.TREE_COMPARE_G
	}
	return tree.TREE_COMPARE_L
}

// Data
func (t *Tree) Elem() (data int) {
	if t == nil {
		return 0
	}
	return t.Data
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
	if root.Left != nil && root.Left.Data == data {
		return root
	}
	if root.Right != nil && root.Right.Data == data {
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
	if root.Data == data {
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
	if lc != nil && lc.Elem() < minN.Elem() {
		minN = lc
	}
	rc := min(root.Right)
	if rc != nil && rc.Elem() < minN.Elem() {
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
	if lc != nil && lc.Elem() > maxN.Elem() {
		maxN = lc
	}
	rc := max(root.Right)
	if rc != nil && rc.Elem() > maxN.Elem() {
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
	*p = append(*p, root.Data)
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
	*m = append(*m, root.Data)
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
	*p = append(*p, root.Data)

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
		Data:  data,
	}
}
