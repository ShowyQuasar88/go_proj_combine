package classical_150

import "github.com/showyquasar88/proj-combine/algorithm/common"

// addTwoNumbers 两数相加
func addTwoNumbers(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	head, in := &common.ListNode{}, 0
	head.Next = l1
	cur := head
	for l1 != nil && l2 != nil {
		cur.Next.Val = l1.Val + l2.Val + in
		if cur.Next.Val >= 10 {
			cur.Next.Val -= 10
			in = 1
		} else {
			in = 0
		}
		l1 = l1.Next
		l2 = l2.Next
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
		for l1 != nil {
			cur.Next.Val = l1.Val + in
			if cur.Next.Val >= 10 {
				cur.Next.Val -= 10
				in = 1
			} else {
				in = 0
			}
			l1 = l1.Next
			cur = cur.Next
		}
	}
	if l2 != nil {
		cur.Next = l2
		for l2 != nil {
			cur.Next.Val = l2.Val + in
			if cur.Next.Val >= 10 {
				cur.Next.Val -= 10
				in = 1
			} else {
				in = 0
			}
			l2 = l2.Next
			cur = cur.Next
		}
	}
	if in != 0 {
		cur.Next = &common.ListNode{
			Val:  in,
			Next: nil,
		}
	}
	return head.Next
}
