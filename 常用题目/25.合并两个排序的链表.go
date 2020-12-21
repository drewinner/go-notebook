package 常用题目

/**
*	https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/
*	输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
*	输入：1->2->4, 1->3->4
*	输出：1->1->2->3->4->4
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	h := &ListNode{}
	cur := h
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			cur = l2
			l2 = l2.Next
		} else {
			cur = l1
			l1 = l1.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return h.Next
}
