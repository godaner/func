package binary

import "container/list"

// Tree
type Tree struct {
	Left, Right *Tree
	Date        int
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
		val:=queue.Front()
		queue.Remove(val)

		node := val.Value.(*Tree)
		n = append(n, node.Date)

		//put child
		if node.Left!=nil{
			queue.PushBack(node.Left)
		}
		if node.Right!=nil{
			queue.PushBack(node.Right)
		}
	}
	return n
}

// Depth
//  树的最大深度
func (t *Tree) MaxDepth() (maxDep int) {
	curDep := 0
	depth(t, curDep, &maxDep)
	return maxDep
}
func depth(root *Tree, curDep int, maxDep *int) {
	if root == nil {
		return
	}
	curDep++
	if curDep > *maxDep {
		*maxDep = curDep
	}
	depth(root.Left, curDep, maxDep)
	depth(root.Left, curDep, maxDep)
}

// Find
func (t *Tree) Find(data int) (exits bool) {
	find(t, data, &exits)
	return exits
}
func find(root *Tree, data int, exits *bool) {
	if root == nil {
		return
	}
	if (*exits) == true {
		return
	}
	if data == root.Date {
		*exits = true
		return
	}
	find(root.Left, data, exits)
	find(root.Right, data, exits)
}

// Min
func (t *Tree) Min() (minN int) {
	min(t, &minN)
	return minN
}

// min
func min(root *Tree, minN *int) {
	if root == nil {
		return
	}
	if root.Date < *minN {
		*minN = root.Date
	}
	min(root.Left, minN)
	min(root.Right, minN)
}

// Max
func (t *Tree) Max() (maxN int) {
	max(t, &maxN)
	return maxN
}

// max
func max(root *Tree, maxN *int) {
	if root == nil {
		return
	}
	if root.Date > *maxN {
		*maxN = root.Date
	}
	max(root.Left, maxN)
	max(root.Right, maxN)
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
