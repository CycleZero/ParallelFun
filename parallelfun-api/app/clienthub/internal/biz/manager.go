package biz

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spaolacci/murmur3"
	"hash"
	"sync"
	"sync/atomic"
	"time"
)

type ConnManager struct {
	Shards     []*Shard
	ShardCount int
	//Hash32      hash.Hash32
	HashPool    sync.Pool
	ReceiveChan chan []byte
	resMapLock  sync.RWMutex
	resMap      map[string]chan []byte

	healthCheckInterval time.Duration

	ctx       context.Context
	closeFunc context.CancelFunc
	closeOnce sync.Once
	closed    atomic.Bool
}

func NewConnManager(ctx context.Context, shardCount int) *ConnManager {
	mctx, cancel := context.WithCancel(ctx)
	m := &ConnManager{
		ShardCount: shardCount,
		HashPool: sync.Pool{
			New: func() any {
				h := murmur3.New32()
				return h
			},
		},
		Shards:      NewShards(shardCount),
		ReceiveChan: make(chan []byte, 1024),
		resMap:      make(map[string]chan []byte),
		ctx:         mctx,
		closeFunc:   cancel,
	}
	go m.GoHandleReceiveRpcMsg()
	return m
}

func NewShards(num int) []*Shard {
	shards := make([]*Shard, num)
	for i := 0; i < num; i++ {
		shards[i] = &Shard{}
		shards[i].Conns = make(map[string]*ClientConn)

	}
	return shards
}

func (m *ConnManager) Close() {
	m.closeOnce.Do(func() {
		if m.closed.Load() == true {
			return
		}
		m.closed.Store(true)
		m.closeFunc()

	})
}

func (m *ConnManager) NewConn_Test(url string, secret string, clientId string) (*ClientConn, error) {
	shard := m.GetShard(clientId)
	shard.Mu.RLock()
	conn := shard.Conns[clientId]
	shard.Mu.RUnlock()
	if conn != nil {
		return conn, nil
	}
	shard.Mu.Lock()
	defer shard.Mu.Unlock()

	// 双重检查，避免重复创建
	conn = shard.Conns[clientId]
	if conn != nil {
		return conn, nil
	}

	// 创建新连接
	conn, err := NewClientConn(m.ctx, url, secret, clientId, m.ReceiveChan)
	if err != nil {
		return nil, fmt.Errorf("client Connect Failed: %v", err)
	}
	shard.Conns[clientId] = conn
	return conn, nil
}

type Shard struct {
	Mu    sync.RWMutex
	Conns map[string]*ClientConn // key: clientID
}

func (m *ConnManager) healthCheckLoop() {
	ticker := time.NewTicker(m.healthCheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-m.ctx.Done():
			return
		case <-ticker.C:
			m.performHealthCheck()
		}
	}
}

// 执行健康检查
func (m *ConnManager) performHealthCheck() {
	for _, shard := range m.Shards {
		shard.Mu.Lock()
		for clientID, conn := range shard.Conns {
			// 检查连接是否有效，这里需要 ClientConn 提供 IsAlive 方法
			if !conn.IsAlive() {
				delete(shard.Conns, clientID)
				conn.Close()
			}
		}
		shard.Mu.Unlock()
	}
}

func (m *ConnManager) SendRpcMsg(ctx context.Context, clientId string, req *RpcRequest) (chan []byte, error) {
	conn, err := m.GetConn(clientId)
	if err != nil {
		return nil, err
	}
	req.Id = conn.NextId()
	//TODO 大量请求下分配空间较多
	resChan := make(chan []byte, 1)
	//closeChan := make(chan byte, 1)
	msg, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	conn.WriteChan <- msg

	m.resMapLock.Lock()
	m.resMap[req.Id] = resChan
	m.resMapLock.Unlock()
	go func() {
		select {
		case <-ctx.Done():
			m.RemoveResChan(req.Id)
		case <-time.After(time.Second * 10):
			m.RemoveResChan(req.Id)
		}
	}()

	return resChan, nil
}

func (m *ConnManager) RemoveResChan(reqId string) {
	m.resMapLock.RLock()
	_, exist := m.resMap[reqId]
	m.resMapLock.RUnlock()
	if !exist {
		return
	}
	m.resMapLock.Lock()
	delete(m.resMap, reqId)
	m.resMapLock.Unlock()
}

func (m *ConnManager) GoHandleReceiveRpcMsg() {
	for {
		select {
		case <-m.ctx.Done():
			return
		case msg := <-m.ReceiveChan:
			log.Info("ReceiveRpcMsg", string(msg))
			mid := GetMsgId(msg)
			m.resMapLock.RLock()
			resChan := m.resMap[mid]
			m.resMapLock.RUnlock()
			if resChan != nil {
				resChan <- msg
				close(resChan)
				m.resMapLock.Lock()
				delete(m.resMap, mid)
				m.resMapLock.Unlock()
			}

		}
	}
}
func (m *ConnManager) GetShard(clientId string) *Shard {
	return m.Shards[m.GetShardIndex(clientId)]
}

// GetConn 该函数返回的连接一定是连通的,若始终无法连接则报错
func (m *ConnManager) GetConn(clientId string) (*ClientConn, error) {

	shard := m.GetShard(clientId)
	shard.Mu.RLock()
	conn := shard.Conns[clientId]
	shard.Mu.RUnlock()
	if conn != nil {
		return conn, nil
	}

	// 如果不存在，使用写锁创建
	shard.Mu.Lock()
	defer shard.Mu.Unlock()

	// 双重检查，避免重复创建
	conn = shard.Conns[clientId]
	if conn != nil {
		return conn, nil
	}

	// 创建新连接
	conn, err := NewClientConn(m.ctx, "", "", clientId, m.ReceiveChan)
	if err != nil {
		return nil, fmt.Errorf("client Connect Failed: %v", err)
	}
	shard.Conns[clientId] = conn
	return conn, nil
}

func (m *ConnManager) GetShardIndex(clientID string) int {
	// 步骤1：重置哈希器（复用实例，避免每次创建新对象）
	h := m.HashPool.Get().(hash.Hash32)
	h.Reset()

	// 步骤2：写入clientID字符串，计算哈希值
	// 注意：Write方法对字符串的处理是按字节（UTF-8编码），支持包含字母、符号等
	_, _ = h.Write([]byte(clientID))
	hashValue := h.Sum32() // 32位哈希值
	m.HashPool.Put(h)
	// 步骤3：映射到分片索引
	if isPowerOfTwo(m.ShardCount) {
		// 若分片数是2的幂，用位运算&代替%，效率更高
		return int(hashValue & uint32(m.ShardCount-1))
	} else {
		// 非2的幂，用取模
		return int(hashValue % uint32(m.ShardCount))
	}
}

func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

func GetMsgId(msg []byte) string {
	m, err := ParseResponse(msg)
	if err != nil {
		return ""
	}
	return m.GetID()
}
