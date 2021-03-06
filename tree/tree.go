package tree

const (
	TREE_COMPARE_SAME = 1
	TREE_COMPARE_L    = -1
	TREE_COMPARE_G    = 1
)

// Tree
type Tree interface {
	Elem() (data int)
	Width() (w int)
	Compare(tar Tree) (res int)
	Find(data int) (r Tree)
	FindParent(data int) (r Tree)
	Depth() (dep int)
	Rm(data int) (tt Tree)
	Print()
	Min() (r Tree)
	Max() (r Tree)
	Pre() (p []int)
	Mid() (m []int)
	Post() (p []int)
	BFS() (n []int)
}

// HuffmanTree
type HuffmanTree interface {
	Tree
	Code(data int) (code string)
	Codes() (cs map[int]string)
}

// Addable
type Addable interface {
	Add(data int) (tt Addable)
}
