package binary

import (
	"strings"
)

const (
	nilNode = "#"
)

// Tree
type Tree struct {
	Left, Right *Tree
	Date        string
}

// Pre
func (t *Tree) Pre() (p string) {
	pre(t, &p)
	return p
}

// pre
func pre(root *Tree, p *string) {
	if root == nil {
		return
	}
	*p = *p + root.Date
	pre(root.Left, p)
	pre(root.Right, p)

}
// Mid
func (t *Tree) Mid() (m string) {
	mid(t, &m)
	return m
}

// mid
func mid(root *Tree, m *string) {
	if root == nil {
		return
	}
	mid(root.Left, m)
	*m = *m + root.Date
	mid(root.Right, m)

}
// Post
func (t *Tree) Post() (p string) {
	post(t, &p)
	return p
}

// post
func post(root *Tree, p *string) {
	if root == nil {
		return
	}
	post(root.Left, p)
	post(root.Right, p)
	*p = *p + root.Date

}
// Build
//  通过前序遍历建立二叉树
//  ABDH##I##E##CF#J##G##
func Build(nodes string) (t *Tree) {
	ns := strings.Split(nodes, "")
	i := -1 // 第几次输入
	return build(ns, &i)
}

// build
//  nodes　总的输入
//  i　第几次输入
func build(nodes []string, i *int) (t *Tree) {
	l := len(nodes)
	if l == 0 || l < *i {
		return nil
	}
	*i = (*i) + 1
	data := nodes[*i]
	if data == nilNode {
		return nil
	}
	node := newTreeNode(data)
	node.Left = build(nodes, i)
	node.Right = build(nodes, i)
	return node
}

// newTreeNode
func newTreeNode(data string) (tn *Tree) {
	return &Tree{
		Left:  nil,
		Right: nil,
		Date:  data,
	}
}
