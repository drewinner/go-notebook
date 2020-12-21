package 常用题目

/**
*	https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/
*	给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
*	返回删除后的链表的头节点。
*	输入: head = [4,5,1,9], val = 5
*	输出: [4,1,9]
*	解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
 */
func DeleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	if head.Val == val {
		return head.Next
	}

	pre, cur := head, head.Next
	for cur != nil && cur.Val != val {
		pre, cur = cur, cur.Next
	}
	if cur != nil {
		pre.Next = cur.Next
	}
	return head
}
