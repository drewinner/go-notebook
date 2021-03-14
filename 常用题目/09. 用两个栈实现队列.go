package 常用题目

import "container/list"

/**
*	https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
*	栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
*	分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

 */

type CQueue struct {
	in, out *list.List
}

func Constructor1() CQueue {
	return CQueue{
		in:  list.New(),
		out: list.New(),
	}
}

//添加元素
func (this *CQueue) AppendTail(value int) {
	this.in.PushBack(value)
}

//从队列中取出整数并删除
func (this *CQueue) DeleteHead() int {
	//如果第二个栈为空
	if this.out.Len() == 0 {
		for this.in.Len() > 0 {
			this.out.PushBack(this.in.Remove(this.in.Back()))
		}
	}
	if this.out.Len() != 0 {
		e := this.out.Back()
		this.out.Remove(e)
		return e.Value.(int)
	}
	return -1
}
