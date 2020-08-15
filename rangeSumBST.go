package range_sum_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	result := 0
	result = transversal(root, L, R)
	return result
}

func transversal(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	switch {
	case root.Val > R:
		return transversal(root.Left, L, R)
	case root.Val < L:
		return transversal(root.Right, L, R)
	default:
		return transversal(root.Left, L, R) + root.Val + transversal(root.Right, L, R)
	}
}
