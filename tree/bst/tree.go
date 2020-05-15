package bst

import (
	"container/list"
	"github.com/godaner/func/tree"
)

// Tree
type Tree struct {
	Left, Right *Tree
	Date        int
}

func (t *Tree) Rm(data int) (tt tree.Tree) {
	panic("implement me")
}

func (t *Tree) Data() (data int) {
	if t == nil {
		return 0
	}
	return t.Date
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

func (t *Tree) MaxDepth() (maxDep int) {
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
	// left
	if data <= curt.Date {
		if curt.Left == nil {
			curt.Left = newTreeNode(data)
			return
		}
		add(curt.Left, data)
		return
	}
	// right
	if curt.Right == nil {
		curt.Right = newTreeNode(data)
		return
	}
	add(curt.Right, data)
}

// Build
func Build(datas []int) (t *Tree) {
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
