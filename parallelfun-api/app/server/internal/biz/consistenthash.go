package biz

import (
	"crypto/sha1"
	"encoding/binary"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/go-kratos/kratos/v2/registry"
)

// ConsistentHash 一致性哈希实现（适配Kratos的服务实例）
type ConsistentHash struct {
	replicas     int                                  // 虚拟节点数量
	circle       map[uint32]*registry.ServiceInstance // 虚拟节点哈希 -> 实际节点（格式：ip:port）
	sortedHashes []uint32                             // 排序的哈希环
	mu           sync.RWMutex
}

func NewConsistentHash(replicas int) *ConsistentHash {
	return &ConsistentHash{
		replicas: replicas,
		circle:   make(map[uint32]*registry.ServiceInstance),
	}
}

// 计算哈希值（复用之前的实现）
func (c *ConsistentHash) hashKey(key string) uint32 {
	h := sha1.New()
	h.Write([]byte(key))
	return binary.BigEndian.Uint32(h.Sum(nil))
}

// Update 根据Kratos服务实例更新哈希环
func (c *ConsistentHash) Update(instances []*registry.ServiceInstance) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// 清空旧节点
	c.circle = make(map[uint32]*registry.ServiceInstance)
	c.sortedHashes = []uint32{}

	// 添加新节点（从Kratos实例中提取ip:port）
	for _, ins := range instances {
		var node string // Kratos实例的Endpoint格式为 "ip:port"
		for _, endpoint := range ins.Endpoints {
			if strings.HasPrefix(endpoint, "grpc") {
				u, err := url.Parse(endpoint)
				if err != nil {
					// 处理解析错误，或者跳过该 endpoint
					continue
				}
				node = u.Host
				break
			}
		}
		for i := 0; i < c.replicas; i++ {
			replicaKey := node + ":" + strconv.FormatInt(int64(i), 10)
			hash := c.hashKey(replicaKey)
			c.circle[hash] = ins
			c.sortedHashes = append(c.sortedHashes, hash)
		}
	}

	// 排序哈希环
	sort.Slice(c.sortedHashes, func(i, j int) bool {
		return c.sortedHashes[i] < c.sortedHashes[j]
	})
}

// Get 根据key选择节点
func (c *ConsistentHash) Get(key string) (*registry.ServiceInstance, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if len(c.circle) == 0 {
		return nil, false
	}

	hash := c.hashKey(key)
	// 二分查找
	idx := sort.Search(len(c.sortedHashes), func(i int) bool {
		return c.sortedHashes[i] >= hash
	})
	if idx == len(c.sortedHashes) {
		idx = 0
	}

	return c.circle[c.sortedHashes[idx]], true
}
