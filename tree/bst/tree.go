package bst

// Tree
type Tree struct {
	Left, Right *Tree
	Date        int
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
