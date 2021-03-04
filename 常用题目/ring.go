package 常用题目

type Ring struct {
	next,pre *Ring
	Value interface{}
}

/**
*	初始化环形链表
 */
func InitRing() (r *Ring) {
	r = new(Ring)
	r.next,r.pre = r,r
	return r
}


/**
*	创建N个节点的环形链表
 */
 func (r *Ring) NewRing(n int) bool {
	 if n <= 0 {
		 return false
	 }
	 //存放前一个节点
	 p:= r
	 for i:=0;i<n;i++ {
	 	p.next = &Ring{
			pre:   p,
			Value: i,
		}
	 	p = p.next
	 }
	 p.next,r.pre = r,p
	 return true
 }

func (r *Ring) DeleteNode(del int) error {
	t := r
	if r == nil {
		return nil
	}
	for {
		if t.Value == del {
			t.pre.next = t.next
			return nil
		}
		if t.next == r {
			return nil
		}
		t = t.next
	}
}

func (r *Ring) CountRing() int {
	t := r
	count := 0
	if r == nil {
		return 0
	}
	for {
		if t.next != nil {
			count++
		}
		if t.next == r {
			return count
		}
		t = t.next
	}
}