package classical_150

import (
	"github.com/showyquasar88/proj-combine/algorithm/common"
)

// copyRandomList 复制随机链表
func copyRandomList(head *common.Node) *common.Node {
	cur := head
	for cur != nil {
		newNode := &common.Node{Val: cur.Val}
		newNode.Next = cur.Next
		cur.Next = newNode
		cur = newNode.Next
	}
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	newHead, prev := &common.Node{}, head
	cur = newHead
	for prev != nil {
		cur.Next = prev.Next
		prev.Next = prev.Next.Next
		cur = cur.Next
		prev = prev.Next
	}
	return newHead.Next
}
