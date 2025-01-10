package classical_150

import "github.com/showyquasar88/proj-combine/algorithm/common"

// isSymmetric 对称二叉树
func isSymmetric(root *common.TreeNode) bool {
	if root == nil {
		return false
	}
	return isSymmetricProcess(root.Left, root.Right)
}

func isSymmetricProcess(left *common.TreeNode, right *common.TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left == nil || right == nil) || (left.Val != right.Val) {
		return false
	}
	return isSymmetricProcess(left.Left, right.Right) && isSymmetricProcess(left.Right, right.Left)
}
