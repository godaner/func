package binary

// Tree
type Tree struct {
	Left, Right *Tree
	Date        int
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

// Build
//  通过前序遍历建立二叉树
//  ABDH##I##E##CF#J##G##
func Build(nodes []*int) (t *Tree) {
	i := -1 // 第几次输入
	return build(nodes, &i)
}

// build
//  nodes　总的输入
//  i　第几次输入
func build(nodes []*int, i *int) (t *Tree) {
	l := len(nodes)
	if l == 0 || l < *i {
		return nil
	}
	*i = (*i) + 1
	data := nodes[*i]
	if data == nil {
		return nil
	}
	node := newTreeNode(data)
	node.Left = build(nodes, i)
	node.Right = build(nodes, i)
	return node
}

// newTreeNode
func newTreeNode(data *int) (tn *Tree) {
	return &Tree{
		Left:  nil,
		Right: nil,
		Date:  *data,
	}
}
