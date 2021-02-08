package lru

import "container/list"

type Cache struct {
	//元素最大个数
	MaxEntries int
	//当一个entries清除的时候、回调函数
	OnEvicted func(key Key, value interface{})
	//双向链表
	ll *list.List
	//map存放元素指针、便于索引到元素
	cache map[interface{}]*list.Element
}
type Key interface{}

type entry struct {
	key   Key
	value interface{}
}

/**
*	创建新的cache
 */
func New(maxEntries int) *Cache {
	return &Cache{
		MaxEntries: maxEntries,
		ll:         list.New(),
		cache:      make(map[interface{}]*list.Element),
	}
}

/**
*	添加元素到cache中
 */
func (c *Cache) Add(key Key, value interface{}) {
	if c.cache == nil {
		c.cache = make(map[interface{}]*list.Element)
		c.ll = list.New()
	}
	//判断元素是否存在、存在将访问的元素移动到链表头部、覆盖map内的值
	if ee, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ee)
		ee.Value.(*entry).value = value
		return
	}
	ele := c.ll.PushFront(&entry{
		key:   key,
		value: value,
	})
	c.cache[key] = ele
	if c.MaxEntries != 0 && c.ll.Len() > c.MaxEntries {
		c.RemoveOldest()
	}

}

/**
*	获取元素
 */
func (c *Cache) Get(key Key) (value interface{}, ok bool) {
	if c.cache == nil {
		return
	}
	if ele, hit := c.cache[key]; hit {
		c.ll.MoveToFront(ele)
		return ele.Value.(*entry).value, true
	}
	return nil, false
}

/**
*	删除元素
 */
func (c *Cache) Remove(key Key) {
	if c.cache == nil {
		return
	}
	if ele, hit := c.cache[key]; hit {
		c.removeElement(ele)
	}
}

/**
*	删除最老元素
 */
func (c *Cache) RemoveOldest() {
	if c.cache == nil {
		return
	}
	//获取链表中最后一个元素、删除
	ele := c.ll.Back()
	if ele != nil {
		c.removeElement(ele)
	}
}

/**
*	移除元素
*	删除链表中元素、删除map中元素
 */
func (c *Cache) removeElement(e *list.Element) {
	c.ll.Remove(e)
	kv := e.Value.(*entry)
	delete(c.cache, kv.key)
	if c.OnEvicted != nil {
		c.OnEvicted(kv.key, kv.value)
	}

}

/**
*	返回元素个数
 */
func (c *Cache) Len() int {
	if c.cache == nil {
		return 0
	}
	return c.ll.Len()
}

/**
*	删除所有元素
 */
func (c *Cache) Clear() {
	if c.OnEvicted != nil {
		for _, e := range c.cache {
			kv := e.Value.(*entry)
			c.OnEvicted(kv.key, kv.value)
		}
	}
	c.cache = nil
	c.ll = nil
}
