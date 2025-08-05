package treeNode

type TreeNode struct {
	Value    int
	Children []*TreeNode
}

func (n *TreeNode) AddChild(child *TreeNode) {
	n.Children = append(n.Children, child)
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{
		Value:    value,
		Children: make([]*TreeNode, 0),
	}
}
