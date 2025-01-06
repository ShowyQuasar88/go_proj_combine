package classical_150

import "github.com/showyquasar88/proj-combine/algorithm/common"

// mergeTwoLists 合并两个有序链表
func mergeTwoLists(list1 *common.ListNode, list2 *common.ListNode) *common.ListNode {
	head := &common.ListNode{}
	cur := head
	for list1 != nil || list2 != nil {
		if list1 == nil {
			cur.Next = list2
			break
		}
		if list2 == nil {
			cur.Next = list1
			break
		}
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	return head.Next
}
