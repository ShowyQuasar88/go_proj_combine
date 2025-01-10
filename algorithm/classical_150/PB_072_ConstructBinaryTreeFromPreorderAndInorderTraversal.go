package classical_150

import "github.com/showyquasar88/proj-combine/algorithm/common"

// buildTree 从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *common.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	mid, root := 0, &common.TreeNode{Val: preorder[0]}
	for inorder[mid] != preorder[0] {
		mid++
	}
	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
	return root
}
