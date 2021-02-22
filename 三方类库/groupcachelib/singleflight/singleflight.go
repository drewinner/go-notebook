package singleflight

import "sync"

type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	//有可能修改m、所以先上锁
	g.mu.Lock()
	//如果m为nil、则初始化一个
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	//如果m中存在对应的key的请求、则线程不会直接再次访问key、所以释放锁
	//然后阻塞等待已经存在的请求得到的结果
	if c, ok := g.m[key]; ok {
		//解锁
		g.mu.Unlock()
		//阻塞
		c.wg.Wait()
		//如果存在的请求已经完成、则阻塞状态会解除、得到正确结果
		return c.val, c.err
	}
	//如果不存在对该key的请求、则本线程要进行实际的请求、保持m的锁定状态
	//创建一个实际请求结构体
	c := new(call)
	//为了保证其他的相同请求的堵塞
	c.wg.Add(1)
	//组织好映射关系
	g.m[key] = c
	//解锁m
	g.mu.Unlock()
	//执行真正的请求函数，得到对该key请求的结果
	c.val, c.err = fn()
	//得到结果后取消其他请求的堵塞
	c.wg.Done()
	//该次请求完成后，要从已存在请求map中删掉
	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()
	//返回请求结果
	return c.val, c.err
}
