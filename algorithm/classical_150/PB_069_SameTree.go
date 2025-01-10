package classical_150

import "github.com/showyquasar88/proj-combine/algorithm/common"

// isSameTree 相同的树
func isSameTree(p *common.TreeNode, q *common.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil || q == nil) || (p.Val != q.Val) {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
