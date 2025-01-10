package classical_150

import "github.com/showyquasar88/proj-combine/algorithm/common"

// maxDepth 二叉树的最大深度
func maxDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := maxDepth(root.Left), maxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}
