package 常用题目

/**
*	https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/
*	定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
*	输入: 1->2->3->4->5->NULL
	输出: 5->4->3->2->1->NULL
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	for head != nil {
		//防丢失、先将节点保存
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}
	return pre

}
