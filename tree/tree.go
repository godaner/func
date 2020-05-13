package tree

type Tree interface {
	Find(data int) (exits bool)
	MaxDepth() (maxDep int)
	Min() (minN int)
	Max() (maxN int)
	Pre() (p []int)
	Mid() (m []int)
	Post() (p []int)
}
