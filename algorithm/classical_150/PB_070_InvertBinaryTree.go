package classical_150

import "github.com/showyquasar88/proj-combine/algorithm/common"

// invertTree 翻转二叉树
func invertTree(root *common.TreeNode) *common.TreeNode {
	if root == nil {
		return nil
	}
	left := root.Left
	root.Left = invertTree(root.Right)
	root.Right = invertTree(left)
	return root
}
