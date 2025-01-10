package classical_150

import "github.com/showyquasar88/proj-combine/algorithm/common"

// buildTreePost 从中序与后序遍历序列构造二叉树
func buildTreePost(inorder []int, postorder []int) *common.TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	mid := 0
	for inorder[mid] != postorder[len(postorder)-1] {
		mid++
	}
	root := &common.TreeNode{Val: postorder[len(postorder)-1]}
	root.Left = buildTreePost(inorder[:mid], postorder[:mid])
	root.Right = buildTreePost(inorder[mid+1:], postorder[mid:len(postorder)-1])
	return root
}
