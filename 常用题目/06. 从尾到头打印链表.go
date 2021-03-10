package 常用题目

/**
*	https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
*	输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）
*	输入：head = [1,3,2]
*	输出：[2,3,1]
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	len := len(arr)
	for i, j := 0, len-1; i < j; {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return arr

}

var rs []int

func ReversePrint01(head *ListNode) []int {
	if head != nil {
		if head.Next != nil {
			ReversePrint01(head.Next)
		}
		rs = append(rs, head.Val)
	}
	return rs

}

//链表
func ReversePrint03(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	if head.Next == nil {
		return []int{head.Val}
	}
	var pre *ListNode
	lLen := 0
	for head != nil {
		lLen++
		//防丢失、保留
		tmpNode := head.Next
		//反转指针
		head.Next = pre
		//pre指针等于当前位置
		pre = head
		head = tmpNode
	}
	arr := make([]int, lLen)
	i := 0
	for pre != nil {
		arr[i] = pre.Val
		i++
		pre = pre.Next
	}
	return arr
}
