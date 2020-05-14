package tree

const (
	TREE_COMPARE_SAME = 1
	TREE_COMPARE_L    = -1
	TREE_COMPARE_G    = 1
)

type Tree interface {
	Data() (data int)
	Compare(tar Tree) (res int)
	Find(data int) (r Tree)
	FindParent(data int) (r Tree)
	MaxDepth() (maxDep int)
	Min() (r Tree)
	Max() (r Tree)
	Pre() (p []int)
	Mid() (m []int)
	Post() (p []int)
	BFS() (p []int)
}
