package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32

// Map 包含了所有的key的哈希
// Map 是一致性哈希的主要数据结构，
type Map struct {
	hash     Hash           // 允许替换自定义的 Hash 函数，默认采用 crc32.ChecksumIEEE 算法
	replicas int            // 虚拟节点的倍数
	keys     []int          // sorted
	hashMap  map[int]string // 虚拟节点与真实节点的映射表，键值-虚拟节点的哈希值，值-真实节点的名称
}

func NewMap(replicas int, fn Hash) *Map {
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

// 新增真实节点/机器
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key))) // 计算虚拟节点的哈希值
			m.keys = append(m.keys, hash)                      // 将虚拟节点哈希值放置环上
			m.hashMap[hash] = key                              // 虚拟节点与真实节点映射
		}
		//fmt.Println(m.hashMap)
	}
	// 环上哈希值排序
	sort.Ints(m.keys)
}

// 选择节点
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))
	// 二分查找
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})

	return m.hashMap[m.keys[idx%len(m.keys)]]
}
