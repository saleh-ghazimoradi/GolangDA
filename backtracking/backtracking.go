package backtracking

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func HasValidPath(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return backtrack(root, true)
}

func backtrack(node *TreeNode, valid bool) bool {
	if node == nil {
		return false
	}

	if node.Value == 0 {
		return false
	}

	if node.Left == nil && node.Right == nil {
		return valid
	}

	return backtrack(node.Left, valid) || backtrack(node.Right, valid)
}

func NewNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}
