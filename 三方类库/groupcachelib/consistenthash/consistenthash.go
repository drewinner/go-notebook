package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

//peer 代表一个缓存服务器

//hash函数、如果不定义crc32.ChecksumIEEE
type Hash func(data []byte) uint32

//一致性hash
type Map struct {
	hash     Hash           //hash函数
	replicas int            //每个peer节点会产生几个虚拟节点
	keys     []int          //存储所有虚拟节点的hash值
	hashMap  map[int]string //根据hash值找到对应的peer
}

/**
*	新建hash环
*	@param:replicas int 每个节点有几个虚拟节点
*	@fn :hash 函数、如果为空使用crc32.ChecksumIEEE
 */
func New(replicas int, fn Hash) *Map {
	m := &Map{
		hash:     fn,
		replicas: replicas,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}
func (m *Map) IsEmpty() bool {
	return len(m.keys) == 0
}

/**
*	将一些keys加入到hash中
*	keys 目前看是机器pro://ip:port
 */
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		//一个key就是一个peer节点
		//对于每个key都要生成replicas个虚拟节点、并把虚拟节点和key(真正节点)的对应关系存起来
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	//排序、利用二分法查找
	sort.Ints(m.keys)
}

/**
*	这边的key为一个cache中的key，即要查询的key
*	计算这个key的hash值，找到与之最为接近的虚拟节点，再通过虚拟节点找到具体的peer
 */
func (m *Map) Get(key string) string {
	//1. 计算key对应的hash值
	hash := int(m.hash([]byte(key)))
	//2. 使用二分法查找key对应的hash最近的虚拟节点
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	//3.一致性hash的虚拟节点应该构成一个环
	if idx == len(m.keys) {
		idx = 0
	}
	return m.hashMap[m.keys[idx]]
}
