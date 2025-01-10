package easy

import "github.com/showyquasar88/proj-combine/algorithm/common/list"

// mergeTwoLists 合并两个有序链表
func mergeTwoLists(list1 *list.ListNode, list2 *list.ListNode) *list.ListNode {
	dummy := &list.ListNode{}
	cur := dummy
	for list1 != nil || list2 != nil {
		if list1 == nil {
			cur.Next = list2
			break
		}
		if list2 == nil {
			cur.Next = list1
			break
		}
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}
